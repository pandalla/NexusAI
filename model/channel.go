package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Channel 渠道表
type Channel struct {
	ChannelID          uint64         `gorm:"column:channel_id;primaryKey;autoIncrement" json:"channel_id"`
	ChannelName        string         `gorm:"column:channel_name;size:100;not null" json:"channel_name"`
	ChannelLevel       string         `gorm:"column:channel_level;size:50;index;not null" json:"channel_level"`
	ChannelDescription string         `gorm:"column:channel_description;type:text" json:"channel_description"`
	ChannelModels      common.JSON    `gorm:"column:channel_models;type:json" json:"channel_models"`
	PriceFactor        float64        `gorm:"column:price_factor;type:decimal(10,2);not null" json:"price_factor"`
	UpstreamConfig     common.JSON    `gorm:"column:upstream_config;type:json;not null" json:"upstream_config"`
	AuthConfig         common.JSON    `gorm:"column:auth_config;type:json;not null" json:"auth_config"`
	RetryConfig        common.JSON    `gorm:"column:retry_config;type:json" json:"retry_config"`
	RateLimit          common.JSON    `gorm:"column:rate_limit;type:json" json:"rate_limit"`
	ModelMapping       common.JSON    `gorm:"column:model_mapping;type:json" json:"model_mapping"`
	TestModels         common.JSON    `gorm:"column:test_models;type:json" json:"test_models"`
	Status             int8           `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt          time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Channel) TableName() string {
	return "channels"
}
