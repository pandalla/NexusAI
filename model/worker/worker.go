package worker

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// Worker 存储系统中的工作节点信息，包括节点配置、资源限制和扩缩容规则等
type Worker struct {
	// 工作节点唯一标识
	WorkerID string `gorm:"column:worker_id;type:char(36);primaryKey;default:(UUID())" json:"worker_id"`

	// 关联的集群ID
	ClusterID string `gorm:"column:cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(ClusterID)" json:"cluster_id"`

	// 工作节点名称
	WorkerName string `gorm:"column:worker_name;size:100;not null" json:"worker_name"`

	// 节点类型(text/image/video/audio)
	WorkerType string `gorm:"column:worker_type;size:50;index;not null" json:"worker_type"`

	// 节点组
	WorkerGroup string `gorm:"column:worker_group;size:50;index;not null" json:"worker_group"`

	// 节点角色(master/slave)
	WorkerRole string `gorm:"column:worker_role;size:20;index;not null" json:"worker_role"`

	// 节点优先级
	WorkerPriority int `gorm:"column:worker_priority;not null;default:0" json:"worker_priority"`

	// 节点配置选项
	WorkerOptions common.JSON `gorm:"column:worker_options;type:json" json:"worker_options"`

	// 节点状态(1:在线 2:离线 3:维护)
	WorkerStatus int8 `gorm:"column:worker_status;index;not null;default:1" json:"worker_status"`

	// 负载权重
	LoadBalance int `gorm:"column:load_balance;not null;default:100" json:"load_balance"`

	// 最大实例数
	MaxInstances int `gorm:"column:max_instances;not null;default:10" json:"max_instances"`

	// 最小实例数
	MinInstances int `gorm:"column:min_instances;not null;default:1" json:"min_instances"`

	// 当前实例数
	CurrentInstances int `gorm:"column:current_instances;not null;default:0" json:"current_instances"`

	// 资源限制配置
	ResourceLimits common.JSON `gorm:"column:resource_limits;type:json" json:"resource_limits"`

	// 扩缩容规则
	ScalingRules common.JSON `gorm:"column:scaling_rules;type:json" json:"scaling_rules"`

	// 节点状态 1:正常 0:禁用
	Status int8 `gorm:"column:status;index;not null;default:1" json:"status"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`

	// 软删除时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (Worker) TableName() string {
	return "workers"
}
