package worker

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// WorkerNode 存储工作节点实例信息，包括节点状态、资源使用和性能指标等
type WorkerNode struct {
	// 节点实例唯一标识
	NodeID string `gorm:"column:node_id;type:char(36);primaryKey;default:(UUID())" json:"node_id"`

	// 关联的工作节点ID
	WorkerID string `gorm:"column:worker_id;type:char(36);index;not null;foreignKey:Worker(WorkerID)" json:"worker_id"`

	// 关联的集群ID
	ClusterID string `gorm:"column:cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(ClusterID)" json:"cluster_id"`

	// 节点IP地址
	NodeIP string `gorm:"column:node_ip;size:50;not null" json:"node_ip"`

	// 节点端口
	NodePort int `gorm:"column:node_port;not null" json:"node_port"`

	// 节点区域
	NodeRegion string `gorm:"column:node_region;size:50;index" json:"node_region"`

	// 节点可用区
	NodeZone string `gorm:"column:node_zone;size:50;index" json:"node_zone"`

	// 节点实例配置
	NodeOptions common.JSON `gorm:"column:node_options;type:json" json:"node_options"`

	// 实例状态(1:运行中 2:启动中 3:停止中 4:已停止 5:异常)
	NodeStatus int8 `gorm:"column:node_status;index;not null;default:1" json:"node_status"`

	// CPU使用率
	CPUUsage float64 `gorm:"column:cpu_usage;type:decimal(5,2);default:0.00" json:"cpu_usage"`

	// 内存使用率
	MemoryUsage float64 `gorm:"column:memory_usage;type:decimal(5,2);default:0.00" json:"memory_usage"`

	// 网络统计信息
	NetworkStats common.JSON `gorm:"column:network_stats;type:json" json:"network_stats"`

	// 性能统计信息
	PerformanceStats common.JSON `gorm:"column:performance_stats;type:json" json:"performance_stats"`

	// 最后健康检查时间
	HealthCheck time.Time `gorm:"column:health_check;not null;default:CURRENT_TIMESTAMP(3)" json:"health_check"`

	// 最后错误信息
	LastError string `gorm:"column:last_error;type:text" json:"last_error"`

	// 启动时间
	StartupTime time.Time `gorm:"column:startup_time;not null" json:"startup_time"`

	// 实例状态 1:正常 0:禁用
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (WorkerNode) TableName() string {
	return "worker_nodes"
}
