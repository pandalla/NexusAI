package model

import (
	"nexus-ai/utils"
)

// PaymentInfo 支付信息
type PaymentInfo struct {
	InvoiceTitle    string `json:"invoice_title"`     // 发票抬头
	InvoiceType     string `json:"invoice_type"`      // 发票类型
	TaxNumber       string `json:"tax_number"`        // 税号
	InvoiceContent  string `json:"invoice_content"`   // 发票内容
	InvoiceEmail    string `json:"invoice_email"`     // 发票接收邮箱
	InvoiceAddress  string `json:"invoice_address"`   // 发票邮寄地址
	InvoicePhone    string `json:"invoice_phone"`     // 发票联系电话
	InvoiceRemark   string `json:"invoice_remark"`    // 发票备注
	BankName        string `json:"bank_name"`         // 开户行
	BankAccount     string `json:"bank_account"`      // 银行账号
	BankAccountName string `json:"bank_account_name"` // 开户名
}

// CompanyInfo 企业信息
type CompanyInfo struct {
	CompanyName         string `json:"company_name"`         // 企业名称
	BusinessLicense     string `json:"business_license"`     // 营业执照号
	LegalRepresentative string `json:"legal_representative"` // 法定代表人
	RegisteredAddress   string `json:"registered_address"`   // 注册地址
	RegisteredCapital   string `json:"registered_capital"`   // 注册资本
	BusinessScope       string `json:"business_scope"`       // 经营范围
	ContactPerson       string `json:"contact_person"`       // 联系人
	ContactPhone        string `json:"contact_phone"`        // 联系电话
	ContactEmail        string `json:"contact_email"`        // 联系邮箱
}

// PlatformOptions 支付平台配置
type PlatformOptions struct {
	AppID          string            `json:"app_id"`          // 应用ID
	MerchantID     string            `json:"merchant_id"`     // 商户ID
	PrivateKey     string            `json:"private_key"`     // 私钥
	PublicKey      string            `json:"public_key"`      // 公钥
	NotifyKey      string            `json:"notify_key"`      // 通知密钥
	ExtraParams    map[string]string `json:"extra_params"`    // 额外参数
	Environment    string            `json:"environment"`     // 环境(sandbox/production)
	PaymentMethods []string          `json:"payment_methods"` // 支持的支付方式
}

// RefundInfo 退款信息
type RefundInfo struct {
	RefundID        string  `json:"refund_id"`        // 退款ID
	RefundAmount    float64 `json:"refund_amount"`    // 退款金额
	RefundReason    string  `json:"refund_reason"`    // 退款原因
	RefundStatus    int8    `json:"refund_status"`    // 退款状态
	RefundTime      string  `json:"refund_time"`      // 退款时间
	ProcessorID     string  `json:"processor_id"`     // 处理人ID
	ProcessorName   string  `json:"processor_name"`   // 处理人姓名
	ProcessorRemark string  `json:"processor_remark"` // 处理备注
}

// Payment DTO结构
type Payment struct {
	PaymentID            string           `json:"payment_id"`             // 支付记录唯一标识
	UserID               string           `json:"user_id"`                // 关联的用户ID
	PaymentPlatform      string           `json:"payment_platform"`       // 支付平台
	PaymentScene         string           `json:"payment_scene"`          // 支付场景
	PaymentMethod        string           `json:"payment_method"`         // 支付方式
	PaymentCurrency      string           `json:"payment_currency"`       // 支付币种
	PaymentAmount        float64          `json:"payment_amount"`         // 支付金额
	PaymentStatus        int8             `json:"payment_status"`         // 支付状态
	PaymentType          string           `json:"payment_type"`           // 支付类型
	PaymentOrderNo       string           `json:"payment_order_no"`       // 支付订单号
	PaymentTransactionID string           `json:"payment_transaction_id"` // 第三方交易ID
	PaymentTitle         string           `json:"payment_title"`          // 支付标题
	PaymentDescription   string           `json:"payment_description"`    // 支付描述
	NotifyURL            string           `json:"notify_url"`             // 支付回调通知地址
	PaymenterName        string           `json:"paymenter_name"`         // 付款人姓名
	PaymenterEmail       string           `json:"paymenter_email"`        // 付款人邮箱
	PaymenterPhone       string           `json:"paymenter_phone"`        // 付款人电话
	PaymentInfo          PaymentInfo      `json:"payment_info"`           // 账单信息
	CompanyInfo          CompanyInfo      `json:"company_info"`           // 企业付款信息
	CallbackData         interface{}      `json:"callback_data"`          // 支付回调数据
	PlatformOptions      PlatformOptions  `json:"platform_options"`       // 支付平台配置
	PaymentTime          utils.MySQLTime  `json:"payment_time"`           // 支付成功时间
	ExpireTime           utils.MySQLTime  `json:"expire_time"`            // 支付过期时间
	RefundInfo           RefundInfo       `json:"refund_info"`            // 退款信息
	CreatedAt            utils.MySQLTime  `json:"created_at"`             // 创建时间
	UpdatedAt            utils.MySQLTime  `json:"updated_at"`             // 更新时间
	DeletedAt            *utils.MySQLTime `json:"deleted_at,omitempty"`   // 删除时间
}
