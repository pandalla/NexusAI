package log

import (
	"nexus-ai/common"
	"nexus-ai/utils"

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

	CreatedAt utils.MySQLTime `gorm:"column:created_at;index;not null" json:"created_at"` // 记录创建时间
	UpdatedAt utils.MySQLTime `gorm:"column:updated_at;not null" json:"updated_at"`       // 记录更新时间
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (MasterLog) TableName() string {
	return "master_logs"
}

// BeforeCreate 在创建记录前自动设置时间
func (masterLog *MasterLog) BeforeCreate(tx *gorm.DB) error {
	masterLog.CreatedAt = utils.MySQLTime(utils.GetTime())
	masterLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}

// BeforeUpdate 在更新记录前自动设置更新时间
func (masterLog *MasterLog) BeforeUpdate(tx *gorm.DB) error {
	masterLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}
