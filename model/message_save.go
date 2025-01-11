package model

import (
	"time"

	"nexus-ai/common"
)

// MessageSave 消息存储表
type MessageSave struct {
	MessageID        uint64      `gorm:"column:message_id;primaryKey;autoIncrement" json:"message_id"`
	UserID           uint64      `gorm:"column:user_id;index;not null" json:"user_id"`
	TokenID          uint64      `gorm:"column:token_id;index;not null" json:"token_id"`
	ModelID          uint64      `gorm:"column:model_id;index;not null" json:"model_id"`
	ChannelID        uint64      `gorm:"column:channel_id;index;not null" json:"channel_id"`
	ConversationID   string      `gorm:"column:conversation_id;size:64;index" json:"conversation_id"`
	ParentMessageID  string      `gorm:"column:parent_message_id;size:64;index" json:"parent_message_id"`
	MessageType      string      `gorm:"column:message_type;size:20;index;not null" json:"message_type"`
	MessageRole      string      `gorm:"column:message_role;size:20;not null" json:"message_role"`
	BusinessType     string      `gorm:"column:business_type;size:50;index" json:"business_type"`
	IndustryType     string      `gorm:"column:industry_type;size:50;index" json:"industry_type"`
	SceneType        string      `gorm:"column:scene_type;size:50;index" json:"scene_type"`
	MessageTitle     string      `gorm:"column:message_title;size:255" json:"message_title"`
	MessageContent   string      `gorm:"column:message_content;type:text;not null" json:"message_content"`
	MessageTokens    int         `gorm:"column:message_tokens" json:"message_tokens"`
	PromptTokens     int         `gorm:"column:prompt_tokens" json:"prompt_tokens"`
	CompletionTokens int         `gorm:"column:completion_tokens" json:"completion_tokens"`
	MessageStatus    int8        `gorm:"column:message_status;index;not null" json:"message_status"`
	QualityScore     float64     `gorm:"column:quality_score;type:decimal(3,2)" json:"quality_score"`
	RelevanceScore   float64     `gorm:"column:relevance_score;type:decimal(3,2)" json:"relevance_score"`
	UserFeedback     int8        `gorm:"column:user_feedback" json:"user_feedback"`
	UserComment      string      `gorm:"column:user_comment;size:500" json:"user_comment"`
	UsageInfo        common.JSON `gorm:"column:usage_info;type:json" json:"usage_info"`
	MessageConfig    common.JSON `gorm:"column:message_config;type:json" json:"message_config"`
	PromptTemplate   common.JSON `gorm:"column:prompt_template;type:json" json:"prompt_template"`
	Keywords         common.JSON `gorm:"column:keywords;type:json" json:"keywords"`
	Sentiment        string      `gorm:"column:sentiment;size:20" json:"sentiment"`
	IntentTags       common.JSON `gorm:"column:intent_tags;type:json" json:"intent_tags"`
	EntityTags       common.JSON `gorm:"column:entity_tags;type:json" json:"entity_tags"`
	Classification   common.JSON `gorm:"column:classification;type:json" json:"classification"`
	MessageExtra     common.JSON `gorm:"column:message_extra;type:json" json:"message_extra"`
	MessageTags      common.JSON `gorm:"column:message_tags;type:json" json:"message_tags"`
	IsSensitive      int8        `gorm:"column:is_sensitive;not null;default:0" json:"is_sensitive"`
	IsPublic         int8        `gorm:"column:is_public;not null;default:0" json:"is_public"`
	ShareCode        string      `gorm:"column:share_code;size:32;uniqueIndex" json:"share_code"`
	ErrorType        string      `gorm:"column:error_type;size:50" json:"error_type"`
	ErrorMessage     string      `gorm:"column:error_message;type:text" json:"error_message"`
	RetryCount       int         `gorm:"column:retry_count" json:"retry_count"`
	Latency          int         `gorm:"column:latency" json:"latency"`
	RequestID        string      `gorm:"column:request_id;size:64;index" json:"request_id"`
	CreatedAt        time.Time   `gorm:"column:created_at;index;not null" json:"created_at"`
	UpdatedAt        time.Time   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName 表名
func (MessageSave) TableName() string {
	return "message_saves"
}
