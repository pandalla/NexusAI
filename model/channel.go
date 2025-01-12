package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 系统中所有AI服务渠道的配置信息，包括上游服务配置、认证信息、重试策略等
type Channel struct {
	ChannelID          string `gorm:"column:channel_id;type:char(36);primaryKey;default:(UUID())" json:"channel_id"` // 渠道唯一标识
	ChannelGroupID     string `gorm:"column:channel_group_id;type:char(36);not null;index" json:"channel_group_id"`  // 渠道组ID
	ChannelName        string `gorm:"column:channel_name;size:100;not null" json:"channel_name"`                     // 渠道名称
	ChannelDescription string `gorm:"column:channel_description;type:text" json:"channel_description"`               // 渠道描述信息
	Status             int8   `gorm:"column:status;index;not null;default:1" json:"status"`                          // 渠道状态，1:正常 0:禁用

	ChannelModels      common.JSON `gorm:"column:channel_models;type:json" json:"channel_models"`             // 支持的模型列表配置
	ChannelPriceFactor common.JSON `gorm:"column:channel_price_factor;type:json" json:"channel_price_factor"` // 渠道价格系数
	UpstreamOptions    common.JSON `gorm:"column:upstream_options;type:json" json:"upstream_options"`         // 上游服务配置信息(endpoint/timeout等)
	AuthOptions        common.JSON `gorm:"column:auth_options;type:json" json:"auth_options"`                 // 认证配置信息(key/secret等)
	RetryOptions       common.JSON `gorm:"column:retry_options;type:json" json:"retry_options"`               // 重试策略配置
	RateLimit          common.JSON `gorm:"column:rate_limit;type:json" json:"rate_limit"`                     // 速率限制配置
	ModelMapping       common.JSON `gorm:"column:model_mapping;type:json" json:"model_mapping"`               // 模型映射配置，用于模型名称转换
	TestModels         common.JSON `gorm:"column:test_models;type:json" json:"test_models"`                   // 测试模型配置

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 渠道创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 渠道信息最后更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Channel) TableName() string {
	return "channels"
}
