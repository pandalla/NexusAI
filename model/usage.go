package model

import (
	"time"
)

// Usage 存储系统中所有API调用的用量记录，包括token使用量、计费信息等
type Usage struct {
	// 用量记录唯一标识
	UsageID string `gorm:"column:usage_id;type:char(36);primaryKey;default:(UUID())" json:"usage_id"`

	// 关联的令牌ID
	TokenID string `gorm:"column:token_id;type:char(36);index;not null;foreignKey:Token(TokenID)" json:"token_id"`

	// 关联的用户ID
	UserID string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`

	// 关联的渠道ID
	ChannelID string `gorm:"column:channel_id;type:char(36);index;not null;foreignKey:Channel(ChannelID)" json:"channel_id"`

	// 关联的模型ID
	ModelID string `gorm:"column:model_id;type:char(36);index;not null;foreignKey:Model(ModelID)" json:"model_id"`

	// Token数量(按量收费时使用)
	TokensCount int `gorm:"column:tokens_count;not null;default:0" json:"tokens_count"`

	// 请求次数(按次收费时使用)
	TimesCount int `gorm:"column:times_count;not null;default:0" json:"times_count"`

	// 单价，支持6位小数
	UnitPrice float64 `gorm:"column:unit_price;type:decimal(10,6);not null;default:0.000000" json:"unit_price"`

	// 价格倍率
	PriceFactor float64 `gorm:"column:price_factor;type:decimal(10,2);not null;default:1.00" json:"price_factor"`

	// 总金额，支持6位小数
	TotalAmount float64 `gorm:"column:total_amount;type:decimal(10,6);not null;default:0.000000" json:"total_amount"`

	// 计费类型(usage:按量,times:按次)
	UsageType string `gorm:"column:usage_type;size:20;index;not null" json:"usage_type"`

	// 关联的请求ID，用于日志追踪
	RequestID string `gorm:"column:request_id;size:64;index" json:"request_id"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`
}

// TableName 表名
func (Usage) TableName() string {
	return "usages"
}
