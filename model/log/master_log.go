package log

import (
	"time"

	"nexus-ai/common"
)

// MasterLog 存储主服务的日志信息，包括节点管理、资源调度等操作的记录
type MasterLog struct {
	// 日志唯一标识
	MasterLogID string `gorm:"column:master_log_id;type:char(36);primaryKey;default:(UUID())" json:"master_log_id"`

	// 主服务节点ID
	MasterID string `gorm:"column:master_id;type:char(36);index;not null" json:"master_id"`

	// 日志级别(info/warn/error)
	LogLevel string `gorm:"column:log_level;size:20;index;not null" json:"log_level"`

	// 事件类型(node/redis/mysql/worker等)
	EventType string `gorm:"column:event_type;size:50;index;not null" json:"event_type"`

	// 操作类型
	Operation string `gorm:"column:operation;size:50;not null" json:"operation"`

	// 关联的请求ID
	RequestID string `gorm:"column:request_id;size:64;index" json:"request_id"`

	// 目标类型(redis/mysql/worker)
	TargetType string `gorm:"column:target_type;size:50" json:"target_type"`

	// 目标节点
	TargetNode string `gorm:"column:target_node;size:100" json:"target_node"`

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
func (MasterLog) TableName() string {
	return "master_logs"
}
