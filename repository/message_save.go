package repository

import (
	"fmt"
	"math/rand"
	"nexus-ai/common"
	dto "nexus-ai/dto/model"
	"nexus-ai/model"
	"nexus-ai/utils"
	"time"

	"gorm.io/gorm"
)

// MessageSaveRepository 消息存储仓储接口
type MessageSaveRepository interface {
	Create(message *dto.MessageSave) error
	Update(message *dto.MessageSave) error
	Delete(messageID string) error
	GetByID(messageID string) (*dto.MessageSave, error)
	GetByRequestID(requestID string) (*dto.MessageSave, error)
	List(page, pageSize int) ([]*dto.MessageSave, int64, error)
	ListByUser(userID string, page, pageSize int) ([]*dto.MessageSave, int64, error)
	ListByToken(tokenID string, page, pageSize int) ([]*dto.MessageSave, int64, error)
	ListByModel(modelID string, page, pageSize int) ([]*dto.MessageSave, int64, error)
	ListByChannel(channelID string, page, pageSize int) ([]*dto.MessageSave, int64, error)
	ListByStatus(status int8, page, pageSize int) ([]*dto.MessageSave, int64, error)
	Benchmark(count int) error
}

type messageSaveRepository struct {
	db *gorm.DB
}

// NewMessageSaveRepository 创建消息存储仓储实例
func NewMessageSaveRepository(db *gorm.DB) MessageSaveRepository {
	return &messageSaveRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *messageSaveRepository) convertToDTO(model *model.MessageSave) *dto.MessageSave {
	if model == nil {
		return nil
	}

	var messageOptions dto.MessageOptions
	var promptTemplate dto.PromptTemplate
	var messageExtra dto.MessageExtra

	if err := model.MessageOptions.ToStruct(&messageOptions); err != nil {
		utils.SysError("解析消息配置失败:" + err.Error())
	}

	if err := model.PromptTemplate.ToStruct(&promptTemplate); err != nil {
		utils.SysError("解析提示词模板失败:" + err.Error())
	}

	if err := model.MessageExtra.ToStruct(&messageExtra); err != nil {
		utils.SysError("解析消息额外信息失败:" + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.MessageSave{
		MessageID:        model.MessageID,
		UserID:           model.UserID,
		TokenID:          model.TokenID,
		ModelID:          model.ModelID,
		ChannelID:        model.ChannelID,
		RequestID:        model.RequestID,
		MessageTitle:     model.MessageTitle,
		MessageContent:   model.MessageContent,
		MessageTokens:    model.MessageTokens,
		PromptTokens:     model.PromptTokens,
		CompletionTokens: model.CompletionTokens,
		Latency:          model.Latency,
		MessageStatus:    model.MessageStatus,
		RetryCount:       model.RetryCount,
		ErrorType:        model.ErrorType,
		ErrorInfo:        model.ErrorInfo,
		MessageOptions:   messageOptions,
		PromptTemplate:   promptTemplate,
		MessageExtra:     messageExtra,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *messageSaveRepository) convertToModel(dto *dto.MessageSave) (*model.MessageSave, error) {
	if dto == nil {
		return nil, nil
	}

	messageOptionsJSON, err := common.FromStruct(dto.MessageOptions)
	if err != nil {
		return nil, fmt.Errorf("转换消息配置失败: %w", err)
	}

	promptTemplateJSON, err := common.FromStruct(dto.PromptTemplate)
	if err != nil {
		return nil, fmt.Errorf("转换提示词模板失败: %w", err)
	}

	messageExtraJSON, err := common.FromStruct(dto.MessageExtra)
	if err != nil {
		return nil, fmt.Errorf("转换消息额外信息失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &model.MessageSave{
		MessageID:        dto.MessageID,
		UserID:           dto.UserID,
		TokenID:          dto.TokenID,
		ModelID:          dto.ModelID,
		ChannelID:        dto.ChannelID,
		RequestID:        dto.RequestID,
		MessageTitle:     dto.MessageTitle,
		MessageContent:   dto.MessageContent,
		MessageTokens:    dto.MessageTokens,
		PromptTokens:     dto.PromptTokens,
		CompletionTokens: dto.CompletionTokens,
		Latency:          dto.Latency,
		MessageStatus:    dto.MessageStatus,
		RetryCount:       dto.RetryCount,
		ErrorType:        dto.ErrorType,
		ErrorInfo:        dto.ErrorInfo,
		MessageOptions:   messageOptionsJSON,
		PromptTemplate:   promptTemplateJSON,
		MessageExtra:     messageExtraJSON,
		CreatedAt:        dto.CreatedAt,
		UpdatedAt:        dto.UpdatedAt,
		DeletedAt:        deletedAt,
	}, nil
}

// Create 创建消息记录
func (r *messageSaveRepository) Create(message *dto.MessageSave) error {
	model, err := r.convertToModel(message)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新消息记录
func (r *messageSaveRepository) Update(message *dto.MessageSave) error {
	modelData, err := r.convertToModel(message)
	if err != nil {
		return err
	}
	return r.db.Model(&model.MessageSave{}).Where("message_id = ?", message.MessageID).Updates(modelData).Error
}

// Delete 删除消息记录
func (r *messageSaveRepository) Delete(messageID string) error {
	return r.db.Delete(&model.MessageSave{}, "message_id = ?", messageID).Error
}

// GetByID 根据ID获取消息记录
func (r *messageSaveRepository) GetByID(messageID string) (*dto.MessageSave, error) {
	var message model.MessageSave
	if err := r.db.Where("message_id = ?", messageID).First(&message).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&message), nil
}

// GetByRequestID 根据请求ID获取消息记录
func (r *messageSaveRepository) GetByRequestID(requestID string) (*dto.MessageSave, error) {
	var message model.MessageSave
	if err := r.db.Where("request_id = ?", requestID).First(&message).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&message), nil
}

// List 获取消息记录列表
func (r *messageSaveRepository) List(page, pageSize int) ([]*dto.MessageSave, int64, error) {
	var total int64
	var messages []model.MessageSave

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.MessageSave{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MessageSave, len(messages))
	for i, m := range messages {
		dtoList[i] = r.convertToDTO(&m)
	}

	return dtoList, total, nil
}

// ListByUser 获取用户的消息记录列表
func (r *messageSaveRepository) ListByUser(userID string, page, pageSize int) ([]*dto.MessageSave, int64, error) {
	var total int64
	var messages []model.MessageSave

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.MessageSave{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MessageSave, len(messages))
	for i, m := range messages {
		dtoList[i] = r.convertToDTO(&m)
	}

	return dtoList, total, nil
}

// ListByToken 获取令牌的消息记录列表
func (r *messageSaveRepository) ListByToken(tokenID string, page, pageSize int) ([]*dto.MessageSave, int64, error) {
	var total int64
	var messages []model.MessageSave

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.MessageSave{}).Where("token_id = ?", tokenID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MessageSave, len(messages))
	for i, m := range messages {
		dtoList[i] = r.convertToDTO(&m)
	}

	return dtoList, total, nil
}

// ListByModel 获取模型的消息记录列表
func (r *messageSaveRepository) ListByModel(modelID string, page, pageSize int) ([]*dto.MessageSave, int64, error) {
	var total int64
	var messages []model.MessageSave

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.MessageSave{}).Where("model_id = ?", modelID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MessageSave, len(messages))
	for i, m := range messages {
		dtoList[i] = r.convertToDTO(&m)
	}

	return dtoList, total, nil
}

// ListByChannel 获取渠道的消息记录列表
func (r *messageSaveRepository) ListByChannel(channelID string, page, pageSize int) ([]*dto.MessageSave, int64, error) {
	var total int64
	var messages []model.MessageSave

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.MessageSave{}).Where("channel_id = ?", channelID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MessageSave, len(messages))
	for i, m := range messages {
		dtoList[i] = r.convertToDTO(&m)
	}

	return dtoList, total, nil
}

