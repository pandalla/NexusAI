package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Token 存储系统中所有API访问令牌的信息，包括权限、配置和使用限制等
type Token struct {
	// 令牌唯一标识
	TokenID string `gorm:"column:token_id;type:char(36);primaryKey;default:(UUID())" json:"token_id"`

	// 关联的用户ID，与users表关联
	UserID string `gorm:"column:user_id;index;not null;type:char(36);foreignKey:User(UserID)" json:"user_id"`

	// 令牌名称，用于标识和管理，最大长度100字符
	TokenName string `gorm:"column:token_name;size:100;index;not null" json:"token_name"`

	// 令牌密钥，用于API认证，唯一索引，最大长度255字符
	TokenKey string `gorm:"column:token_key;size:255;uniqueIndex;not null" json:"token_key"`

	// 令牌状态，1:正常 0:禁用
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 令牌组配置，包含令牌的分组信息(如普通/高级/尊享/专线/特权等)
	TokenGroup common.JSON `gorm:"column:token_group;type:json" json:"token_group"`

	// 支持的模型列表配置，关联models表的model_id
	TokenModels common.JSON `gorm:"column:token_models;type:json" json:"token_models"`

	// 令牌的渠道配置，指定使用channels表的channel_id对应的渠道
	TokenChannels common.JSON `gorm:"column:token_channels;type:json" json:"token_channels"`

	// 令牌的特殊配置选项(如频率限制/并发限制/IP白名单等)
	TokenOptions common.JSON `gorm:"column:token_options;type:json" json:"token_options"`

	// 令牌过期时间，为空表示永不过期
	ExpireTime *time.Time `gorm:"column:expire_time" json:"expire_time"`

	// 令牌创建时间
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 令牌信息最后更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Token) TableName() string {
	return "tokens"
}
