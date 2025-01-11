package model

import (
	"nexus-ai/common"
	"time"

	"gorm.io/gorm"
)

// 渠道组信息，包括渠道组ID、名称、描述、价格系数和配置等
type ChannelGroup struct {
	ChannelGroupID          string `gorm:"column:channel_group_id;type:char(36);primaryKey;default:(UUID())" json:"channel_group_id"` // 渠道组ID
	ChannelGroupName        string `gorm:"column:channel_group_name;size:100;not null" json:"channel_group_name"`                     // 渠道组名称
	ChannelGroupDescription string `gorm:"column:channel_group_description;type:text" json:"channel_group_description"`               // 渠道组描述

	ChannelPriceFactor  common.JSON `gorm:"column:channel_price_factor;type:json" json:"channel_price_factor"`   // 渠道组价格系数
	ChannelGroupOptions common.JSON `gorm:"column:channel_group_options;type:json" json:"channel_group_options"` // 渠道组配置

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                                                      // 删除时间
}

func (ChannelGroup) TableName() string {
	return "channel_groups"
}
