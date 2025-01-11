package model

import (
	"time"

	"nexus-ai/common"
)

// Payment 支付表
type Payment struct {
	PayID          uint64      `gorm:"column:pay_id;primaryKey;autoIncrement" json:"pay_id"`
	UserID         uint64      `gorm:"column:user_id;index;not null" json:"user_id"`
	OrderNo        string      `gorm:"column:order_no;size:64;uniqueIndex;not null" json:"order_no"`
	PayPlatform    string      `gorm:"column:pay_platform;size:50;index;not null" json:"pay_platform"`
	PayMethod      string      `gorm:"column:pay_method;size:50;not null" json:"pay_method"`
	PayType        string      `gorm:"column:pay_type;size:20;index;not null" json:"pay_type"`
	PayScene       string      `gorm:"column:pay_scene;size:20;not null" json:"pay_scene"`
	PayCurrency    string      `gorm:"column:pay_currency;size:10;not null" json:"pay_currency"`
	PayAmount      float64     `gorm:"column:pay_amount;type:decimal(10,2);not null" json:"pay_amount"`
	PayStatus      int8        `gorm:"column:pay_status;index;not null" json:"pay_status"`
	PayTitle       string      `gorm:"column:pay_title;size:255;not null" json:"pay_title"`
	PayDesc        string      `gorm:"column:pay_desc;size:1000" json:"pay_desc"`
	PayerName      string      `gorm:"column:payer_name;size:100" json:"payer_name"`
	PayerEmail     string      `gorm:"column:payer_email;size:100" json:"payer_email"`
	PayerPhone     string      `gorm:"column:payer_phone;size:20" json:"payer_phone"`
	CompanyInfo    common.JSON `gorm:"column:company_info;type:json" json:"company_info"`
	BillingInfo    common.JSON `gorm:"column:billing_info;type:json" json:"billing_info"`
	NotifyURL      string      `gorm:"column:notify_url;size:255" json:"notify_url"`
	TransactionID  string      `gorm:"column:transaction_id;size:64" json:"transaction_id"`
	CallbackData   common.JSON `gorm:"column:callback_data;type:json" json:"callback_data"`
	PlatformConfig common.JSON `gorm:"column:platform_config;type:json" json:"platform_config"`
	RefundInfo     common.JSON `gorm:"column:refund_info;type:json" json:"refund_info"`
	ExpireTime     time.Time   `gorm:"column:expire_time;not null" json:"expire_time"`
	PayTime        *time.Time  `gorm:"column:pay_time" json:"pay_time"`
	CreatedAt      time.Time   `gorm:"column:created_at;index;not null" json:"created_at"`
	UpdatedAt      time.Time   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName 表名
func (Payment) TableName() string {
	return "pays"
}
