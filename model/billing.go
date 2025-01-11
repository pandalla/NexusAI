package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Billing 账单表
type Billing struct {
	BillingID       uint64         `gorm:"column:billing_id;primaryKey;autoIncrement" json:"billing_id"`
	UserID          uint64         `gorm:"column:user_id;index;not null" json:"user_id"`
	BillingNo       string         `gorm:"column:billing_no;size:64;uniqueIndex;not null" json:"billing_no"`
	BillingType     string         `gorm:"column:billing_type;size:20;index;not null" json:"billing_type"`
	BillingCycle    string         `gorm:"column:billing_cycle;size:20;not null" json:"billing_cycle"`
	BillingMoney    float64        `gorm:"column:billing_money;type:decimal(10,2);not null" json:"billing_money"`
	BillingCurrency string         `gorm:"column:billing_currency;size:10;not null" json:"billing_currency"`
	BillingDetails  common.JSON    `gorm:"column:billing_details;type:json;not null" json:"billing_details"`
	QuotaDetails    common.JSON    `gorm:"column:quota_details;type:json;not null" json:"quota_details"`
	BillingStatus   int8           `gorm:"column:billing_status;index;not null" json:"billing_status"`
	StartTime       time.Time      `gorm:"column:start_time" json:"start_time"`
	EndTime         time.Time      `gorm:"column:end_time" json:"end_time"`
	DueTime         time.Time      `gorm:"column:due_time;not null" json:"due_time"`
	PayTime         *time.Time     `gorm:"column:pay_time" json:"pay_time"`
	Remark          string         `gorm:"column:remark;size:255" json:"remark"`
	CreatedAt       time.Time      `gorm:"column:created_at;index;not null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Billing) TableName() string {
	return "billings"
}
