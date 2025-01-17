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

// RequestLogRepository 请求日志仓储接口
type RequestLogRepository interface {
	Create(log *dto.RequestLog) error
	Update(log *dto.RequestLog) error
	Delete(logID string) error
	GetByID(logID string) (*dto.RequestLog, error)
	GetByRequestID(requestID string) (*dto.RequestLog, error)
	List(page, pageSize int) ([]*dto.RequestLog, int64, error)
	ListByUser(userID string, page, pageSize int) ([]*dto.RequestLog, int64, error)
	ListByLevel(logLevel string, page, pageSize int) ([]*dto.RequestLog, int64, error)
	ListByEventType(eventType string, page, pageSize int) ([]*dto.RequestLog, int64, error)
	Benchmark(count int) error
}

type requestLogRepository struct {
	db *gorm.DB
}

// NewRequestLogRepository 创建请求日志仓储实例
func NewRequestLogRepository(db *gorm.DB) RequestLogRepository {
	return &requestLogRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *requestLogRepository) convertToDTO(model *log.RequestLog) *dto.RequestLog {
	if model == nil {
		return nil
	}

	var logDetails dto.RequestLogDetails
	if err := model.LogDetails.ToStruct(&logDetails); err != nil {
		utils.SysError("解析日志详情失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.RequestLog{
		RequestLogID:    model.RequestLogID,
		UserID:          model.UserID,
		ChannelID:       model.ChannelID,
		ModelID:         model.ModelID,
		TokenID:         model.TokenID,
		MasterID:        model.MasterID,
		RequestID:       model.RequestID,
		RequestType:     model.RequestType,
		RequestPath:     model.RequestPath,
		RequestMethod:   model.RequestMethod,
		RequestHeaders:  model.RequestHeaders,
		RequestParams:   model.RequestParams,
		RequestTokens:   model.RequestTokens,
		RequestTime:     model.RequestTime,
		RequestStatus:   model.RequestStatus,
		ResponseTime:    model.ResponseTime,
		TotalTime:       model.TotalTime,
		ResponseCode:    model.ResponseCode,
		ResponseHeaders: model.ResponseHeaders,
		ErrorMessage:    model.ErrorMessage,
		ClientIP:        model.ClientIP,
		EventType:       model.EventType,
		LogLevel:        model.LogLevel,
		LogDetails:      logDetails,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		DeletedAt:       deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *requestLogRepository) convertToModel(dto *dto.RequestLog) (*log.RequestLog, error) {
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

	return &log.RequestLog{
		RequestLogID:    dto.RequestLogID,
		UserID:          dto.UserID,
		ChannelID:       dto.ChannelID,
		ModelID:         dto.ModelID,
		TokenID:         dto.TokenID,
		MasterID:        dto.MasterID,
		RequestID:       dto.RequestID,
		RequestType:     dto.RequestType,
		RequestPath:     dto.RequestPath,
		RequestMethod:   dto.RequestMethod,
		RequestHeaders:  dto.RequestHeaders,
		RequestParams:   dto.RequestParams,
		RequestTokens:   dto.RequestTokens,
		RequestTime:     dto.RequestTime,
		RequestStatus:   dto.RequestStatus,
		ResponseTime:    dto.ResponseTime,
		TotalTime:       dto.TotalTime,
		ResponseCode:    dto.ResponseCode,
		ResponseHeaders: dto.ResponseHeaders,
		ErrorMessage:    dto.ErrorMessage,
		ClientIP:        dto.ClientIP,
		EventType:       dto.EventType,
		LogLevel:        dto.LogLevel,
		LogDetails:      logDetailsJSON,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
		DeletedAt:       deletedAt,
	}, nil
}

// Create 创建日志记录
func (r *requestLogRepository) Create(requestLog *dto.RequestLog) error {
	model, err := r.convertToModel(requestLog)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新日志记录
func (r *requestLogRepository) Update(requestLog *dto.RequestLog) error {
	model, err := r.convertToModel(requestLog)
	if err != nil {
		return err
	}
	return r.db.Model(&log.RequestLog{}).Where("request_log_id = ?", requestLog.RequestLogID).Updates(model).Error
}

// Delete 删除日志记录
func (r *requestLogRepository) Delete(requestLogID string) error {
	return r.db.Delete(&log.RequestLog{}, "request_log_id = ?", requestLogID).Error
}

// GetByID 根据ID获取日志记录
func (r *requestLogRepository) GetByID(requestLogID string) (*dto.RequestLog, error) {
	var requestLog log.RequestLog
	if err := r.db.Where("request_log_id = ?", requestLogID).First(&requestLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&requestLog), nil
}

// GetByRequestID 根据请求ID获取日志记录
func (r *requestLogRepository) GetByRequestID(requestID string) (*dto.RequestLog, error) {
	var requestLog log.RequestLog
	if err := r.db.Where("request_id = ?", requestID).First(&requestLog).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&requestLog), nil
}

// List 获取日志记录列表
func (r *requestLogRepository) List(page, pageSize int) ([]*dto.RequestLog, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var logs []log.RequestLog

	offset := (page - 1) * pageSize

	if err := r.db.Model(&log.RequestLog{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count request logs failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("find request logs failed: %w", err)
	}

	dtoList := make([]*dto.RequestLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByUser 获取指定用户的日志记录列表
func (r *requestLogRepository) ListByUser(userID string, page, pageSize int) ([]*dto.RequestLog, int64, error) {
	var total int64
	var logs []log.RequestLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RequestLog{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RequestLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByLevel 获取指定级别的日志记录列表
func (r *requestLogRepository) ListByLevel(logLevel string, page, pageSize int) ([]*dto.RequestLog, int64, error) {
	var total int64
	var logs []log.RequestLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RequestLog{}).Where("log_level = ?", logLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RequestLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// ListByEventType 获取指定事件类型的日志记录列表
func (r *requestLogRepository) ListByEventType(eventType string, page, pageSize int) ([]*dto.RequestLog, int64, error) {
	var total int64
	var logs []log.RequestLog

	offset := (page - 1) * pageSize
	query := r.db.Model(&log.RequestLog{}).Where("event_type = ?", eventType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.RequestLog, len(logs))
	for i, l := range logs {
		dtoList[i] = r.convertToDTO(&l)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *requestLogRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行请求日志基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testLog := &dto.RequestLog{
			RequestLogID:    utils.GenerateRandomUUID(12),
			UserID:          utils.GenerateRandomUUID(12),
			ChannelID:       utils.GenerateRandomUUID(12),
			ModelID:         utils.GenerateRandomUUID(12),
			TokenID:         utils.GenerateRandomUUID(12),
			MasterID:        utils.GenerateRandomUUID(12),
			RequestID:       fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			RequestType:     []string{"http", "rpc"}[rand.Intn(2)],
			RequestPath:     fmt.Sprintf("/api/test/%d", i),
			RequestMethod:   []string{"GET", "POST", "PUT", "DELETE"}[rand.Intn(4)],
			RequestHeaders:  common.JSON{},
			RequestParams:   common.JSON{},
			RequestTokens:   rand.Intn(100),
			RequestTime:     utils.MySQLTime(time.Now()),
			RequestStatus:   []int8{1, 2, 3}[rand.Intn(3)],
			ResponseTime:    utils.MySQLTime(time.Now().Add(time.Duration(rand.Intn(5000)) * time.Millisecond)),
			TotalTime:       rand.Intn(5000),
			ResponseCode:    []int{200, 400, 401, 403, 404, 500}[rand.Intn(6)],
			ResponseHeaders: common.JSON{},
			ErrorMessage:    fmt.Sprintf("测试错误消息 %d", i),
			ClientIP:        fmt.Sprintf("192.168.1.%d", rand.Intn(255)),
			EventType:       []string{"request", "response"}[rand.Intn(2)],
			LogLevel:        []string{"info", "warn", "error"}[rand.Intn(3)],
			LogDetails: dto.RequestLogDetails{
				Content: fmt.Sprintf("test_value_%d", i),
			},
		}

		// 创建
		if err := r.Create(testLog); err != nil {
			utils.SysError("创建日志记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdLog, err := r.GetByID(testLog.RequestLogID)
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
		if err := r.Delete(createdLog.RequestLogID); err != nil {
			utils.SysError("删除日志记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
