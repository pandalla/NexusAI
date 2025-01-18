package log

import (
	"nexus-ai/utils"
)

type MasterLogDetails struct {
	Content string `json:"content"`
}

// MasterLog DTO结构
type MasterLog struct {
	MasterLogID string `json:"master_log_id"` // 日志唯一标识
	MasterID    string `json:"master_id"`     // 主服务节点ID

	EventType  string           `json:"event_type"`  // 事件类型(service/redis/mysql/worker等)
	LogLevel   string           `json:"log_level"`   // 日志级别(info/warn/error)
	LogDetails MasterLogDetails `json:"log_details"` // 详细日志信息
	RequestID  string           `json:"request_id"`  // 关联的请求ID

	ErrorType    string `json:"error_type"`    // 错误类型
	ErrorMessage string `json:"error_message"` // 错误信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
