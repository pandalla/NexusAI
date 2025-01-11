package model

import (
	"time"
)

// Usage 用量表
type Usage struct {
	UsageID     uint64    `gorm:"column:usage_id;primaryKey;autoIncrement" json:"usage_id"`
	TokenID     uint64    `gorm:"column:token_id;index;not null" json:"token_id"`
	UserID      uint64    `gorm:"column:user_id;index;not null" json:"user_id"`
	ChannelID   uint64    `gorm:"column:channel_id;index;not null" json:"channel_id"`
	ModelID     uint64    `gorm:"column:model_id;index;not null" json:"model_id"`
	TokensCount int       `gorm:"column:tokens_count" json:"tokens_count"`
	TimesCount  int       `gorm:"column:times_count" json:"times_count"`
	UnitPrice   float64   `gorm:"column:unit_price;type:decimal(10,6);not null" json:"unit_price"`
	PriceFactor float64   `gorm:"column:price_factor;type:decimal(10,2);not null" json:"price_factor"`
	TotalAmount float64   `gorm:"column:total_amount;type:decimal(10,6);not null" json:"total_amount"`
	UsageType   string    `gorm:"column:usage_type;size:20;index;not null" json:"usage_type"`
	RequestID   string    `gorm:"column:request_id;size:64" json:"request_id"`
	CreatedAt   time.Time `gorm:"column:created_at;index;not null" json:"created_at"`
}

// TableName 表名
func (Usage) TableName() string {
	return "usages"
}
