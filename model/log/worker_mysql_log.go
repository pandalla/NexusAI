package log

import (
	"nexus-ai/common"
	"nexus-ai/utils"

	"gorm.io/gorm"
)

// WorkerMySQLLog 存储工作节点的MySQL操作日志，包括查询执行、性能监控等信息
type WorkerMySQLLog struct {
	WorkerMySQLLogID string `gorm:"column:worker_mysql_log_id;type:char(36);primaryKey;default:(UUID())" json:"worker_mysql_log_id"`
	WorkerClusterID  string `gorm:"column:worker_cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(WorkerClusterID)" json:"worker_cluster_id"` // 关联的集群ID
	WorkerGroupID    string `gorm:"column:worker_group_id;type:char(36);index;not null;foreignKey:WorkerGroup(WorkerGroupID)" json:"worker_group_id"`         // 关联的工作节点组ID
	WorkerNodeID     string `gorm:"column:worker_node_id;type:char(36);index;not null;foreignKey:WorkerNode(WorkerNodeID)" json:"worker_node_id"`             // 关联的节点实例ID
	RequestID        string `gorm:"column:request_id;size:64;index" json:"request_id"`                                                                        // 关联的请求ID

	EventType     string      `gorm:"column:event_type;size:50;index;not null" json:"event_type"`     // 事件类型(connect/query/error等)
	LogLevel      string      `gorm:"column:log_level;size:20;index;not null" json:"log_level"`       // 日志级别(info/warn/error)
	LogDetails    common.JSON `gorm:"column:log_details;type:json" json:"log_details"`                // 详细日志信息
	Operation     string      `gorm:"column:operation;size:50" json:"operation"`                      // 操作类型(connect/query/error等)
	ExecutionTime int         `gorm:"column:execution_time;not null;default:0" json:"execution_time"` // 执行时间(ms)
	ErrorType     string      `gorm:"column:error_type;size:50" json:"error_type"`                    // 错误类型
	ErrorMessage  string      `gorm:"column:error_message;type:text" json:"error_message"`            // 错误信息

	CreatedAt utils.MySQLTime `gorm:"column:created_at;index;not null" json:"created_at"` // 记录创建时间
	UpdatedAt utils.MySQLTime `gorm:"column:updated_at;not null" json:"updated_at"`       // 记录更新时间
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (WorkerMySQLLog) TableName() string {
	return "worker_mysql_logs"
}

// BeforeCreate 在创建记录前自动设置时间
func (workerMySQLLog *WorkerMySQLLog) BeforeCreate(tx *gorm.DB) error {
	workerMySQLLog.CreatedAt = utils.MySQLTime(utils.GetTime())
	workerMySQLLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}

// BeforeUpdate 在更新记录前自动设置更新时间
func (workerMySQLLog *WorkerMySQLLog) BeforeUpdate(tx *gorm.DB) error {
	workerMySQLLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}
