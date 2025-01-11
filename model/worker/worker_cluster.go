package worker

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// WorkerCluster 工作节点集群表
type WorkerCluster struct {
	ClusterID         uint64         `gorm:"column:cluster_id;primaryKey;autoIncrement" json:"cluster_id"`
	ClusterName       string         `gorm:"column:cluster_name;size:100;not null" json:"cluster_name"`
	TotalWorkers      int            `gorm:"column:total_workers;not null" json:"total_workers"`
	ActiveWorkers     int            `gorm:"column:active_workers;not null" json:"active_workers"`
	TotalInstances    int            `gorm:"column:total_instances;not null" json:"total_instances"`
	ActiveInstances   int            `gorm:"column:active_instances;not null" json:"active_instances"`
	ClusterStatus     int8           `gorm:"column:cluster_status;index;not null" json:"cluster_status"`
	ClusterConfig     common.JSON    `gorm:"column:cluster_config;type:json" json:"cluster_config"`
	ResourceUsage     common.JSON    `gorm:"column:resource_usage;type:json" json:"resource_usage"`
	PerformanceStats  common.JSON    `gorm:"column:performance_stats;type:json" json:"performance_stats"`
	AlertConfig       common.JSON    `gorm:"column:alert_config;type:json" json:"alert_config"`
	MaintenanceWindow common.JSON    `gorm:"column:maintenance_window;type:json" json:"maintenance_window"`
	Status            int8           `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt         time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (WorkerCluster) TableName() string {
	return "worker_clusters"
}
