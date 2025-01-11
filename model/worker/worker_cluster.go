package worker

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// WorkerCluster 存储工作节点集群信息，包括集群配置、资源使用和性能统计等
type WorkerCluster struct {
	// 集群唯一标识
	ClusterID string `gorm:"column:cluster_id;type:char(36);primaryKey;default:(UUID())" json:"cluster_id"`

	// 集群名称
	ClusterName string `gorm:"column:cluster_name;size:100;not null" json:"cluster_name"`

	// 工作节点总数
	TotalWorkers int `gorm:"column:total_workers;not null;default:0" json:"total_workers"`

	// 活跃节点数
	ActiveWorkers int `gorm:"column:active_workers;not null;default:0" json:"active_workers"`

	// 实例总数
	TotalInstances int `gorm:"column:total_instances;not null;default:0" json:"total_instances"`

	// 活跃实例数
	ActiveInstances int `gorm:"column:active_instances;not null;default:0" json:"active_instances"`

	// 集群状态(1:正常 2:部分可用 3:异常)
	ClusterStatus int8 `gorm:"column:cluster_status;index;not null;default:1" json:"cluster_status"`

	// 集群配置信息
	ClusterConfig common.JSON `gorm:"column:cluster_config;type:json" json:"cluster_config"`

	// 资源使用统计
	ResourceUsage common.JSON `gorm:"column:resource_usage;type:json" json:"resource_usage"`

	// 性能统计信息
	PerformanceStats common.JSON `gorm:"column:performance_stats;type:json" json:"performance_stats"`

	// 告警配置
	AlertConfig common.JSON `gorm:"column:alert_config;type:json" json:"alert_config"`

	// 维护窗口配置
	MaintenanceWindow common.JSON `gorm:"column:maintenance_window;type:json" json:"maintenance_window"`

	// 集群状态 1:正常 0:禁用
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (WorkerCluster) TableName() string {
	return "worker_clusters"
}
