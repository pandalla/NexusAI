package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Model 存储系统支持的AI模型信息，包括定价、配置和服务提供商等信息
type Model struct {
	// 模型唯一标识
	ModelID string `gorm:"column:model_id;type:char(36);primaryKey;default:(UUID())" json:"model_id"`

	// 模型名称，最大长度100字符
	ModelName string `gorm:"column:model_name;size:100;uniqueIndex;not null" json:"model_name"`

	// 模型别名映射配置
	ModelAlias common.JSON `gorm:"column:model_alias;type:json" json:"model_alias"`

	// 模型单价，支持6位小数
	ModelPrice float64 `gorm:"column:model_price;type:decimal(10,6);not null;default:0.000000" json:"model_price"`

	// 计费类型(usage:按量,times:按次)
	PriceType string `gorm:"column:price_type;size:20;not null" json:"price_type"`

	// 模型描述信息
	ModelDescription string `gorm:"column:model_description;type:text" json:"model_description"`

	// 模型配置信息(频率限制/并发数等)
	ModelConfig common.JSON `gorm:"column:model_config;type:json" json:"model_config"`

	// 模型类型(如text/image/audio/video等)
	ModelType string `gorm:"column:model_type;size:50;index;not null" json:"model_type"`

	// 服务提供商
	Provider string `gorm:"column:provider;size:50;index;not null" json:"provider"`

	// 模型状态，1:正常 0:禁用
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 模型创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 模型信息最后更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Model) TableName() string {
	return "models"
}
