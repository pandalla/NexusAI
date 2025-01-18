package log

import (
	"nexus-ai/common"
	"nexus-ai/utils"
)

type WorkerLogDetails struct {
	Content string `json:"content"`
}

// WorkerLog DTO结构
type WorkerLog struct {
	WorkerLogID     string `json:"worker_log_id"`     // 日志唯一标识
	WorkerClusterID string `json:"worker_cluster_id"` // 工作节点ID
	WorkerGroupID   string `json:"worker_group_id"`   // 工作节点ID
	WorkerNodeID    string `json:"worker_node_id"`    // 工作节点ID
	RequestID       string `json:"request_id"`        // 关联的请求ID

	EventType     string           `json:"event_type"`     // 事件类型(task/redis/mysql等)
	LogLevel      string           `json:"log_level"`      // 日志级别(info/warn/error)
	LogDetails    WorkerLogDetails `json:"log_details"`    // 详细日志信息
	ResourceType  string           `json:"resource_type"`  // 资源类型(cpu/memory/disk等)
	ResourceUsage common.JSON      `json:"resource_usage"` // 资源使用情况
	ErrorType     string           `json:"error_type"`     // 错误类型
	ErrorMessage  string           `json:"error_message"`  // 错误信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
