package model

import (
	"time"

	"nexus-ai/common"
)

// Quota 配额表
type Quota struct {
	QuotaID         uint64      `gorm:"column:quota_id;primaryKey;autoIncrement" json:"quota_id"`
	UserID          uint64      `gorm:"column:user_id;index;not null" json:"user_id"`
	PayID           uint64      `gorm:"column:pay_id;index;not null" json:"pay_id"`
	QuotaAmount     float64     `gorm:"column:quota_amount;type:decimal(10,2);not null" json:"quota_amount"`
	RemainingAmount float64     `gorm:"column:remaining_amount;type:decimal(10,2);not null" json:"remaining_amount"`
	FrozenAmount    float64     `gorm:"column:frozen_amount;type:decimal(10,2);not null" json:"frozen_amount"`
	QuotaType       string      `gorm:"column:quota_type;size:20;index;not null" json:"quota_type"`
	QuotaLevel      string      `gorm:"column:quota_level;size:20;index;not null" json:"quota_level"`
	ValidPeriod     int         `gorm:"column:valid_period;not null" json:"valid_period"`
	StartTime       time.Time   `gorm:"column:start_time;not null" json:"start_time"`
	ExpireTime      time.Time   `gorm:"column:expire_time;not null" json:"expire_time"`
	QuotaConfig     common.JSON `gorm:"column:quota_config;type:json" json:"quota_config"`
	Status          int8        `gorm:"column:status;index;not null" json:"status"`
	CreatedAt       time.Time   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt       time.Time   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName 表名
func (Quota) TableName() string {
	return "quotas"
}
