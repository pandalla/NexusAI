package log

import (
	"fmt"
	"math/rand"
	"nexus-ai/common"
	dto "nexus-ai/dto/model/log"
	"nexus-ai/model/log"
	"nexus-ai/utils"
	"time"

	"gorm.io/gorm"
)

// MasterLogRepository 主服务日志仓储接口
type MasterLogRepository interface {
	Create(masterLog *dto.MasterLog) error
	Update(masterLog *dto.MasterLog) error
	Delete(masterLogID string) error
	GetByID(masterLogID string) (*dto.MasterLog, error)
	GetByRequestID(requestID string) (*dto.MasterLog, error)
	List(page, pageSize int) ([]*dto.MasterLog, int64, error)
	ListByMaster(masterID string, page, pageSize int) ([]*dto.MasterLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.MasterLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.MasterLog, int64, error)
	Benchmark(count int) error
}

type masterLogRepository struct {
	db *gorm.DB
}

// NewMasterLogRepository 创建主服务日志仓储实例
func NewMasterLogRepository(db *gorm.DB) MasterLogRepository {
	return &masterLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *masterLogRepository) convertToDTO(model *log.MasterLog) *dto.MasterLog {
	if model == nil {
		return nil
	}

	var logDetails dto.MasterLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.MasterLog{
		MasterLogID:  model.MasterLogID,
		MasterID:     model.MasterID,
		EventType:    model.EventType,
		LogLevel:     model.LogLevel,
		LogDetails:   logDetails,
		RequestID:    model.RequestID,
		ErrorType:    model.ErrorType,
		ErrorMessage: model.ErrorMessage,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *masterLogRepository) convertToModel(dto *dto.MasterLog) (*log.MasterLog, error) {
	if dto == nil {
		return nil, nil
	}

	logDetailsJSON, err := common.FromStruct(dto.LogDetails)
	if err != nil {
		return nil, fmt.Errorf("转换日志详情失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &log.MasterLog{
		MasterLogID:  dto.MasterLogID,
		MasterID:     dto.MasterID,
		EventType:    dto.EventType,
		LogLevel:     dto.LogLevel,
		LogDetails:   logDetailsJSON,
		RequestID:    dto.RequestID,
		ErrorType:    dto.ErrorType,
		ErrorMessage: dto.ErrorMessage,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		DeletedAt:    deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *masterLogRepository) Create(masterLog *dto.MasterLog) error {
	model, err := r.convertToModel(masterLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *masterLogRepository) Update(masterLog *dto.MasterLog) error {
	model, err := r.convertToModel(masterLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.MasterLog{}).Where("master_log_id = ?", masterLog.MasterLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *masterLogRepository) Delete(masterLogID string) error {
	return r.db.Delete(&log.MasterLog{}, "master_log_id = ?", masterLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *masterLogRepository) GetByID(masterLogID string) (*dto.MasterLog, error) {
	var masterLog log.MasterLog
	if err := r.db.Where("master_log_id = ?", masterLogID).First(&masterLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&masterLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *masterLogRepository) GetByRequestID(requestID string) (*dto.MasterLog, error) {
	var masterLog log.MasterLog
	if err := r.db.Where("request_id = ?", requestID).First(&masterLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&masterLog), nil
}

// List 获取日志记录列表
func (r *masterLogRepository) List(page, pageSize int) ([]*dto.MasterLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.MasterLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.MasterLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count master logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find master logs failed: %w", err)
	}

	dtoList := make([]*dto.MasterLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByMaster 获取指定主服务的日志记录列表
func (r *masterLogRepository) ListByMaster(masterID string, page, pageSize int) ([]*dto.MasterLog, int64, error) {
	var masterLogs []log.MasterLog
	var total int64

	if err := r.db.Model(&log.MasterLog{}).Where("master_id = ?", masterID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("master_id = ?", masterID).Offset(offset).Limit(pageSize).Find(&masterLogs).Error; err != nil {
		return nil, 0, err
	}

	dtos := make([]*dto.MasterLog, len(masterLogs))
	for i, masterLog := range masterLogs {
		dtos[i] = r.convertToDTO(&masterLog)
	}

	return dtos, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *masterLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.MasterLog, int64, error) {
	if logLevel == "" {
		return nil, 0, fmt.Errorf("log_level cannot be empty")
	}
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.MasterLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.MasterLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count master logs failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find master logs failed: %w", err)
	}

	dtos := make([]*dto.MasterLog, len(logs))
	for i, l := range logs {
		dtos[i] = r.convertToDTO(&l)
	}

	return dtos, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *masterLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.MasterLog, int64, error) {
	if eventType == "" {
		return nil, 0, fmt.Errorf("event_type cannot be empty")
	}
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.MasterLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.MasterLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count master logs failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find master logs failed: %w", err)
	}

	dtos := make([]*dto.MasterLog, len(logs))
	for i, l := range logs {
		dtos[i] = r.convertToDTO(&l)
	}

	return dtos, total, nil
}

// Benchmark 执行基准测试
func (r *masterLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行主服务日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.MasterLog{
			MasterLogID: utils.GenerateRandomUUID(12),
			MasterID:    utils.GenerateRandomUUID(12),
			RequestID:   fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			EventType:   []string{"service", "redis", "mysql", "worker"}[rand.Intn(4)],
			LogLevel:    []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.MasterLogDetails{
				Content: fmt.Sprintf("test_value_%d", i),
			},
			ErrorType:    "test_error",
			ErrorMessage: fmt.Sprintf("测试错误消息 %d", i),
		}

		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		if _, err := r.GetByID(testLog.MasterLogID); err != nil {
			utils.SysError("获取日志记录失败: " + err.Error())
			return err
		}

		testLog.LogLevel = []string{"info", "warn", "error"}[rand.Intn(3)]
		if err := r.Update(testLog); err != nil {
			utils.SysError("更新日志记录失败: " + err.Error())
			return err
		}

		if err := r.Delete(testLog.MasterLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo(fmt.Sprintf("基准测试完成，总耗时: %s, 平均每组操作耗时: %s",
		duration.String(),
		(duration / time.Duration(count)).String()))
	return nil
}
