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

// GatewayLogRepository 网关日志仓储接口
type GatewayLogRepository interface {
	Create(log *dto.GatewayLog) error
	Update(log *dto.GatewayLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.GatewayLog, error)
	GetByRequestID(requestID string) (*dto.GatewayLog, error)
	List(page, pageSize int) ([]*dto.GatewayLog, int64, error)
	ListByGateway(gatewayID string, page, pageSize int) ([]*dto.GatewayLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.GatewayLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.GatewayLog, int64, error)
	Benchmark(count int) error
}

type gatewayLogRepository struct {
	db *gorm.DB
}

// NewGatewayLogRepository 创建网关日志仓储实例
func NewGatewayLogRepository(db *gorm.DB) GatewayLogRepository {
	return &gatewayLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *gatewayLogRepository) convertToDTO(model *log.GatewayLog) *dto.GatewayLog {
	if model == nil {
		return nil
	}

	var logDetails dto.GatewayLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.GatewayLog{
		GatewayLogID: model.GatewayLogID,
		GatewayID:    model.GatewayID,
		EventType:    model.EventType,
		LogLevel:     model.LogLevel,
		LogDetails:   logDetails,
		ErrorType:    model.ErrorType,
		ErrorMessage: model.ErrorMessage,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *gatewayLogRepository) convertToModel(dto *dto.GatewayLog) (*log.GatewayLog, error) {
	if dto == nil {
		return nil, nil
	}

	logDetailsJSON, err := common.FromStruct(dto.LogDetails)
	if err != nil {
		return nil, fmt.Errorf("转换日志详细信息失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &log.GatewayLog{
		GatewayLogID: dto.GatewayLogID,
		GatewayID:    dto.GatewayID,
		EventType:    dto.EventType,
		LogLevel:     dto.LogLevel,
		LogDetails:   logDetailsJSON,
		ErrorType:    dto.ErrorType,
		ErrorMessage: dto.ErrorMessage,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		DeletedAt:    deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *gatewayLogRepository) Create(gatewayLog *dto.GatewayLog) error {
	model, err := r.convertToModel(gatewayLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *gatewayLogRepository) Update(gatewayLog *dto.GatewayLog) error {
	model, err := r.convertToModel(gatewayLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.GatewayLog{}).Where("gateway_log_id = ?", gatewayLog.GatewayLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *gatewayLogRepository) Delete(gatewayLogID string) error {
	return r.db.Delete(&log.GatewayLog{}, "gateway_log_id = ?", gatewayLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *gatewayLogRepository) GetByID(gatewayLogID string) (*dto.GatewayLog, error) {
	var gatewayLog log.GatewayLog
	if err := r.db.Where("gateway_log_id = ?", gatewayLogID).First(&gatewayLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&gatewayLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *gatewayLogRepository) GetByRequestID(requestID string) (*dto.GatewayLog, error) {
	var gatewayLog log.GatewayLog
	if err := r.db.Where("request_id = ?", requestID).First(&gatewayLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&gatewayLog), nil
}

// List 获取日志记录列表
func (r *gatewayLogRepository) List(page, pageSize int) ([]*dto.GatewayLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.GatewayLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.GatewayLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count gateway logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find gateway logs failed: %w", err)
	}

	dtoList := make([]*dto.GatewayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByGateway 获取指定网关的日志记录列表
func (r *gatewayLogRepository) ListByGateway(gatewayID string, page, pageSize int) ([]*dto.GatewayLog, int64, error) {
	if gatewayID == "" {
		return nil, 0, fmt.Errorf("gateway_id cannot be empty")
	}
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.GatewayLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.GatewayLog{}).Where("gateway_id = ?", gatewayID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count gateway logs failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find gateway logs failed: %w", err)
	}

	dtoList := make([]*dto.GatewayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *gatewayLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.GatewayLog, int64, error) {
	var total int64
	var logs []log.GatewayLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.GatewayLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.GatewayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *gatewayLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.GatewayLog, int64, error) {
	var total int64
	var logs []log.GatewayLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.GatewayLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.GatewayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *gatewayLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行网关日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.GatewayLog{
			GatewayLogID: utils.GenerateRandomUUID(12),
			GatewayID:    utils.GenerateRandomUUID(12),
			LogLevel:     []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.GatewayLogDetails{
				Content: "test_value",
			},
			EventType:    []string{"request", "response", "error"}[rand.Intn(3)],
			ErrorMessage: fmt.Sprintf("测试错误消息 %d", i),
			ErrorType:    []string{"rate_limit", "server_error", "timeout", "network"}[rand.Intn(4)],
		}

		// 创建
		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdLog, err := r.GetByID(testLog.GatewayLogID)
		if err != nil {
			utils.SysError("获取创建的日志记录失败: " + err.Error())
			return err
		}

		// 更新
		createdLog.LogLevel = []string{"info", "warn", "error"}[rand.Intn(3)]
		if err := r.Update(createdLog); err != nil {
			utils.SysError("更新日志记录失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdLog.GatewayLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
