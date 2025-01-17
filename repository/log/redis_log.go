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

// RedisLogRepository Redis日志仓储接口
type RedisLogRepository interface {
	Create(log *dto.RedisLog) error
	Update(log *dto.RedisLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.RedisLog, error)
	GetByRequestID(requestID string) (*dto.RedisLog, error)
	List(page, pageSize int) ([]*dto.RedisLog, int64, error)
	ListByNode(nodeID string, page, pageSize int) ([]*dto.RedisLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.RedisLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.RedisLog, int64, error)
	Benchmark(count int) error
}

type redisLogRepository struct {
	db *gorm.DB
}

// NewRedisLogRepository 创建Redis日志仓储实例
func NewRedisLogRepository(db *gorm.DB) RedisLogRepository {
	return &redisLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *redisLogRepository) convertToDTO(model *log.RedisLog) *dto.RedisLog {
	if model == nil {
		return nil
	}

	var logDetails dto.RedisLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.RedisLog{
		RedisLogID:    model.RedisLogID,
		ServiceID:     model.ServiceID,
		RequestID:     model.RequestID,
		NodeType:      model.NodeType,
		EventType:     model.EventType,
		LogLevel:      model.LogLevel,
		LogDetails:    logDetails,
		Operation:     model.Operation,
		KeyPattern:    model.KeyPattern,
		ExecutionTime: model.ExecutionTime,
		ErrorType:     model.ErrorType,
		ErrorMessage:  model.ErrorMessage,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *redisLogRepository) convertToModel(dto *dto.RedisLog) (*log.RedisLog, error) {
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

	return &log.RedisLog{
		RedisLogID:    dto.RedisLogID,
		ServiceID:     dto.ServiceID,
		RequestID:     dto.RequestID,
		NodeType:      dto.NodeType,
		EventType:     dto.EventType,
		LogLevel:      dto.LogLevel,
		LogDetails:    logDetailsJSON,
		Operation:     dto.Operation,
		KeyPattern:    dto.KeyPattern,
		ExecutionTime: dto.ExecutionTime,
		ErrorType:     dto.ErrorType,
		ErrorMessage:  dto.ErrorMessage,
		CreatedAt:     dto.CreatedAt,
		UpdatedAt:     dto.UpdatedAt,
		DeletedAt:     deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *redisLogRepository) Create(redisLog *dto.RedisLog) error {
	model, err := r.convertToModel(redisLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *redisLogRepository) Update(redisLog *dto.RedisLog) error {
	model, err := r.convertToModel(redisLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.RedisLog{}).Where("redis_log_id = ?", redisLog.RedisLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *redisLogRepository) Delete(redisLogID string) error {
	return r.db.Delete(&log.RedisLog{}, "redis_log_id = ?", redisLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *redisLogRepository) GetByID(redisLogID string) (*dto.RedisLog, error) {
	var redisLog log.RedisLog
	if err := r.db.Where("redis_log_id = ?", redisLogID).First(&redisLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&redisLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *redisLogRepository) GetByRequestID(requestID string) (*dto.RedisLog, error) {
	var redisLog log.RedisLog
	if err := r.db.Where("request_id = ?", requestID).First(&redisLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&redisLog), nil
}

// List 获取日志记录列表
func (r *redisLogRepository) List(page, pageSize int) ([]*dto.RedisLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.RedisLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.RedisLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count redis logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find redis logs failed: %w", err)
	}

	dtoList := make([]*dto.RedisLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByNode 获取指定节点的日志记录列表
func (r *redisLogRepository) ListByNode(nodeID string, page, pageSize int) ([]*dto.RedisLog, int64, error) {
	var total int64
	var logs []log.RedisLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RedisLog{}).Where("node_id = ?", nodeID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RedisLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *redisLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.RedisLog, int64, error) {
	var total int64
	var logs []log.RedisLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RedisLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RedisLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *redisLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.RedisLog, int64, error) {
	var total int64
	var logs []log.RedisLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RedisLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RedisLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *redisLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行Redis日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.RedisLog{
			RedisLogID: utils.GenerateRandomUUID(12),
			ServiceID:  utils.GenerateRandomUUID(12),
			RequestID:  fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			NodeType:   []string{"master", "worker"}[rand.Intn(2)],
			EventType:  []string{"command", "connection", "error"}[rand.Intn(3)],
			LogLevel:   []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.RedisLogDetails{
				Content: fmt.Sprintf("test_value_%d", i),
			},
			Operation:     []string{"GET", "SET", "DEL", "HGET"}[rand.Intn(4)],
			KeyPattern:    fmt.Sprintf("test:key:%d", i),
			ExecutionTime: rand.Intn(1000),
			ErrorType:     "test_error",
			ErrorMessage:  fmt.Sprintf("测试错误消息 %d", i),
		}

		// 创建
		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdLog, err := r.GetByID(testLog.RedisLogID)
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
		if err := r.Delete(createdLog.RedisLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
