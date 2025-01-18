package log

import (
	"nexus-ai/utils"
)

type GatewayLogDetails struct {
	Content string `json:"content"`
}

// GatewayLog DTO结构
type GatewayLog struct {
	GatewayLogID string            `json:"gateway_log_id"` // 日志唯一标识
	GatewayID    string            `json:"gateway_id"`     // 网关节点ID
	EventType    string            `json:"event_type"`     // 事件类型(request/response/error等)
	LogLevel     string            `json:"log_level"`      // 日志级别(info/warn/error)
	LogDetails   GatewayLogDetails `json:"log_details"`    // 详细日志信息
	ErrorType    string            `json:"error_type"`     // 错误类型
	ErrorMessage string            `json:"error_message"`  // 错误信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
