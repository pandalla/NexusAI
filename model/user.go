package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 存储系统中所有用户的基本信息、认证信息、配额信息等
type User struct {
	// 用户唯一标识
	UserID string `gorm:"column:user_id;type:char(36);primaryKey;default:(UUID())" json:"user_id"`

	// 用户名，1-50字符
	Username string `gorm:"column:username;size:50" json:"username"`

	// 密码哈希，最大长度255字符
	Password string `gorm:"column:password;size:255" json:"password"`

	// 邮箱地址，唯一索引，最大长度100字符
	Email string `gorm:"column:email;size:100;uniqueIndex" json:"email"`

	// 手机号码，唯一索引，最大长度20字符
	Phone string `gorm:"column:phone;size:20;uniqueIndex" json:"phone"`

	// 用户组，最大长度50字符
	UserGroup string `gorm:"column:user_group;size:50;index;not null;default:default" json:"user_group"`

	// 第三方登录相关的认证信息(如GitHub、微信、Google等)
	OAuthInfo common.JSON `gorm:"column:oauth_info;type:json" json:"oauth_info"`

	// 各种资源使用配额信息
	UserQuota common.JSON `gorm:"column:user_quota;type:json" json:"user_quota"`

	// 个性化设置和偏好选项
	UserOptions common.JSON `gorm:"column:user_options;type:json" json:"user_options"`

	// 最后一次成功登录的时间
	LastLoginTime *time.Time `gorm:"column:last_login_time" json:"last_login_time"`

	// 最后一次登录的IP地址
	LastLoginIP string `gorm:"column:last_login_ip;size:100" json:"last_login_ip"`

	// 用户状态，1:正常 0:禁用
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 用户账号的创建时间
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 用户信息的最后更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}
