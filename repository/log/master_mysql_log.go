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

// MasterMySQLLogRepository 主服务MySQL日志仓储接口
type MasterMySQLLogRepository interface {
	Create(log *dto.MasterMySQLLog) error
	Update(log *dto.MasterMySQLLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.MasterMySQLLog, error)
	GetByRequestID(requestID string) (*dto.MasterMySQLLog, error)
	List(page, pageSize int) ([]*dto.MasterMySQLLog, int64, error)
	ListByMaster(masterID string, page, pageSize int) ([]*dto.MasterMySQLLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.MasterMySQLLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.MasterMySQLLog, int64, error)
	Benchmark(count int) error
}

type masterMySQLLogRepository struct {
	db *gorm.DB
}

// NewMasterMySQLLogRepository 创建主服务MySQL日志仓储实例
func NewMasterMySQLLogRepository(db *gorm.DB) MasterMySQLLogRepository {
	return &masterMySQLLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *masterMySQLLogRepository) convertToDTO(model *log.MasterMySQLLog) *dto.MasterMySQLLog {
	if model == nil {
		return nil
	}

	var logDetails dto.MasterMySQLLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.MasterMySQLLog{
		MasterMySQLLogID: model.MasterMySQLLogID,
		MasterID:         model.MasterID,
		RequestID:        model.RequestID,
		EventType:        model.EventType,
		LogLevel:         model.LogLevel,
		LogDetails:       logDetails,
		Operation:        model.Operation,
		ExecutionTime:    model.ExecutionTime,
		ErrorType:        model.ErrorType,
		ErrorMessage:     model.ErrorMessage,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *masterMySQLLogRepository) convertToModel(dto *dto.MasterMySQLLog) (*log.MasterMySQLLog, error) {
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

	return &log.MasterMySQLLog{
		MasterMySQLLogID: dto.MasterMySQLLogID,
		MasterID:         dto.MasterID,
		RequestID:        dto.RequestID,
		EventType:        dto.EventType,
		LogLevel:         dto.LogLevel,
		LogDetails:       logDetailsJSON,
		Operation:        dto.Operation,
		ExecutionTime:    dto.ExecutionTime,
		ErrorType:        dto.ErrorType,
		ErrorMessage:     dto.ErrorMessage,
		CreatedAt:        dto.CreatedAt,
		UpdatedAt:        dto.UpdatedAt,
		DeletedAt:        deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *masterMySQLLogRepository) Create(masterMySQLLog *dto.MasterMySQLLog) error {
	model, err := r.convertToModel(masterMySQLLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *masterMySQLLogRepository) Update(masterMySQLLog *dto.MasterMySQLLog) error {
	model, err := r.convertToModel(masterMySQLLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.MasterMySQLLog{}).Where("master_mysql_log_id = ?", masterMySQLLog.MasterMySQLLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *masterMySQLLogRepository) Delete(masterMySQLLogID string) error {
	return r.db.Delete(&log.MasterMySQLLog{}, "master_mysql_log_id = ?", masterMySQLLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *masterMySQLLogRepository) GetByID(masterMySQLLogID string) (*dto.MasterMySQLLog, error) {
	var masterMySQLLog log.MasterMySQLLog
	if err := r.db.Where("master_mysql_log_id = ?", masterMySQLLogID).First(&masterMySQLLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&masterMySQLLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *masterMySQLLogRepository) GetByRequestID(requestID string) (*dto.MasterMySQLLog, error) {
	var masterMySQLLog log.MasterMySQLLog
	if err := r.db.Where("request_id = ?", requestID).First(&masterMySQLLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&masterMySQLLog), nil
}

// List 获取日志记录列表
func (r *masterMySQLLogRepository) List(page, pageSize int) ([]*dto.MasterMySQLLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.MasterMySQLLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.MasterMySQLLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count master mysql logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find master mysql logs failed: %w", err)
	}

	dtoList := make([]*dto.MasterMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByMaster 获取指定主服务的日志记录列表
func (r *masterMySQLLogRepository) ListByMaster(masterID string, page, pageSize int) ([]*dto.MasterMySQLLog, int64, error) {
	if masterID == "" {
		return nil, 0, fmt.Errorf("master_id cannot be empty")
	}
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.MasterMySQLLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.MasterMySQLLog{}).Where("master_id = ?", masterID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count master mysql logs failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find master mysql logs failed: %w", err)
	}

	dtoList := make([]*dto.MasterMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *masterMySQLLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.MasterMySQLLog, int64, error) {
	var total int64
	var logs []log.MasterMySQLLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.MasterMySQLLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MasterMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *masterMySQLLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.MasterMySQLLog, int64, error) {
	var total int64
	var logs []log.MasterMySQLLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.MasterMySQLLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MasterMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *masterMySQLLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行主服务MySQL日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.MasterMySQLLog{
			MasterMySQLLogID: utils.GenerateRandomUUID(12),
			MasterID:         utils.GenerateRandomUUID(12),
			RequestID:        fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			EventType:        []string{"connect", "query", "error"}[rand.Intn(3)],
			LogLevel:         []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.MasterMySQLLogDetails{
				Content: fmt.Sprintf("SELECT * FROM test_table_%d", i),
			},
			Operation:     []string{"SELECT", "INSERT", "UPDATE", "DELETE"}[rand.Intn(4)],
			ExecutionTime: rand.Intn(1000), // 0-1000ms
			ErrorType:     "test_error",
			ErrorMessage:  fmt.Sprintf("测试错误消息 %d", i),
		}

		// 创建
		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdLog, err := r.GetByID(testLog.MasterMySQLLogID)
		if err != nil {
			utils.SysError("获取创建的日志记录失败: " + err.Error())
			return err
		}

		// 更新
		createdLog.Operation = []string{"SELECT", "INSERT", "UPDATE", "DELETE"}[rand.Intn(4)]
		createdLog.ExecutionTime = rand.Intn(1000)
		if err := r.Update(createdLog); err != nil {
			utils.SysError("更新日志记录失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdLog.MasterMySQLLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
