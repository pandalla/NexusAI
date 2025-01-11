package model

import (
	"nexus-ai/common"
	"time"

	"gorm.io/gorm"
)

type UserGroup struct {
	UserGroupID          string `gorm:"column:user_group_id;type:char(36);primaryKey;default:(UUID())" json:"user_group_id"` // 用户组唯一标识
	UserGroupName        string `gorm:"column:user_group_name;size:100;not null" json:"user_group_name"`                     // 用户组名称
	UserGroupDescription string `gorm:"column:user_group_description;type:text" json:"user_group_description"`               // 用户组描述

	UserGroupPriceFactor common.JSON `gorm:"column:user_group_price_factor;type:json" json:"user_group_price_factor"` // 用户组价格系数
	UserGroupOptions     common.JSON `gorm:"column:user_group_options;type:json" json:"user_group_options"`           // 用户组配置

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                                                      // 删除时间
}

func (UserGroup) TableName() string {
	return "user_groups"
}
