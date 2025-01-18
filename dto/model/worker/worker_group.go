package worker

import (
	"nexus-ai/utils"
)

type WorkerGroupOptions struct {
	MaxWorkerNodes int `json:"max_worker_nodes"`
}
type WorkerGroupResourceLimits struct {
	CPUUsage     float64 `json:"cpu_usage"`
	MemoryUsage  float64 `json:"memory_usage"`
	DiskUsage    float64 `json:"disk_usage"`
	NetworkUsage float64 `json:"network_usage"`
}
type WorkerGroupScalingRules struct {
	ScalingType  string `json:"scaling_type"`
	ScalingValue int    `json:"scaling_value"`
}

type WorkerGroup struct {
	WorkerGroupID          string `json:"worker_group_id"`
	WorkerClusterID        string `json:"worker_cluster_id"`
	WorkerGroupName        string `json:"worker_group_name"`
	WorkerGroupDescription string `json:"worker_group_description"`

	WorkerGroupType     string             `json:"worker_group_type"`
	WorkerGroupRole     string             `json:"worker_group_role"`
	WorkerGroupPriority int                `json:"worker_group_priority"`
	WorkerGroupOptions  WorkerGroupOptions `json:"worker_group_options"`

	WorkerGroupStatus           int8                      `json:"worker_group_status"`
	WorkerGroupLoadBalance      int                       `json:"worker_group_load_balance"`
	WorkerGroupMaxInstances     int                       `json:"worker_group_max_instances"`
	WorkerGroupMinInstances     int                       `json:"worker_group_min_instances"`
	WorkerGroupCurrentInstances int                       `json:"worker_group_current_instances"`
	WorkerGroupResourceLimits   WorkerGroupResourceLimits `json:"worker_group_resource_limits"`
	WorkerGroupScalingRules     WorkerGroupScalingRules   `json:"worker_group_scaling_rules"`
	Status                      int8                      `json:"status"`

	CreatedAt utils.MySQLTime  `json:"created_at"`
	UpdatedAt utils.MySQLTime  `json:"updated_at"`
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"`
}
