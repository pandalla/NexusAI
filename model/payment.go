package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Payment 存储系统中所有支付交易记录，包括支付信息、账单信息和退款信息等
type Payment struct {
	// 支付记录唯一标识
	PayID string `gorm:"column:pay_id;type:char(36);primaryKey;default:(UUID())" json:"pay_id"`

	// 关联的用户ID
	UserID string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`

	// 支付订单号，系统生成的唯一编号
	OrderNo string `gorm:"column:order_no;size:64;uniqueIndex;not null" json:"order_no"`

	// 支付平台(stripe/paddle/alipay/wechat等)
	PayPlatform string `gorm:"column:pay_platform;size:50;index;not null" json:"pay_platform"`

	// 支付方式(信用卡/银行转账/支付宝/微信等)
	PayMethod string `gorm:"column:pay_method;size:50;not null" json:"pay_method"`

	// 支付类型(个人/企业)
	PayType string `gorm:"column:pay_type;size:20;index;not null" json:"pay_type"`

	// 支付场景(充值/订阅等)
	PayScene string `gorm:"column:pay_scene;size:20;index;not null" json:"pay_scene"`

	// 支付币种
	PayCurrency string `gorm:"column:pay_currency;size:10;not null" json:"pay_currency"`

	// 支付金额
	PayAmount float64 `gorm:"column:pay_amount;type:decimal(10,2);not null;default:0.00" json:"pay_amount"`

	// 支付状态(1:待支付 2:支付中 3:支付成功 4:支付失败 5:已退款)
	PayStatus int8 `gorm:"column:pay_status;index;not null;default:1" json:"pay_status"`

	// 支付标题
	PayTitle string `gorm:"column:pay_title;size:255;not null" json:"pay_title"`

	// 支付描述
	PayDesc string `gorm:"column:pay_desc;size:1000" json:"pay_desc"`

	// 付款人姓名/企业名称
	PayerName string `gorm:"column:payer_name;size:100" json:"payer_name"`

	// 付款人邮箱
	PayerEmail string `gorm:"column:payer_email;size:100" json:"payer_email"`

	// 付款人电话
	PayerPhone string `gorm:"column:payer_phone;size:20" json:"payer_phone"`

	// 企业付款信息(营业执照/税号等)
	CompanyInfo common.JSON `gorm:"column:company_info;type:json" json:"company_info"`

	// 账单信息(发票信息等)
	BillingInfo common.JSON `gorm:"column:billing_info;type:json" json:"billing_info"`

	// 支付回调通知地址
	NotifyURL string `gorm:"column:notify_url;size:255" json:"notify_url"`

	// 第三方交易ID
	TransactionID string `gorm:"column:transaction_id;size:64;index" json:"transaction_id"`

	// 支付回调数据
	CallbackData common.JSON `gorm:"column:callback_data;type:json" json:"callback_data"`

	// 支付平台特定配置
	PlatformConfig common.JSON `gorm:"column:platform_config;type:json" json:"platform_config"`

	// 退款信息
	RefundInfo common.JSON `gorm:"column:refund_info;type:json" json:"refund_info"`

	// 支付过期时间
	ExpireTime time.Time `gorm:"column:expire_time;not null" json:"expire_time"`

	// 支付成功时间
	PayTime *time.Time `gorm:"column:pay_time" json:"pay_time"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Payment) TableName() string {
	return "pays"
}
