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

// WorkerLogRepository 工作节点日志仓储接口
type WorkerLogRepository interface {
	Create(log *dto.WorkerLog) error
	Update(log *dto.WorkerLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.WorkerLog, error)
	GetByRequestID(requestID string) (*dto.WorkerLog, error)
	List(page, pageSize int) ([]*dto.WorkerLog, int64, error)
	ListByWorker(workerID string, page, pageSize int) ([]*dto.WorkerLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.WorkerLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.WorkerLog, int64, error)
	Benchmark(count int) error
}

type workerLogRepository struct {
	db *gorm.DB
}

// NewWorkerLogRepository 创建工作节点日志仓储实例
func NewWorkerLogRepository(db *gorm.DB) WorkerLogRepository {
	return &workerLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *workerLogRepository) convertToDTO(model *log.WorkerLog) *dto.WorkerLog {
	if model == nil {
		return nil
	}

	var logDetails dto.WorkerLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.WorkerLog{
		WorkerLogID:     model.WorkerLogID,
		WorkerClusterID: model.WorkerClusterID,
		WorkerGroupID:   model.WorkerGroupID,
		WorkerNodeID:    model.WorkerNodeID,
		RequestID:       model.RequestID,
		EventType:       model.EventType,
		LogLevel:        model.LogLevel,
		LogDetails:      logDetails,
		ResourceType:    model.ResourceType,
		ResourceUsage:   model.ResourceUsage,
		ErrorType:       model.ErrorType,
		ErrorMessage:    model.ErrorMessage,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		DeletedAt:       deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *workerLogRepository) convertToModel(dto *dto.WorkerLog) (*log.WorkerLog, error) {
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

	return &log.WorkerLog{
		WorkerLogID:     dto.WorkerLogID,
		WorkerClusterID: dto.WorkerClusterID,
		WorkerGroupID:   dto.WorkerGroupID,
		WorkerNodeID:    dto.WorkerNodeID,
		RequestID:       dto.RequestID,
		EventType:       dto.EventType,
		LogLevel:        dto.LogLevel,
		LogDetails:      logDetailsJSON,
		ResourceType:    dto.ResourceType,
		ResourceUsage:   dto.ResourceUsage,
		ErrorType:       dto.ErrorType,
		ErrorMessage:    dto.ErrorMessage,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
		DeletedAt:       deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *workerLogRepository) Create(workerLog *dto.WorkerLog) error {
	model, err := r.convertToModel(workerLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *workerLogRepository) Update(workerLog *dto.WorkerLog) error {
	model, err := r.convertToModel(workerLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.WorkerLog{}).Where("worker_log_id = ?", workerLog.WorkerLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *workerLogRepository) Delete(workerLogID string) error {
	return r.db.Delete(&log.WorkerLog{}, "worker_log_id = ?", workerLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *workerLogRepository) GetByID(workerLogID string) (*dto.WorkerLog, error) {
	var workerLog log.WorkerLog
	if err := r.db.Where("worker_log_id = ?", workerLogID).First(&workerLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&workerLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *workerLogRepository) GetByRequestID(requestID string) (*dto.WorkerLog, error) {
	var workerLog log.WorkerLog
	if err := r.db.Where("request_id = ?", requestID).First(&workerLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&workerLog), nil
}

// List 获取日志记录列表
func (r *workerLogRepository) List(page, pageSize int) ([]*dto.WorkerLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.WorkerLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.WorkerLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker logs failed: %w", err)
	}

	dtoList := make([]*dto.WorkerLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByWorker 获取指定工作节点的日志记录列表
func (r *workerLogRepository) ListByWorker(workerID string, page, pageSize int) ([]*dto.WorkerLog, int64, error) {
	var total int64
	var logs []log.WorkerLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.WorkerLog{}).Where("worker_id = ?", workerID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.WorkerLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *workerLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.WorkerLog, int64, error) {
	var total int64
	var logs []log.WorkerLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.WorkerLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.WorkerLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *workerLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.WorkerLog, int64, error) {
	var total int64
	var logs []log.WorkerLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.WorkerLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.WorkerLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *workerLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行工作节点日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.WorkerLog{
			WorkerLogID:     utils.GenerateRandomUUID(12),
			WorkerClusterID: utils.GenerateRandomUUID(12),
			WorkerGroupID:   utils.GenerateRandomUUID(12),
			WorkerNodeID:    utils.GenerateRandomUUID(12),
			RequestID:       fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			EventType:       []string{"task", "process", "error"}[rand.Intn(3)],
			LogLevel:        []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.WorkerLogDetails{
				Content: fmt.Sprintf("test_value_%d", i),
			},
			ResourceType:  []string{"cpu", "memory", "disk"}[rand.Intn(3)],
			ResourceUsage: common.JSON{},
			ErrorType:     "test_error",
			ErrorMessage:  fmt.Sprintf("测试错误消息 %d", i),
		}

		// 创建
		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdLog, err := r.GetByID(testLog.WorkerLogID)
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
		if err := r.Delete(createdLog.WorkerLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
