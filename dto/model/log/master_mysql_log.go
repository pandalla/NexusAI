package log

import (
	"nexus-ai/utils"
)

type MasterMySQLLogDetails struct {
	Content string `json:"content"`
}

// MasterMySQLLog DTO结构
type MasterMySQLLog struct {
	MasterMySQLLogID string `json:"master_mysql_log_id"` // 日志唯一标识
	MasterID         string `json:"master_id"`           // 主服务节点ID
	RequestID        string `json:"request_id"`          // 关联的请求ID

	EventType  string                `json:"event_type"`  // 事件类型(connect/query/error等)
	LogLevel   string                `json:"log_level"`   // 日志级别(info/warn/error)
	LogDetails MasterMySQLLogDetails `json:"log_details"` // 详细日志信息

	Operation     string `json:"operation"`      // 操作类型
	ExecutionTime int    `json:"execution_time"` // 执行时间(ms)
	ErrorType     string `json:"error_type"`     // 错误类型
	ErrorMessage  string `json:"error_message"`  // 错误信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
