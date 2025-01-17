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

// WorkerMySQLLogRepository 工作节点MySQL日志仓储接口
type WorkerMySQLLogRepository interface {
	Create(log *dto.WorkerMySQLLog) error
	Update(log *dto.WorkerMySQLLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.WorkerMySQLLog, error)
	GetByRequestID(requestID string) (*dto.WorkerMySQLLog, error)
	List(page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error)
	ListByWorker(workerID string, page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error)
	Benchmark(count int) error
}

type workerMySQLLogRepository struct {
	db *gorm.DB
}

// NewWorkerMySQLLogRepository 创建工作节点MySQL日志仓储实例
func NewWorkerMySQLLogRepository(db *gorm.DB) WorkerMySQLLogRepository {
	return &workerMySQLLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *workerMySQLLogRepository) convertToDTO(model *log.WorkerMySQLLog) *dto.WorkerMySQLLog {
	if model == nil {
		return nil
	}

	var logDetails dto.WorkerMySQLLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.WorkerMySQLLog{
		WorkerMySQLLogID: model.WorkerMySQLLogID,
		WorkerClusterID:  model.WorkerClusterID,
		WorkerGroupID:    model.WorkerGroupID,
		WorkerNodeID:     model.WorkerNodeID,
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
func (r *workerMySQLLogRepository) convertToModel(dto *dto.WorkerMySQLLog) (*log.WorkerMySQLLog, error) {
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

	return &log.WorkerMySQLLog{
		WorkerMySQLLogID: dto.WorkerMySQLLogID,
		WorkerClusterID:  dto.WorkerClusterID,
		WorkerGroupID:    dto.WorkerGroupID,
		WorkerNodeID:     dto.WorkerNodeID,
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
func (r *workerMySQLLogRepository) Create(workerMySQLLog *dto.WorkerMySQLLog) error {
	model, err := r.convertToModel(workerMySQLLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *workerMySQLLogRepository) Update(workerMySQLLog *dto.WorkerMySQLLog) error {
	model, err := r.convertToModel(workerMySQLLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.WorkerMySQLLog{}).Where("worker_mysql_log_id = ?", workerMySQLLog.WorkerMySQLLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *workerMySQLLogRepository) Delete(workerMySQLLogID string) error {
	return r.db.Delete(&log.WorkerMySQLLog{}, "worker_mysql_log_id = ?", workerMySQLLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *workerMySQLLogRepository) GetByID(workerMySQLLogID string) (*dto.WorkerMySQLLog, error) {
	var workerMySQLLog log.WorkerMySQLLog
	if err := r.db.Where("worker_mysql_log_id = ?", workerMySQLLogID).First(&workerMySQLLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&workerMySQLLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *workerMySQLLogRepository) GetByRequestID(requestID string) (*dto.WorkerMySQLLog, error) {
	var workerMySQLLog log.WorkerMySQLLog
	if err := r.db.Where("request_id = ?", requestID).First(&workerMySQLLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&workerMySQLLog), nil
}

// List 获取日志记录列表
func (r *workerMySQLLogRepository) List(page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.WorkerMySQLLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.WorkerMySQLLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker mysql logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker mysql logs failed: %w", err)
	}

	dtoList := make([]*dto.WorkerMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByWorker 获取指定工作节点的日志记录列表
func (r *workerMySQLLogRepository) ListByWorker(workerID string, page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error) {
	var total int64
	var logs []log.WorkerMySQLLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.WorkerMySQLLog{}).Where("worker_id = ?", workerID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.WorkerMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *workerMySQLLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error) {
	var total int64
	var logs []log.WorkerMySQLLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.WorkerMySQLLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.WorkerMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *workerMySQLLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.WorkerMySQLLog, int64, error) {
	var total int64
	var logs []log.WorkerMySQLLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.WorkerMySQLLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.WorkerMySQLLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *workerMySQLLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行工作节点MySQL日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.WorkerMySQLLog{
			WorkerMySQLLogID: utils.GenerateRandomUUID(12),
			WorkerClusterID:  utils.GenerateRandomUUID(12),
			WorkerGroupID:    utils.GenerateRandomUUID(12),
			WorkerNodeID:     utils.GenerateRandomUUID(12),
			RequestID:        fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			EventType:        []string{"query", "transaction", "connection"}[rand.Intn(3)],
			LogLevel:         []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.WorkerMySQLLogDetails{
				Content: fmt.Sprintf("test_value_%d", i),
			},
			Operation:     []string{"SELECT", "INSERT", "UPDATE", "DELETE"}[rand.Intn(4)],
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
		createdLog, err := r.GetByID(testLog.WorkerMySQLLogID)
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
		if err := r.Delete(createdLog.WorkerMySQLLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
