package model

import (
	"nexus-ai/common"
	"nexus-ai/utils"

	"gorm.io/gorm"
)

// 系统中所有API访问令牌的信息，包括权限、配置和使用限制等
type Token struct {
	TokenID   string `gorm:"column:token_id;type:char(36);primaryKey;default:(UUID())" json:"token_id"`          // 令牌唯一标识
	UserID    string `gorm:"column:user_id;index;not null;type:char(36);foreignKey:User(UserID)" json:"user_id"` // 关联的用户ID
	TokenName string `gorm:"column:token_name;size:100;index;not null" json:"token_name"`                        // 令牌名称，用于标识和管理，最大长度100字符
	TokenKey  string `gorm:"column:token_key;size:255;uniqueIndex;not null" json:"token_key"`                    // 令牌密钥，用于API认证，唯一索引，最大长度255字符
	Status    int8   `gorm:"column:status;index;not null;default:1" json:"status"`                               // 令牌状态，1:正常 0:禁用

	TokenQuotaTotal  float64 `gorm:"column:token_quota_total;type:decimal(10,2);not null;default:0.00" json:"token_quota_total"`   // 令牌配额
	TokenQuotaUsed   float64 `gorm:"column:token_quota_used;type:decimal(10,2);not null;default:0.00" json:"token_quota_used"`     // 已使用配额
	TokenQuotaLeft   float64 `gorm:"column:token_quota_left;type:decimal(10,2);not null;default:0.00" json:"token_quota_left"`     // 剩余配额
	TokenQuotaFrozen float64 `gorm:"column:token_quota_frozen;type:decimal(10,2);not null;default:0.00" json:"token_quota_frozen"` // 冻结配额

	TokenChannels common.JSON     `gorm:"column:token_channels;type:json" json:"token_channels"` // 令牌的渠道配置，指定使用channels表的channel_id对应的渠道
	TokenModels   common.JSON     `gorm:"column:token_models;type:json" json:"token_models"`     // 支持的模型列表配置，关联models表的model_id
	TokenOptions  common.JSON     `gorm:"column:token_options;type:json" json:"token_options"`   // 令牌的特殊配置选项(如频率限制/并发限制/IP白名单等)
	ExpireTime    utils.MySQLTime `gorm:"column:expire_time" json:"expire_time"`                 // 令牌过期时间，为空表示永不过期

	CreatedAt utils.MySQLTime `gorm:"column:created_at;not null" json:"created_at"` // 令牌创建时间
	UpdatedAt utils.MySQLTime `gorm:"column:updated_at;not null" json:"updated_at"` // 令牌信息最后更新时间
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`          // 软删除时间
}

// TableName 表名
func (Token) TableName() string {
	return "tokens"
}

// BeforeCreate 在创建记录前自动设置时间
func (token *Token) BeforeCreate(tx *gorm.DB) error {
	token.CreatedAt = utils.MySQLTime(utils.GetTime())
	token.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}

// BeforeUpdate 在更新记录前自动设置更新时间
func (token *Token) BeforeUpdate(tx *gorm.DB) error {
	token.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}
