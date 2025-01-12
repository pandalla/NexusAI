package log

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 主服务的MySQL操作日志，包括查询执行、性能监控等信息
type MasterMySQLLog struct {
	MasterMySQLLogID string `gorm:"column:master_mysql_log_id;type:char(36);primaryKey;default:(UUID())" json:"master_mysql_log_id"`
	MasterID         string `gorm:"column:master_id;type:char(36);index;not null" json:"master_id"` // 主服务节点ID
	RequestID        string `gorm:"column:request_id;size:64;index" json:"request_id"`              // 关联的请求ID

	LogLevel  string `gorm:"column:log_level;size:20;index;not null" json:"log_level"`   // 日志级别(info/warn/error)
	EventType string `gorm:"column:event_type;size:50;index;not null" json:"event_type"` // 事件类型(connect/query/error等)

	Operation     string      `gorm:"column:operation;size:50" json:"operation"`             // 操作类型
	ExecutionTime int         `gorm:"column:execution_time;default:0" json:"execution_time"` // 执行时间(ms)
	ErrorType     string      `gorm:"column:error_type;size:50" json:"error_type"`           // 错误类型
	ErrorMessage  string      `gorm:"column:error_message;type:text" json:"error_message"`   // 错误信息
	LogDetails    common.JSON `gorm:"column:log_details;type:json" json:"log_details"`       // 详细日志信息

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 记录创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 记录更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (MasterMySQLLog) TableName() string {
	return "master_mysql_logs"
}
