package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Billing 存储用户的账单信息，包括计费周期、金额和支付状态等
type Billing struct {
	// 账单唯一标识
	BillingID string `gorm:"column:billing_id;type:char(36);primaryKey;default:(UUID())" json:"billing_id"`

	// 关联的用户ID
	UserID string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`

	// 账单编号，系统生成的唯一编号
	BillingNo string `gorm:"column:billing_no;size:64;uniqueIndex;not null" json:"billing_no"`

	// 账单类型(消费/退款)
	BillingType string `gorm:"column:billing_type;size:20;index;not null" json:"billing_type"`

	// 账单周期(一次性/日/周/月)
	BillingCycle string `gorm:"column:billing_cycle;size:20;not null" json:"billing_cycle"`

	// 账单金额
	BillingMoney float64 `gorm:"column:billing_money;type:decimal(10,2);not null;default:0.00" json:"billing_money"`

	// 账单币种
	BillingCurrency string `gorm:"column:billing_currency;size:10;not null" json:"billing_currency"`

	// 账单明细(各模型用量统计)
	BillingDetails common.JSON `gorm:"column:billing_details;type:json;not null" json:"billing_details"`

	// 配额使用明细
	QuotaDetails common.JSON `gorm:"column:quota_details;type:json;not null" json:"quota_details"`

	// 账单状态(1:未出账 2:已出账 3:已支付 4:已逾期)
	BillingStatus int8 `gorm:"column:billing_status;index;not null;default:1" json:"billing_status"`

	// 账单开始时间
	StartTime time.Time `gorm:"column:start_time;not null" json:"start_time"`

	// 账单结束时间
	EndTime time.Time `gorm:"column:end_time;not null" json:"end_time"`

	// 账单到期时间
	DueTime time.Time `gorm:"column:due_time;not null" json:"due_time"`

	// 支付时间
	PayTime *time.Time `gorm:"column:pay_time" json:"pay_time"`

	// 账单备注
	Remark string `gorm:"column:remark;size:255" json:"remark"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Billing) TableName() string {
	return "billings"
}
