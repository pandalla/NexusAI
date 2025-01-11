package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Quota 存储用户的配额信息，包括额度、使用情况和有效期等
type Quota struct {
	// 配额记录唯一标识 uuid
	QuotaID string `gorm:"column:quota_id;type:char(36);primaryKey;default:(UUID())" json:"quota_id"`

	// 关联的用户ID
	UserID string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`

	// 关联的支付ID
	PayID string `gorm:"column:pay_id;type:char(36);index;not null;foreignKey:Payment(PayID)" json:"pay_id"`

	// 配额总金额
	QuotaAmount float64 `gorm:"column:quota_amount;type:decimal(10,2);not null;default:0.00" json:"quota_amount"`

	// 剩余金额
	RemainingAmount float64 `gorm:"column:remaining_amount;type:decimal(10,2);not null;default:0.00" json:"remaining_amount"`

	// 冻结金额(处理中的请求)
	FrozenAmount float64 `gorm:"column:frozen_amount;type:decimal(10,2);not null;default:0.00" json:"frozen_amount"`

	// 配额类型(充值/赠送/奖励等)
	QuotaType string `gorm:"column:quota_type;size:20;index;not null" json:"quota_type"`

	// 配额等级(普通/高级/尊享等)
	QuotaLevel string `gorm:"column:quota_level;size:20;index;not null" json:"quota_level"`

	// 有效期(天)
	ValidPeriod int `gorm:"column:valid_period;not null;default:0" json:"valid_period"`

	// 生效时间
	StartTime time.Time `gorm:"column:start_time;not null" json:"start_time"`

	// 过期时间
	ExpireTime time.Time `gorm:"column:expire_time;not null" json:"expire_time"`

	// 配额特殊配置
	QuotaConfig common.JSON `gorm:"column:quota_config;type:json" json:"quota_config"`

	// 配额状态(1:正常 2:冻结 3:过期)
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Quota) TableName() string {
	return "quotas"
}
