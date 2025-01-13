package worker

import (
	"nexus-ai/common"
	"nexus-ai/utils"

	"gorm.io/gorm"
)

// 工作节点集群信息，包括集群配置、资源使用和性能统计等
type WorkerCluster struct {
	WorkerClusterID    string `gorm:"column:worker_cluster_id;type:char(36);primaryKey;default:(UUID())" json:"worker_cluster_id"` // 集群唯一标识
	WorkerClusterName  string `gorm:"column:worker_cluster_name;size:100;not null" json:"worker_cluster_name"`                     // 集群名称
	TotalWorkerGroups  int    `gorm:"column:total_worker_groups;not null;default:0" json:"total_worker_groups"`                    // 工作节点组总数
	ActiveWorkerGroups int    `gorm:"column:active_worker_groups;not null;default:0" json:"active_worker_groups"`                  // 活跃节点组数
	TotalWorkerNodes   int    `gorm:"column:total_worker_nodes;not null;default:0" json:"total_worker_nodes"`                      // 工作节点总数
	ActiveWorkerNodes  int    `gorm:"column:active_worker_nodes;not null;default:0" json:"active_worker_nodes"`                    // 活跃工作节点数

	WorkerClusterStatus  int8        `gorm:"column:worker_cluster_status;index;not null;default:1" json:"worker_cluster_status"` // 集群状态(1:正常 2:部分可用 3:异常)
	WorkerClusterOptions common.JSON `gorm:"column:worker_cluster_options;type:json" json:"worker_cluster_options"`              // 集群配置信息
	ResourceUsage        common.JSON `gorm:"column:resource_usage;type:json" json:"resource_usage"`                              // 资源使用统计
	PerformanceStats     common.JSON `gorm:"column:performance_stats;type:json" json:"performance_stats"`                        // 性能统计信息
	AlertOptions         common.JSON `gorm:"column:alert_options;type:json" json:"alert_options"`                                // 告警配置
	MaintenanceWindow    common.JSON `gorm:"column:maintenance_window;type:json" json:"maintenance_window"`                      // 维护窗口配置
	Status               int8        `gorm:"column:status;index;not null;default:1" json:"status"`                               // 集群状态 1:正常 0:禁用

	CreatedAt utils.MySQLTime `gorm:"column:created_at;index;not null" json:"created_at"`
	UpdatedAt utils.MySQLTime `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (WorkerCluster) TableName() string {
	return "worker_clusters"
}

// BeforeCreate 在创建记录前自动设置时间
func (workerCluster *WorkerCluster) BeforeCreate(tx *gorm.DB) error {
	workerCluster.CreatedAt = utils.MySQLTime(utils.GetTime())
	workerCluster.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}

// BeforeUpdate 在更新记录前自动设置更新时间
func (workerCluster *WorkerCluster) BeforeUpdate(tx *gorm.DB) error {
	workerCluster.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}
