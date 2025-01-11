package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Token 令牌表
type Token struct {
	TokenID      uint64         `gorm:"column:token_id;primaryKey;autoIncrement" json:"token_id"`
	UserID       uint64         `gorm:"column:user_id;index;not null" json:"user_id"`
	TokenName    string         `gorm:"column:token_name;size:100;not null" json:"token_name"`
	TokenGroup   common.JSON    `gorm:"column:token_group;type:json" json:"token_group"`
	TokenModels  common.JSON    `gorm:"column:token_models;type:json" json:"token_models"`
	TokenOptions common.JSON    `gorm:"column:token_options;type:json" json:"token_options"`
	TokenKey     string         `gorm:"column:token_key;size:255;uniqueIndex;not null" json:"token_key"`
	ExpireTime   *time.Time     `gorm:"column:expire_time" json:"expire_time"`
	Status       int8           `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Token) TableName() string {
	return "tokens"
}
