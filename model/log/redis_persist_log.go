package log

import (
	"nexus-ai/common"
	"nexus-ai/utils"

	"gorm.io/gorm"
)

// Redis持久化操作日志，包括数据持久化、同步等信息
type RedisPersistLog struct {
	RedisPersistLogID string `gorm:"column:redis_persist_log_id;type:char(36);primaryKey;default:(UUID())" json:"redis_persist_log_id"`
	NodeType          string `gorm:"column:node_type;size:20;index;not null" json:"node_type"`   // 节点类型(master/worker)
	ServiceID         string `gorm:"column:service_id;size:50;index;not null" json:"service_id"` // 节点ID
	EventType         string `gorm:"column:event_type;size:50;index;not null" json:"event_type"` // 事件类型(start/complete/error)

	PersistType  string          `gorm:"column:persist_type;size:50;index;not null" json:"persist_type"` // 持久化类型(rdb/aof/混合)
	TargetTable  string          `gorm:"column:target_table;size:100" json:"target_table"`               // 目标MySQL表
	DataSize     int64           `gorm:"column:data_size;not null;default:0" json:"data_size"`           // 数据大小(bytes)
	AffectedRows int             `gorm:"column:affected_rows;not null;default:0" json:"affected_rows"`   // 影响行数
	StartTime    utils.MySQLTime `gorm:"column:start_time;not null" json:"start_time"`                   // 开始时间
	EndTime      utils.MySQLTime `gorm:"column:end_time" json:"end_time"`                                // 结束时间
	Duration     int             `gorm:"column:duration;not null;default:0" json:"duration"`             // 持续时间(ms)
	ErrorType    string          `gorm:"column:error_type;size:50" json:"error_type"`                    // 错误类型
	ErrorMessage string          `gorm:"column:error_message;type:text" json:"error_message"`            // 错误信息
	LogDetails   common.JSON     `gorm:"column:log_details;type:json" json:"log_details"`                // 详细日志信息

	CreatedAt utils.MySQLTime `gorm:"column:created_at;index;not null" json:"created_at"` // 记录创建时间
	UpdatedAt utils.MySQLTime `gorm:"column:updated_at;not null" json:"updated_at"`       // 记录更新时间
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (RedisPersistLog) TableName() string {
	return "redis_persist_logs"
}

// BeforeCreate 在创建记录前自动设置时间
func (redisPersistLog *RedisPersistLog) BeforeCreate(tx *gorm.DB) error {
	redisPersistLog.CreatedAt = utils.MySQLTime(utils.GetTime())
	redisPersistLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}

// BeforeUpdate 在更新记录前自动设置更新时间
func (redisPersistLog *RedisPersistLog) BeforeUpdate(tx *gorm.DB) error {
	redisPersistLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}
