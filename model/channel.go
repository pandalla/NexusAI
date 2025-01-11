package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Channel 存储系统中所有AI服务渠道的配置信息，包括上游服务配置、认证信息、重试策略等
type Channel struct {
	// 渠道唯一标识
	ChannelID string `gorm:"column:channel_id;type:char(36);primaryKey;default:(UUID())" json:"channel_id"`

	// 渠道名称，最大长度100字符
	ChannelName string `gorm:"column:channel_name;size:100;not null" json:"channel_name"`

	// 渠道等级，用于分级管理，最大长度50字符
	ChannelLevel string `gorm:"column:channel_level;size:50;index;not null" json:"channel_level"`

	// 渠道描述信息
	ChannelDescription string `gorm:"column:channel_description;type:text" json:"channel_description"`

	// 支持的模型列表配置
	ChannelModels common.JSON `gorm:"column:channel_models;type:json" json:"channel_models"`

	// 价格系数，用于计算最终价格
	PriceFactor float64 `gorm:"column:price_factor;type:decimal(10,2);not null;default:1.0" json:"price_factor"`

	// 上游服务配置信息(endpoint/timeout等)
	UpstreamConfig common.JSON `gorm:"column:upstream_config;type:json;not null" json:"upstream_config"`

	// 认证配置信息(key/secret等)
	AuthConfig common.JSON `gorm:"column:auth_config;type:json;not null" json:"auth_config"`

	// 重试策略配置
	RetryConfig common.JSON `gorm:"column:retry_config;type:json" json:"retry_config"`

	// 速率限制配置
	RateLimit common.JSON `gorm:"column:rate_limit;type:json" json:"rate_limit"`

	// 模型映射配置，用于模型名称转换
	ModelMapping common.JSON `gorm:"column:model_mapping;type:json" json:"model_mapping"`

	// 测试模型配置
	TestModels common.JSON `gorm:"column:test_models;type:json" json:"test_models"`

	// 渠道状态，1:正常 0:禁用
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 渠道创建时间
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 渠道信息最后更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Channel) TableName() string {
	return "channels"
}
