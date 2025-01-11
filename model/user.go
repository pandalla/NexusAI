package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	UserID        uint64         `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	Username      string         `gorm:"column:username;size:50" json:"username"`
	Password      string         `gorm:"column:password;size:255" json:"password"`
	Email         string         `gorm:"column:email;size:100;uniqueIndex" json:"email"`
	Phone         string         `gorm:"column:phone;size:20;uniqueIndex" json:"phone"`
	OAuthInfo     common.JSON    `gorm:"column:oauth_info;type:json" json:"oauth_info"`
	UserTokens    common.JSON    `gorm:"column:user_tokens;type:json" json:"user_tokens"`
	UserGroup     string         `gorm:"column:user_group;size:50;index;not null" json:"user_group"`
	UserQuota     common.JSON    `gorm:"column:user_quota;type:json" json:"user_quota"`
	UserOptions   common.JSON    `gorm:"column:user_options;type:json" json:"user_options"`
	LastLoginTime *time.Time     `gorm:"column:last_login_time" json:"last_login_time"`
	LastLoginIP   string         `gorm:"column:last_login_ip;size:50" json:"last_login_ip"`
	Status        int8           `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}
