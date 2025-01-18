package worker

import (
	"nexus-ai/utils"
)

type WorkerClusterOptions struct {
	MaxWorkerGroups int `json:"max_worker_groups"`
	MaxWorkerNodes  int `json:"max_worker_nodes"`
}
type ResourceUsage struct {
	CPUUsage     float64 `json:"cpu_usage"`
	MemoryUsage  float64 `json:"memory_usage"`
	DiskUsage    float64 `json:"disk_usage"`
	NetworkUsage float64 `json:"network_usage"`
}
type PerformanceClusterStats struct {
	CPUUsage     float64 `json:"cpu_usage"`
	MemoryUsage  float64 `json:"memory_usage"`
	DiskUsage    float64 `json:"disk_usage"`
	NetworkUsage float64 `json:"network_usage"`
}
type AlertOptions struct {
	AlertType    string `json:"alert_type"`
	AlertLevel   string `json:"alert_level"`
	AlertMessage string `json:"alert_message"`
}
type MaintenanceWindow struct {
	StartTime utils.MySQLTime `json:"start_time"`
	EndTime   utils.MySQLTime `json:"end_time"`
}

type WorkerCluster struct {
	WorkerClusterID    string `json:"worker_cluster_id"`
	WorkerClusterName  string `json:"worker_cluster_name"`
	TotalWorkerGroups  int    `json:"total_worker_groups"`
	ActiveWorkerGroups int    `json:"active_worker_groups"`
	TotalWorkerNodes   int    `json:"total_worker_nodes"`
	ActiveWorkerNodes  int    `json:"active_worker_nodes"`

	WorkerClusterStatus  int8                    `json:"worker_cluster_status"`
	WorkerClusterOptions WorkerClusterOptions    `json:"worker_cluster_options"`
	ResourceUsage        ResourceUsage           `json:"resource_usage"`
	PerformanceStats     PerformanceClusterStats `json:"performance_stats"`
	AlertOptions         AlertOptions            `json:"alert_options"`
	MaintenanceWindow    MaintenanceWindow       `json:"maintenance_window"`
	Status               int8                    `json:"status"`

	CreatedAt utils.MySQLTime  `json:"created_at"`
	UpdatedAt utils.MySQLTime  `json:"updated_at"`
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"`
}
