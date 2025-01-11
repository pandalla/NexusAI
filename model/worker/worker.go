package worker

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Worker 工作节点表
type Worker struct {
	WorkerID         uint64         `gorm:"column:worker_id;primaryKey;autoIncrement" json:"worker_id"`
	ClusterID        uint64         `gorm:"column:cluster_id;index;not null" json:"cluster_id"`
	WorkerName       string         `gorm:"column:worker_name;size:100;not null" json:"worker_name"`
	WorkerType       string         `gorm:"column:worker_type;size:50;index;not null" json:"worker_type"`
	WorkerGroup      string         `gorm:"column:worker_group;size:50;index;not null" json:"worker_group"`
	WorkerRole       string         `gorm:"column:worker_role;size:20;index;not null" json:"worker_role"`
	WorkerPriority   int            `gorm:"column:worker_priority;not null" json:"worker_priority"`
	WorkerOptions    common.JSON    `gorm:"column:worker_options;type:json" json:"worker_options"`
	WorkerStatus     int8           `gorm:"column:worker_status;index;not null" json:"worker_status"`
	LoadBalance      int            `gorm:"column:load_balance;not null" json:"load_balance"`
	MaxInstances     int            `gorm:"column:max_instances;not null" json:"max_instances"`
	MinInstances     int            `gorm:"column:min_instances;not null" json:"min_instances"`
	CurrentInstances int            `gorm:"column:current_instances;not null" json:"current_instances"`
	ResourceLimits   common.JSON    `gorm:"column:resource_limits;type:json" json:"resource_limits"`
	ScalingRules     common.JSON    `gorm:"column:scaling_rules;type:json" json:"scaling_rules"`
	Status           int8           `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt        time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 表名
func (Worker) TableName() string {
	return "workers"
}
