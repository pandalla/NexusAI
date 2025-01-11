package worker

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// WorkerNode 工作节点实例表
type WorkerNode struct {
	NodeID           uint64         `gorm:"column:node_id;primaryKey;autoIncrement" json:"node_id"`
	WorkerID         uint64         `gorm:"column:worker_id;index;not null" json:"worker_id"`
	ClusterID        uint64         `gorm:"column:cluster_id;index;not null" json:"cluster_id"`
	NodeIP           string         `gorm:"column:node_ip;size:50;not null" json:"node_ip"`
	NodePort         int            `gorm:"column:node_port;not null" json:"node_port"`
	NodeRegion       string         `gorm:"column:node_region;size:50;index" json:"node_region"`
	NodeZone         string         `gorm:"column:node_zone;size:50;index" json:"node_zone"`
	NodeOptions      common.JSON    `gorm:"column:node_options;type:json" json:"node_options"`
	NodeStatus       int8           `gorm:"column:node_status;index;not null" json:"node_status"`
	CPUUsage         float64        `gorm:"column:cpu_usage;type:decimal(5,2)" json:"cpu_usage"`
	MemoryUsage      float64        `gorm:"column:memory_usage;type:decimal(5,2)" json:"memory_usage"`
	NetworkStats     common.JSON    `gorm:"column:network_stats;type:json" json:"network_stats"`
	PerformanceStats common.JSON    `gorm:"column:performance_stats;type:json" json:"performance_stats"`
	HealthCheck      time.Time      `gorm:"column:health_check;not null" json:"health_check"`
	LastError        string         `gorm:"column:last_error;type:text" json:"last_error"`
	StartupTime      time.Time      `gorm:"column:startup_time;not null" json:"startup_time"`
	Status           int8           `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt        time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (WorkerNode) TableName() string {
	return "worker_nodes"
}
