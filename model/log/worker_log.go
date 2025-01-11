package log

import (
	"time"

	"nexus-ai/common"
)

// WorkerLog 存储工作节点的日志信息，包括任务处理、资源使用等操作的记录
type WorkerLog struct {
	// 日志唯一标识
	WorkerLogID string `gorm:"column:worker_log_id;type:char(36);primaryKey;default:(UUID())" json:"worker_log_id"`

	// 关联的集群ID
	ClusterID string `gorm:"column:cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(ClusterID)	" json:"cluster_id"`

	// 关联的工作节点ID
	WorkerID string `gorm:"column:worker_id;type:char(36);index;not null;foreignKey:Worker(WorkerID)" json:"worker_id"`

	// 节点实例ID
	NodeID string `gorm:"column:node_id;type:char(36);index;not null;foreignKey:WorkerNode(NodeID)" json:"node_id"`

	// 日志级别(info/warn/error)
	LogLevel string `gorm:"column:log_level;size:20;index;not null" json:"log_level"`

	// 事件类型(task/redis/mysql等)
	EventType string `gorm:"column:event_type;size:50;index;not null" json:"event_type"`

	// 关联的请求ID
	RequestID string `gorm:"column:request_id;size:64;index" json:"request_id"`

	// 资源类型(cpu/memory/disk等)
	ResourceType string `gorm:"column:resource_type;size:50" json:"resource_type"`

	// 资源使用情况
	ResourceUsage common.JSON `gorm:"column:resource_usage;type:json" json:"resource_usage"`

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
func (WorkerLog) TableName() string {
	return "worker_logs"
}
