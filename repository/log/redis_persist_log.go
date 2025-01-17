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

// RedisPersistLogRepository Redis持久化日志仓储接口
type RedisPersistLogRepository interface {
	Create(log *dto.RedisPersistLog) error
	Update(log *dto.RedisPersistLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.RedisPersistLog, error)
	GetByRequestID(requestID string) (*dto.RedisPersistLog, error)
	List(page, pageSize int) ([]*dto.RedisPersistLog, int64, error)
	ListByNode(nodeID string, page, pageSize int) ([]*dto.RedisPersistLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.RedisPersistLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.RedisPersistLog, int64, error)
	Benchmark(count int) error
}

type redisPersistLogRepository struct {
	db *gorm.DB
}

// NewRedisPersistLogRepository 创建Redis持久化日志仓储实例
func NewRedisPersistLogRepository(db *gorm.DB) RedisPersistLogRepository {
	return &redisPersistLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *redisPersistLogRepository) convertToDTO(model *log.RedisPersistLog) *dto.RedisPersistLog {
	if model == nil {
		return nil
	}

	var logDetails dto.RedisPersistLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.RedisPersistLog{
		RedisPersistLogID: model.RedisPersistLogID,
		NodeType:          model.NodeType,
		ServiceID:         model.ServiceID,
		EventType:         model.EventType,
		LogLevel:          model.LogLevel,
		LogDetails:        logDetails,
		PersistType:       model.PersistType,
		TargetTable:       model.TargetTable,
		DataSize:          model.DataSize,
		AffectedRows:      model.AffectedRows,
		StartTime:         model.StartTime,
		EndTime:           model.EndTime,
		Duration:          model.Duration,
		ErrorType:         model.ErrorType,
		ErrorMessage:      model.ErrorMessage,
		CreatedAt:         model.CreatedAt,
		UpdatedAt:         model.UpdatedAt,
		DeletedAt:         deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *redisPersistLogRepository) convertToModel(dto *dto.RedisPersistLog) (*log.RedisPersistLog, error) {
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

	return &log.RedisPersistLog{
		RedisPersistLogID: dto.RedisPersistLogID,
		NodeType:          dto.NodeType,
		ServiceID:         dto.ServiceID,
		EventType:         dto.EventType,
		LogLevel:          dto.LogLevel,
		LogDetails:        logDetailsJSON,
		PersistType:       dto.PersistType,
		TargetTable:       dto.TargetTable,
		DataSize:          dto.DataSize,
		AffectedRows:      dto.AffectedRows,
		StartTime:         dto.StartTime,
		EndTime:           dto.EndTime,
		Duration:          dto.Duration,
		ErrorType:         dto.ErrorType,
		ErrorMessage:      dto.ErrorMessage,
		CreatedAt:         dto.CreatedAt,
		UpdatedAt:         dto.UpdatedAt,
		DeletedAt:         deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *redisPersistLogRepository) Create(redisPersistLog *dto.RedisPersistLog) error {
	model, err := r.convertToModel(redisPersistLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *redisPersistLogRepository) Update(redisPersistLog *dto.RedisPersistLog) error {
	model, err := r.convertToModel(redisPersistLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.RedisPersistLog{}).Where("redis_persist_log_id = ?", redisPersistLog.RedisPersistLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *redisPersistLogRepository) Delete(redisPersistLogID string) error {
	return r.db.Delete(&log.RedisPersistLog{}, "redis_persist_log_id = ?", redisPersistLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *redisPersistLogRepository) GetByID(redisPersistLogID string) (*dto.RedisPersistLog, error) {
	var redisPersistLog log.RedisPersistLog
	if err := r.db.Where("redis_persist_log_id = ?", redisPersistLogID).First(&redisPersistLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&redisPersistLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *redisPersistLogRepository) GetByRequestID(requestID string) (*dto.RedisPersistLog, error) {
	var redisPersistLog log.RedisPersistLog
	if err := r.db.Where("request_id = ?", requestID).First(&redisPersistLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&redisPersistLog), nil
}

// List 获取日志记录列表
func (r *redisPersistLogRepository) List(page, pageSize int) ([]*dto.RedisPersistLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.RedisPersistLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.RedisPersistLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count redis persist logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find redis persist logs failed: %w", err)
	}

	dtoList := make([]*dto.RedisPersistLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByNode 获取指定节点的日志记录列表
func (r *redisPersistLogRepository) ListByNode(nodeID string, page, pageSize int) ([]*dto.RedisPersistLog, int64, error) {
	var total int64
	var logs []log.RedisPersistLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RedisPersistLog{}).Where("node_id = ?", nodeID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RedisPersistLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *redisPersistLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.RedisPersistLog, int64, error) {
	var total int64
	var logs []log.RedisPersistLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RedisPersistLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RedisPersistLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *redisPersistLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.RedisPersistLog, int64, error) {
	var total int64
	var logs []log.RedisPersistLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RedisPersistLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RedisPersistLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *redisPersistLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行Redis持久化日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.RedisPersistLog{
			RedisPersistLogID: utils.GenerateRandomUUID(12),
			NodeType:          []string{"master", "worker"}[rand.Intn(2)],
			ServiceID:         utils.GenerateRandomUUID(12),
			EventType:         []string{"start", "complete", "error"}[rand.Intn(3)],
			LogLevel:          []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.RedisPersistLogDetails{
				Content: fmt.Sprintf("test_value_%d", i),
			},
			PersistType:  []string{"rdb", "aof", "mixed"}[rand.Intn(3)],
			TargetTable:  fmt.Sprintf("test_table_%d", i),
			DataSize:     rand.Int63n(1024 * 1024 * 100), // 100MB以内
			AffectedRows: rand.Intn(10000),
			StartTime:    utils.MySQLTime(time.Now()),
			EndTime:      utils.MySQLTime(time.Now().Add(time.Duration(rand.Intn(5000)) * time.Millisecond)),
			Duration:     rand.Intn(5000), // 5秒以内
			ErrorType:    "test_error",
			ErrorMessage: fmt.Sprintf("测试错误消息 %d", i),
		}

		// 创建
		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdLog, err := r.GetByID(testLog.RedisPersistLogID)
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
		if err := r.Delete(createdLog.RedisPersistLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
