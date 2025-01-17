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

// RelayLogRepository 中继日志仓储接口
type RelayLogRepository interface {
	Create(log *dto.RelayLog) error
	Update(log *dto.RelayLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.RelayLog, error)
	GetByRequestID(requestID string) (*dto.RelayLog, error)
	List(page, pageSize int) ([]*dto.RelayLog, int64, error)
	ListByRelay(relayID string, page, pageSize int) ([]*dto.RelayLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.RelayLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.RelayLog, int64, error)
	Benchmark(count int) error
}

type relayLogRepository struct {
	db *gorm.DB
}

// NewRelayLogRepository 创建中继日志仓储实例
func NewRelayLogRepository(db *gorm.DB) RelayLogRepository {
	return &relayLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *relayLogRepository) convertToDTO(model *log.RelayLog) *dto.RelayLog {
	if model == nil {
		return nil
	}

	var logDetails dto.RelayLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.RelayLog{
		RelayLogID:        model.RelayLogID,
		UserID:            model.UserID,
		ChannelID:         model.ChannelID,
		ModelID:           model.ModelID,
		TokenID:           model.TokenID,
		WorkerClusterID:   model.WorkerClusterID,
		WorkerGroupID:     model.WorkerGroupID,
		WorkerNodeID:      model.WorkerNodeID,
		RequestID:         model.RequestID,
		RelayStatus:       model.RelayStatus,
		UpstreamRequestID: model.UpstreamRequestID,
		UpstreamStatus:    model.UpstreamStatus,
		UpstreamLatency:   model.UpstreamLatency,
		TotalLatency:      model.TotalLatency,
		RetryCount:        model.RetryCount,
		RetryResult:       model.RetryResult,
		ErrorType:         model.ErrorType,
		ErrorMessage:      model.ErrorMessage,
		RequestHeaders:    model.RequestHeaders,
		RequestBody:       model.RequestBody,
		ResponseHeaders:   model.ResponseHeaders,
		ResponseBody:      model.ResponseBody,
		RequestTokens:     model.RequestTokens,
		ResponseTokens:    model.ResponseTokens,
		QuotaConsumed:     model.QuotaConsumed,
		EventType:         model.EventType,
		LogLevel:          model.LogLevel,
		LogDetails:        logDetails,
		CreatedAt:         model.CreatedAt,
		UpdatedAt:         model.UpdatedAt,
		DeletedAt:         deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *relayLogRepository) convertToModel(dto *dto.RelayLog) (*log.RelayLog, error) {
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

	return &log.RelayLog{
		RelayLogID:        dto.RelayLogID,
		UserID:            dto.UserID,
		ChannelID:         dto.ChannelID,
		ModelID:           dto.ModelID,
		TokenID:           dto.TokenID,
		WorkerClusterID:   dto.WorkerClusterID,
		WorkerGroupID:     dto.WorkerGroupID,
		WorkerNodeID:      dto.WorkerNodeID,
		RequestID:         dto.RequestID,
		RelayStatus:       dto.RelayStatus,
		UpstreamRequestID: dto.UpstreamRequestID,
		UpstreamStatus:    dto.UpstreamStatus,
		UpstreamLatency:   dto.UpstreamLatency,
		TotalLatency:      dto.TotalLatency,
		RetryCount:        dto.RetryCount,
		RetryResult:       dto.RetryResult,
		ErrorType:         dto.ErrorType,
		ErrorMessage:      dto.ErrorMessage,
		RequestHeaders:    dto.RequestHeaders,
		RequestBody:       dto.RequestBody,
		ResponseHeaders:   dto.ResponseHeaders,
		ResponseBody:      dto.ResponseBody,
		RequestTokens:     dto.RequestTokens,
		ResponseTokens:    dto.ResponseTokens,
		QuotaConsumed:     dto.QuotaConsumed,
		EventType:         dto.EventType,
		LogLevel:          dto.LogLevel,
		LogDetails:        logDetailsJSON,
		CreatedAt:         dto.CreatedAt,
		UpdatedAt:         dto.UpdatedAt,
		DeletedAt:         deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *relayLogRepository) Create(relayLog *dto.RelayLog) error {
	model, err := r.convertToModel(relayLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *relayLogRepository) Update(relayLog *dto.RelayLog) error {
	model, err := r.convertToModel(relayLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.RelayLog{}).Where("relay_log_id = ?", relayLog.RelayLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *relayLogRepository) Delete(relayLogID string) error {
	return r.db.Delete(&log.RelayLog{}, "relay_log_id = ?", relayLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *relayLogRepository) GetByID(relayLogID string) (*dto.RelayLog, error) {
	var relayLog log.RelayLog
	if err := r.db.Where("relay_log_id = ?", relayLogID).First(&relayLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&relayLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *relayLogRepository) GetByRequestID(requestID string) (*dto.RelayLog, error) {
	var relayLog log.RelayLog
	if err := r.db.Where("request_id = ?", requestID).First(&relayLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&relayLog), nil
}

// List 获取日志记录列表
func (r *relayLogRepository) List(page, pageSize int) ([]*dto.RelayLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.RelayLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.RelayLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count relay logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find relay logs failed: %w", err)
	}

	dtoList := make([]*dto.RelayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByRelay 获取指定中继节点的日志记录列表
func (r *relayLogRepository) ListByRelay(relayID string, page, pageSize int) ([]*dto.RelayLog, int64, error) {
	var total int64
	var logs []log.RelayLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RelayLog{}).Where("relay_id = ?", relayID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RelayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *relayLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.RelayLog, int64, error) {
	var total int64
	var logs []log.RelayLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RelayLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RelayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *relayLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.RelayLog, int64, error) {
	var total int64
	var logs []log.RelayLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RelayLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RelayLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *relayLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行中继日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.RelayLog{
			RelayLogID:        utils.GenerateRandomUUID(12),
			UserID:            utils.GenerateRandomUUID(12),
			ChannelID:         utils.GenerateRandomUUID(12),
			ModelID:           utils.GenerateRandomUUID(12),
			TokenID:           utils.GenerateRandomUUID(12),
			WorkerClusterID:   utils.GenerateRandomUUID(12),
			WorkerGroupID:     utils.GenerateRandomUUID(12),
			WorkerNodeID:      utils.GenerateRandomUUID(12),
			RequestID:         fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			RelayStatus:       []int8{1, 2, 3, 4, 5}[rand.Intn(5)],
			UpstreamRequestID: fmt.Sprintf("UPSTREAM_REQ_%s", utils.GenerateRandomString(16)),
			UpstreamStatus:    rand.Intn(600),
			UpstreamLatency:   rand.Intn(5000),
			TotalLatency:      rand.Intn(10000),
			RetryCount:        rand.Intn(5),
			RetryResult:       common.JSON{},
			ErrorType:         []string{"rate_limit", "server_error", "timeout", "network"}[rand.Intn(4)],
			ErrorMessage:      fmt.Sprintf("测试错误消息 %d", i),
			RequestHeaders:    common.JSON{},
			RequestBody:       common.JSON{},
			ResponseHeaders:   common.JSON{},
			ResponseBody:      common.JSON{},
			RequestTokens:     rand.Intn(100),
			ResponseTokens:    rand.Intn(100),
			QuotaConsumed:     rand.Float64() * 100,
			EventType:         []string{"request", "response"}[rand.Intn(2)],
			LogLevel:          []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.RelayLogDetails{
				Content: fmt.Sprintf("test_value_%d", i),
			},
		}

		// 创建
		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdLog, err := r.GetByID(testLog.RelayLogID)
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
		if err := r.Delete(createdLog.RelayLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
