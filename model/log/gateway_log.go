package log

import (
	"time"

	"nexus-ai/common"
)

// GatewayLog 存储网关服务的日志信息，包括请求处理、路由转发等操作的记录
type GatewayLog struct {
	// 日志唯一标识
	GatewayLogID string `gorm:"column:gateway_log_id;type:char(36);primaryKey;default:(UUID())" json:"gateway_log_id"`

	// 网关节点ID
	GatewayID string `gorm:"column:gateway_id;type:char(36);index;not null" json:"gateway_id"`

	// 日志级别(info/warn/error)
	LogLevel string `gorm:"column:log_level;size:20;index;not null" json:"log_level"`

	// 事件类型(route/forward/error等)
	EventType string `gorm:"column:event_type;size:50;index;not null" json:"event_type"`

	// 响应状态码
	ResponseCode int `gorm:"column:response_code;not null;default:0" json:"response_code"`

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
func (GatewayLog) TableName() string {
	return "gateway_logs"
}
