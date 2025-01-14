package model

import (
	"nexus-ai/utils"
)

// BillingMoney 账单金额
type BillingMoney struct {
	TotalAmount    float64 `json:"total_amount"`    // 总金额
	DiscountAmount float64 `json:"discount_amount"` // 优惠金额
	ActualAmount   float64 `json:"actual_amount"`   // 实际金额
	PaidAmount     float64 `json:"paid_amount"`     // 已支付金额
	UnpaidAmount   float64 `json:"unpaid_amount"`   // 未支付金额
	RefundAmount   float64 `json:"refund_amount"`   // 退款金额
	OverdueAmount  float64 `json:"overdue_amount"`  // 逾期金额
	OverduePenalty float64 `json:"overdue_penalty"` // 逾期罚金
}

// BillingDetails 账单明细
type BillingDetails struct {
	ModelUsage     map[string]float64 `json:"model_usage"`     // 各模型用量
	ChannelUsage   map[string]float64 `json:"channel_usage"`   // 各渠道用量
	TokenUsage     map[string]int     `json:"token_usage"`     // Token用量
	RequestCount   map[string]int     `json:"request_count"`   // 请求次数
	FailureCount   map[string]int     `json:"failure_count"`   // 失败次数
	AverageLatency map[string]float64 `json:"average_latency"` // 平均延迟
}

// QuotaDetails 配额使用明细
type QuotaDetails struct {
	QuotaUsage     map[string]float64 `json:"quota_usage"`     // 配额使用情况
	QuotaRemaining map[string]float64 `json:"quota_remaining"` // 配额剩余情况
	QuotaFrozen    map[string]float64 `json:"quota_frozen"`    // 配额冻结情况
	QuotaExpired   map[string]float64 `json:"quota_expired"`   // 配额过期情况
}

// Billing DTO结构
type Billing struct {
	BillingID       string           `json:"billing_id"`           // 账单ID
	UserID          string           `json:"user_id"`              // 用户ID
	BillingNo       string           `json:"billing_no"`           // 账单编号
	BillingType     string           `json:"billing_type"`         // 账单类型
	BillingCycle    string           `json:"billing_cycle"`        // 账单周期
	Remark          string           `json:"remark"`               // 账单备注
	BillingCurrency string           `json:"billing_currency"`     // 账单币种
	BillingStatus   int8             `json:"billing_status"`       // 账单状态
	BillingMoney    BillingMoney     `json:"billing_money"`        // 账单金额
	BillingDetails  BillingDetails   `json:"billing_details"`      // 账单明细
	QuotaDetails    QuotaDetails     `json:"quota_details"`        // 配额使用明细
	PayTime         utils.MySQLTime  `json:"pay_time"`             // 支付时间
	StartTime       utils.MySQLTime  `json:"start_time"`           // 账单开始时间
	EndTime         utils.MySQLTime  `json:"end_time"`             // 账单结束时间
	DueTime         utils.MySQLTime  `json:"due_time"`             // 账单到期时间
	CreatedAt       utils.MySQLTime  `json:"created_at"`           // 创建时间
	UpdatedAt       utils.MySQLTime  `json:"updated_at"`           // 更新时间
	DeletedAt       *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
