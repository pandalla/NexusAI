package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Model 模型表
type Model struct {
	ModelID          uint64         `gorm:"column:model_id;primaryKey;autoIncrement" json:"model_id"`
	ModelName        string         `gorm:"column:model_name;size:100;uniqueIndex;not null" json:"model_name"`
	ModelAlias       common.JSON    `gorm:"column:model_alias;type:json" json:"model_alias"`
	ModelPrice       float64        `gorm:"column:model_price;type:decimal(10,6);not null" json:"model_price"`
	PriceType        string         `gorm:"column:price_type;size:20;not null" json:"price_type"`
	ModelDescription string         `gorm:"column:model_description;type:text" json:"model_description"`
	ModelConfig      common.JSON    `gorm:"column:model_config;type:json" json:"model_config"`
	ModelType        string         `gorm:"column:model_type;size:50;index;not null" json:"model_type"`
	Provider         string         `gorm:"column:provider;size:50;not null" json:"provider"`
	Status           int8           `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt        time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Model) TableName() string {
	return "models"
}
