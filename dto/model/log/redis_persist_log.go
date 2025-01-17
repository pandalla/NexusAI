package log

import (
	"nexus-ai/utils"
)

type RedisPersistLogDetails struct {
	Content string `json:"content"`
}

// RedisPersistLog DTO结构
type RedisPersistLog struct {
	RedisPersistLogID string                 `json:"redis_persist_log_id"` // 日志唯一标识
	NodeType          string                 `json:"node_type"`            // 节点类型(master/worker)
	ServiceID         string                 `json:"service_id"`           // 节点ID
	EventType         string                 `json:"event_type"`           // 事件类型(start/complete/error)
	LogLevel          string                 `json:"log_level"`            // 日志级别(info/warn/error)
	LogDetails        RedisPersistLogDetails `json:"log_details"`          // 详细日志信息

	PersistType  string          `json:"persist_type"`  // 持久化类型(rdb/aof/混合)
	TargetTable  string          `json:"target_table"`  // 目标MySQL表
	DataSize     int64           `json:"data_size"`     // 数据大小(字节)
	AffectedRows int             `json:"affected_rows"` // 影响的行数
	StartTime    utils.MySQLTime `json:"start_time"`    // 开始时间
	EndTime      utils.MySQLTime `json:"end_time"`      // 结束时间
	Duration     int             `json:"duration"`      // 持久化耗时(毫秒)
	ErrorType    string          `json:"error_type"`    // 错误类型
	ErrorMessage string          `json:"error_message"` // 错误信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
