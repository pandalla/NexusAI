package log

import (
	"time"

	"nexus-ai/common"
)

// WorkerMySQLLog 工作节点MySQL日志表
type WorkerMySQLLog struct {
	LogID         uint64      `gorm:"column:log_id;primaryKey;autoIncrement" json:"log_id"`
	WorkerID      uint64      `gorm:"column:worker_id;index;not null" json:"worker_id"`
	NodeID        string      `gorm:"column:node_id;size:50;index;not null" json:"node_id"`
	LogLevel      string      `gorm:"column:log_level;size:20;index;not null" json:"log_level"`
	EventType     string      `gorm:"column:event_type;size:50;index;not null" json:"event_type"`
	Operation     string      `gorm:"column:operation;size:50" json:"operation"`
	RequestID     string      `gorm:"column:request_id;size:64;index" json:"request_id"`
	ExecutionTime int         `gorm:"column:execution_time" json:"execution_time"`
	ErrorType     string      `gorm:"column:error_type;size:50" json:"error_type"`
	ErrorMessage  string      `gorm:"column:error_message;type:text" json:"error_message"`
	LogDetails    common.JSON `gorm:"column:log_details;type:json" json:"log_details"`
	CreatedAt     time.Time   `gorm:"column:created_at;index;not null" json:"created_at"`
}

// TableName 表名
func (WorkerMySQLLog) TableName() string {
	return "worker_mysql_logs"
}
