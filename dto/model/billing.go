package model

import (
	"nexus-ai/utils"
)

// BillingMoney 账单金额
type BillingMoney struct {
	TotalAmount               float64 `json:"total_amount"`                // 总金额
	DiscountAmount            float64 `json:"discount_amount"`             // 优惠金额
	ActualAmount              float64 `json:"actual_amount"`               // 实际金额
	PaidAmount                float64 `json:"paid_amount"`                 // 已支付金额
	UnpaidAmount              float64 `json:"unpaid_amount"`               // 未支付金额
	PaymentPlatformCommission float64 `json:"payment_platform_commission"` // 支付平台佣金
	PaymentPlatformFee        float64 `json:"payment_platform_fee"`        // 支付平台手续费
	RefundAmount              float64 `json:"refund_amount"`               // 退款金额
	OverdueAmount             float64 `json:"overdue_amount"`              // 逾期金额
	OverduePenalty            float64 `json:"overdue_penalty"`             // 逾期罚金
}

// BillingDetails 账单明细
type BillingDetails struct {
	ModelUsage      map[string]float64 `json:"model_usage"`       // 各模型总用量
	TextUsage       map[string]float64 `json:"text_usage"`        // 文本用量
	ImageUsage      map[string]float64 `json:"image_usage"`       // 图片用量
	AudioUsage      map[string]float64 `json:"audio_usage"`       // 音频用量
	VideoUsage      map[string]float64 `json:"video_usage"`       // 视频用量
	MultiMediaUsage map[string]float64 `json:"multi_media_usage"` // 多模态用量
	OtherUsage      map[string]float64 `json:"other_usage"`       // 其他用量
}

// QuotaDetails 配额使用明细
type QuotaDetails struct {
	QuotaID         string          `json:"quota_id"`          // 配额ID
	QuotaType       string          `json:"quota_type"`        // 配额类型
	QuotaAmount     float64         `json:"quota_amount"`      // 配额总金额
	QuotaUsage      float64         `json:"quota_usage"`       // 配额使用量
	QuotaFrozen     float64         `json:"quota_frozen"`      // 配额冻结量
	QuotaRemaining  float64         `json:"quota_remaining"`   // 配额剩余量
	QuotaExpired    float64         `json:"quota_expired"`     // 配额过期量
	QuotaExpireTime utils.MySQLTime `json:"quota_expire_time"` // 配额过期时间
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
