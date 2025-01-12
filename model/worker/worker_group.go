package worker

import (
	"nexus-ai/common"
	"nexus-ai/utils"

	"gorm.io/gorm"
)

// 工作节点信息，包括节点配置、资源限制和扩缩容规则等
type WorkerGroup struct {
	WorkerGroupID          string `gorm:"column:worker_group_id;type:char(36);primaryKey;default:(UUID())" json:"worker_group_id"`                                  // 工作节点组唯一标识
	WorkerClusterID        string `gorm:"column:worker_cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(WorkerClusterID)" json:"worker_cluster_id"` // 关联的集群ID
	WorkerGroupName        string `gorm:"column:worker_group_name;size:100;not null" json:"worker_group_name"`                                                      // 工作节点组名称
	WorkerGroupDescription string `gorm:"column:worker_group_description;type:text" json:"worker_group_description"`                                                // 工作节点组描述

	WorkerGroupType     string      `gorm:"column:worker_group_type;size:50;index;not null" json:"worker_group_type"`     // 节点类型(text/image/video/audio)
	WorkerGroupRole     string      `gorm:"column:worker_group_role;size:20;index;not null" json:"worker_group_role"`     // 节点角色(master/slave)
	WorkerGroupPriority int         `gorm:"column:worker_group_priority;not null;default:0" json:"worker_group_priority"` // 节点优先级
	WorkerGroupOptions  common.JSON `gorm:"column:worker_group_options;type:json" json:"worker_group_options"`            // 节点配置选项

	WorkerGroupStatus           int8        `gorm:"column:worker_group_status;index;not null;default:1" json:"worker_group_status"`                 // 节点状态(1:在线 2:离线 3:维护)
	WorkerGroupLoadBalance      int         `gorm:"column:worker_group_load_balance;not null;default:100" json:"worker_group_load_balance"`         // 负载权重
	WorkerGroupMaxInstances     int         `gorm:"column:worker_group_max_instances;not null;default:10" json:"worker_group_max_instances"`        // 最大实例数
	WorkerGroupMinInstances     int         `gorm:"column:worker_group_min_instances;not null;default:1" json:"worker_group_min_instances"`         // 最小实例数
	WorkerGroupCurrentInstances int         `gorm:"column:worker_group_current_instances;not null;default:0" json:"worker_group_current_instances"` // 当前实例数
	WorkerGroupResourceLimits   common.JSON `gorm:"column:worker_group_resource_limits;type:json" json:"worker_group_resource_limits"`              // 资源限制配置
	WorkerGroupScalingRules     common.JSON `gorm:"column:worker_group_scaling_rules;type:json" json:"worker_group_scaling_rules"`                  // 扩缩容规则
	Status                      int8        `gorm:"column:status;index;not null;default:1" json:"status"`                                           // 节点状态 1:正常 0:禁用

	CreatedAt utils.MySQLTime `gorm:"column:created_at;index;not null" json:"created_at"` // 记录创建时间
	UpdatedAt utils.MySQLTime `gorm:"column:updated_at;not null" json:"updated_at"`       // 记录更新时间
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`                // 软删除时间
}

// TableName 表名
func (WorkerGroup) TableName() string {
	return "worker_groups"
}

// BeforeCreate 在创建记录前自动设置时间
func (workerGroup *WorkerGroup) BeforeCreate(tx *gorm.DB) error {
	workerGroup.CreatedAt = utils.MySQLTime(utils.GetTime())
	workerGroup.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}

// BeforeUpdate 在更新记录前自动设置更新时间
func (workerGroup *WorkerGroup) BeforeUpdate(tx *gorm.DB) error {
	workerGroup.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}
