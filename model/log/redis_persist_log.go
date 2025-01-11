package log

import (
	"time"

	"nexus-ai/common"
)

// RedisPersistLog Redis持久化日志表
type RedisPersistLog struct {
	LogID        uint64      `gorm:"column:log_id;primaryKey;autoIncrement" json:"log_id"`
	NodeType     string      `gorm:"column:node_type;size:20;index;not null" json:"node_type"`
	NodeID       string      `gorm:"column:node_id;size:50;index;not null" json:"node_id"`
	PersistType  string      `gorm:"column:persist_type;size:50;index;not null" json:"persist_type"`
	EventType    string      `gorm:"column:event_type;size:50;index;not null" json:"event_type"`
	TargetTable  string      `gorm:"column:target_table;size:100" json:"target_table"`
	DataSize     int64       `gorm:"column:data_size" json:"data_size"`
	AffectedRows int         `gorm:"column:affected_rows" json:"affected_rows"`
	StartTime    time.Time   `gorm:"column:start_time;not null" json:"start_time"`
	EndTime      *time.Time  `gorm:"column:end_time" json:"end_time"`
	Duration     int         `gorm:"column:duration" json:"duration"`
	ErrorType    string      `gorm:"column:error_type;size:50" json:"error_type"`
	ErrorMessage string      `gorm:"column:error_message;type:text" json:"error_message"`
	LogDetails   common.JSON `gorm:"column:log_details;type:json" json:"log_details"`
	CreatedAt    time.Time   `gorm:"column:created_at;index;not null" json:"created_at"`
}

// TableName 表名
func (RedisPersistLog) TableName() string {
	return "redis_persist_logs"
}
