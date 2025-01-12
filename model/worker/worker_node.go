package worker

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 工作节点实例信息，包括节点状态、资源使用和性能指标等
type WorkerNode struct {
	WorkerNodeID    string `gorm:"column:worker_node_id;type:char(36);primaryKey;default:(UUID())" json:"worker_node_id"`
	WorkerGroupID   string `gorm:"column:worker_group_id;type:char(36);index;not null;foreignKey:WorkerGroup(WorkerGroupID)" json:"worker_group_id"`
	WorkerClusterID string `gorm:"column:worker_cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(WorkerClusterID)" json:"worker_cluster_id"` // 关联的集群ID

	NodeIP           string      `gorm:"column:node_ip;size:50;not null" json:"node_ip"`                         // 节点IP地址
	NodePort         int         `gorm:"column:node_port;not null" json:"node_port"`                             // 节点端口
	NodeRegion       string      `gorm:"column:node_region;size:50;index" json:"node_region"`                    // 节点区域
	NodeZone         string      `gorm:"column:node_zone;size:50;index" json:"node_zone"`                        // 节点可用区
	CPUUsage         float64     `gorm:"column:cpu_usage;type:decimal(5,2);default:0.00" json:"cpu_usage"`       // CPU使用率
	MemoryUsage      float64     `gorm:"column:memory_usage;type:decimal(5,2);default:0.00" json:"memory_usage"` // 内存使用率
	NetworkStats     common.JSON `gorm:"column:network_stats;type:json" json:"network_stats"`                    // 网络统计信息
	PerformanceStats common.JSON `gorm:"column:performance_stats;type:json" json:"performance_stats"`            // 性能统计信息
	LastError        string      `gorm:"column:last_error;type:text" json:"last_error"`                          // 最后错误信息
	NodeStatus       int8        `gorm:"column:node_status;index;not null;default:1" json:"node_status"`         // 实例状态(1:运行中 2:启动中 3:停止中 4:已停止 5:异常)
	NodeOptions      common.JSON `gorm:"column:node_options;type:json" json:"node_options"`                      // 节点实例配置

	StartupTime time.Time `gorm:"column:startup_time;not null" json:"startup_time"`     // 启动时间
	Status      int8      `gorm:"column:status;index;not null;default:1" json:"status"` // 实例状态 1:正常 0:禁用

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 记录创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 记录更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                                                      // 软删除时间
}

// TableName 表名
func (WorkerNode) TableName() string {
	return "worker_nodes"
}
