package log

import (
	"nexus-ai/utils"
)

type WorkerMySQLLogDetails struct {
	Content string `json:"content"`
}

// WorkerMySQLLog DTO结构
type WorkerMySQLLog struct {
	WorkerMySQLLogID string `json:"worker_mysql_log_id"` // 日志唯一标识
	WorkerClusterID  string `json:"worker_cluster_id"`   // 工作节点ID
	WorkerGroupID    string `json:"worker_group_id"`     // 工作节点ID
	WorkerNodeID     string `json:"worker_node_id"`      // 工作节点ID
	RequestID        string `json:"request_id"`          // 关联的请求ID

	EventType     string                `json:"event_type"`     // 事件类型(query/transaction等)
	LogLevel      string                `json:"log_level"`      // 日志级别(info/warn/error)
	LogDetails    WorkerMySQLLogDetails `json:"log_details"`    // 详细日志信息
	Operation     string                `json:"operation"`      // 操作类型(query/transaction等)
	ExecutionTime int                   `json:"execution_time"` // 执行时间(毫秒)
	ErrorType     string                `json:"error_type"`     // 错误类型
	ErrorMessage  string                `json:"error_message"`  // 错误信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
