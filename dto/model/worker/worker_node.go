package worker

import (
	"nexus-ai/utils"
)

type NetworkStats struct {
	NetworkUsage float64 `json:"network_usage"`
}
type PerformanceNodeStats struct {
	CPUUsage     float64 `json:"cpu_usage"`
	MemoryUsage  float64 `json:"memory_usage"`
	DiskUsage    float64 `json:"disk_usage"`
	NetworkUsage float64 `json:"network_usage"`
}
type WorkerNodeOptions struct {
	MaxWorkerNodes int `json:"max_worker_nodes"`
}

type WorkerNode struct {
	WorkerNodeID    string `json:"worker_node_id"`
	WorkerGroupID   string `json:"worker_group_id"`
	WorkerClusterID string `json:"worker_cluster_id"`

	NodeIP           string               `json:"node_ip"`
	NodePort         int                  `json:"node_port"`
	NodeRegion       string               `json:"node_region"`
	NodeZone         string               `json:"node_zone"`
	CPUUsage         float64              `json:"cpu_usage"`
	MemoryUsage      float64              `json:"memory_usage"`
	NetworkStats     NetworkStats         `json:"network_stats"`
	PerformanceStats PerformanceNodeStats `json:"performance_stats"`
	LastError        string               `json:"last_error"`
	NodeStatus       int8                 `json:"node_status"`
	NodeOptions      WorkerNodeOptions    `json:"node_options"`

	StartupTime utils.MySQLTime `json:"startup_time"`
	Status      int8            `json:"status"`

	CreatedAt utils.MySQLTime  `json:"created_at"`
	UpdatedAt utils.MySQLTime  `json:"updated_at"`
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"`
}