// ListByStatus 获取指定状态的消息记录列表
func (r *messageSaveRepository) ListByStatus(status int8, page, pageSize int) ([]*dto.MessageSave, int64, error) {
	var total int64
	var messages []model.MessageSave

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.MessageSave{}).Where("message_status = ?", status)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.MessageSave, len(messages))
	for i, m := range messages {
		dtoList[i] = r.convertToDTO(&m)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *messageSaveRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行消息记录基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testMessage := &dto.MessageSave{
			MessageID:        utils.GenerateRandomUUID(12),
			UserID:           utils.GenerateRandomUUID(12),
			TokenID:          utils.GenerateRandomUUID(12),
			ModelID:          utils.GenerateRandomUUID(12),
			ChannelID:        utils.GenerateRandomUUID(12),
			RequestID:        fmt.Sprintf("REQ_%s", utils.GenerateRandomString(16)),
			MessageTitle:     fmt.Sprintf("测试消息 %d", i),
			MessageContent:   fmt.Sprintf("这是一条测试消息内容 %d", i),
			MessageTokens:    rand.Intn(1000),
			PromptTokens:     rand.Intn(500),
			CompletionTokens: rand.Intn(500),
			Latency:          rand.Intn(2000),
			MessageStatus:    int8(rand.Intn(3) + 1),
			RetryCount:       rand.Intn(3),
			MessageOptions: dto.MessageOptions{
				Temperature: float64(rand.Intn(100)) / 100,
				TopP:        float64(rand.Intn(100)) / 100,
				MaxTokens:   rand.Intn(2000) + 1000,
				Stream:      rand.Intn(2) == 1,
			},
			PromptTemplate: dto.PromptTemplate{
				TemplateID:      utils.GenerateRandomString(8),
				TemplateName:    fmt.Sprintf("模板 %d", i),
				TemplateContent: fmt.Sprintf("模板内容 %d", i),
			},
			MessageExtra: dto.MessageExtra{
				Source:    []string{"web", "api", "sdk"}[rand.Intn(3)],
				IP:        fmt.Sprintf("192.168.1.%d", rand.Intn(255)),
				TraceID:   utils.GenerateRandomString(16),
				SessionID: utils.GenerateRandomString(16),
			},
		}

		// 创建
		if err := r.Create(testMessage); err != nil {
			utils.SysError("创建消息记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdMessage, err := r.GetByID(testMessage.MessageID)
		if err != nil {
			utils.SysError("获取创建的消息记录失败: " + err.Error())
			return err
		}

		// 更新
		createdMessage.MessageStatus = int8(rand.Intn(3) + 1)
		if err := r.Update(createdMessage); err != nil {
			utils.SysError("更新消息记录失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdMessage.MessageID); err != nil {
			utils.SysError("删除消息记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
