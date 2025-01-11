package log

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 网关服务的日志信息，包括请求处理、路由转发等操作的记录
type GatewayLog struct {
	GatewayLogID string `gorm:"column:gateway_log_id;type:char(36);primaryKey;default:(UUID())" json:"gateway_log_id"` // 日志唯一标识

	LogLevel     string      `gorm:"column:log_level;size:20;index;not null" json:"log_level"`   // 日志级别(info/warn/error)
	EventType    string      `gorm:"column:event_type;size:50;index;not null" json:"event_type"` // 事件类型(route/forward/error等)
	ErrorType    string      `gorm:"column:error_type;size:50" json:"error_type"`                // 错误类型
	ErrorMessage string      `gorm:"column:error_message;type:text" json:"error_message"`        // 错误信息
	LogDetails   common.JSON `gorm:"column:log_details;type:json" json:"log_details"`            // 详细日志信息

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"` // 记录创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (GatewayLog) TableName() string {
	return "gateway_logs"
}
