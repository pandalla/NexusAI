package log

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 主服务的日志信息，包括节点管理、资源调度等操作的记录
type MasterLog struct {
	MasterLogID string `gorm:"column:master_log_id;type:char(36);primaryKey;default:(UUID())" json:"master_log_id"` // 日志唯一标识
	MasterID    string `gorm:"column:master_id;type:char(36);index;not null" json:"master_id"`                      // 主服务节点ID

	LogLevel  string `gorm:"column:log_level;size:20;index;not null" json:"log_level"`   // 日志级别(info/warn/error)
	EventType string `gorm:"column:event_type;size:50;index;not null" json:"event_type"` // 事件类型(service/redis/mysql/worker等)
	RequestID string `gorm:"column:request_id;size:64;index" json:"request_id"`          // 关联的请求ID

	ErrorType    string      `gorm:"column:error_type;size:50" json:"error_type"`         // 错误类型
	ErrorMessage string      `gorm:"column:error_message;type:text" json:"error_message"` // 错误信息
	LogDetails   common.JSON `gorm:"column:log_details;type:json" json:"log_details"`     // 详细日志信息

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 记录创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 记录更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (MasterLog) TableName() string {
	return "master_logs"
}
