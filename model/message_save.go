package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 存储系统中所有AI对话消息记录
type MessageSave struct {
	// 消息唯一标识
	MessageID string `gorm:"column:message_id;type:char(36);primaryKey;default:(UUID())" json:"message_id"`

	// 关联的用户ID
	UserID string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`

	// 关联的令牌ID
	TokenID string `gorm:"column:token_id;type:char(36);index;not null;foreignKey:Token(TokenID)" json:"token_id"`

	// 关联的模型ID
	ModelID string `gorm:"column:model_id;type:char(36);index;not null;foreignKey:Model(ModelID)" json:"model_id"`

	// 关联的渠道ID
	ChannelID string `gorm:"column:channel_id;type:char(36);index;not null;foreignKey:Channel(ChannelID)" json:"channel_id"`

	// 消息标题
	MessageTitle string `gorm:"column:message_title;size:255" json:"message_title"`

	// 消息内容
	MessageContent string `gorm:"column:message_content;type:text;not null" json:"message_content"`

	// 消息token数
	MessageTokens int `gorm:"column:message_tokens;not null;default:0" json:"message_tokens"`

	// 提示词token数
	PromptTokens int `gorm:"column:prompt_tokens;not null;default:0" json:"prompt_tokens"`

	// 补全token数
	CompletionTokens int `gorm:"column:completion_tokens;not null;default:0" json:"completion_tokens"`

	// 消息状态(1:正常 2:中断 3:异常)
	MessageStatus int8 `gorm:"column:message_status;index;not null;default:1" json:"message_status"`

	// 用量信息
	UsageInfo common.JSON `gorm:"column:usage_info;type:json" json:"usage_info"`

	// 消息配置(温度/top_p等)
	MessageConfig common.JSON `gorm:"column:message_config;type:json" json:"message_config"`

	// 提示词模板
	PromptTemplate common.JSON `gorm:"column:prompt_template;type:json" json:"prompt_template"`

	// 消息额外信息
	MessageExtra common.JSON `gorm:"column:message_extra;type:json" json:"message_extra"`

	// 错误类型
	ErrorType string `gorm:"column:error_type;size:50" json:"error_type"`

	// 错误信息
	ErrorMessage string `gorm:"column:error_message;type:text" json:"error_message"`

	// 重试次数
	RetryCount int `gorm:"column:retry_count;not null;default:0" json:"retry_count"`

	// 响应延迟(ms)
	Latency int `gorm:"column:latency;not null;default:0" json:"latency"`

	// 关联的请求ID，用于日志追踪
	RequestID string `gorm:"column:request_id;size:64;index" json:"request_id"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (MessageSave) TableName() string {
	return "message_saves"
}
