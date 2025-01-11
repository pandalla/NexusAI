package log

import (
	"time"

	"nexus-ai/common"
)

// RedisPersistLog 存储Redis持久化操作日志，包括数据持久化、同步等信息
type RedisPersistLog struct {
	// 日志唯一标识
	RedisPersistLogID string `gorm:"column:redis_persist_log_id;type:char(36);primaryKey;default:(UUID())" json:"redis_persist_log_id"`

	// 节点类型(master/worker)
	NodeType string `gorm:"column:node_type;size:20;index;not null" json:"node_type"`

	// 节点ID
	ServiceID string `gorm:"column:service_id;size:50;index;not null" json:"service_id"`

	// 持久化类型(rdb/aof/混合)
	PersistType string `gorm:"column:persist_type;size:50;index;not null" json:"persist_type"`

	// 事件类型(start/complete/error)
	EventType string `gorm:"column:event_type;size:50;index;not null" json:"event_type"`

	// 目标MySQL表
	TargetTable string `gorm:"column:target_table;size:100" json:"target_table"`

	// 数据大小(bytes)
	DataSize int64 `gorm:"column:data_size;not null;default:0" json:"data_size"`

	// 影响行数
	AffectedRows int `gorm:"column:affected_rows;not null;default:0" json:"affected_rows"`

	// 开始时间
	StartTime time.Time `gorm:"column:start_time;not null" json:"start_time"`

	// 结束时间
	EndTime *time.Time `gorm:"column:end_time" json:"end_time"`

	// 持续时间(ms)
	Duration int `gorm:"column:duration;not null;default:0" json:"duration"`

	// 错误类型
	ErrorType string `gorm:"column:error_type;size:50" json:"error_type"`

	// 错误信息
	ErrorMessage string `gorm:"column:error_message;type:text" json:"error_message"`

	// 详细日志信息
	LogDetails common.JSON `gorm:"column:log_details;type:json" json:"log_details"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`
}

// TableName 表名
func (RedisPersistLog) TableName() string {
	return "redis_persist_logs"
}
