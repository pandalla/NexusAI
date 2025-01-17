package worker

import (
	"fmt"
	"math/rand"
	"nexus-ai/common"
	dto "nexus-ai/dto/model/worker"
	"nexus-ai/model/worker"
	"nexus-ai/utils"
	"time"

	"gorm.io/gorm"
)

// WorkerNodeRepository 工作节点仓储接口
type WorkerNodeRepository interface {
	Create(node *dto.WorkerNode) error
	Update(node *dto.WorkerNode) error
	Delete(nodeID string) error
	GetByID(nodeID string) (*dto.WorkerNode, error)
	List(page, pageSize int) ([]*dto.WorkerNode, int64, error)
	ListByGroup(groupID string, page, pageSize int) ([]*dto.WorkerNode, int64, error)
	ListByCluster(clusterID string, page, pageSize int) ([]*dto.WorkerNode, int64, error)
	ListByStatus(status int8, page, pageSize int) ([]*dto.WorkerNode, int64, error)
	Benchmark(count int) error
}

type workerNodeRepository struct {
	db *gorm.DB
}

// NewWorkerNodeRepository 创建工作节点仓储实例
func NewWorkerNodeRepository(db *gorm.DB) WorkerNodeRepository {
	return &workerNodeRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *workerNodeRepository) convertToDTO(model *worker.WorkerNode) *dto.WorkerNode {
	if model == nil {
		return nil
	}

	var networkStats dto.NetworkStats
	var performanceStats dto.PerformanceNodeStats
	var nodeOptions dto.WorkerNodeOptions

	if err := model.NetworkStats.ToStruct(&networkStats); err != nil {
		utils.SysError("解析网络统计失败: " + err.Error())
	}
	if err := model.PerformanceStats.ToStruct(&performanceStats); err != nil {
		utils.SysError("解析性能统计失败: " + err.Error())
	}
	if err := model.NodeOptions.ToStruct(&nodeOptions); err != nil {
		utils.SysError("解析节点选项失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.WorkerNode{
		WorkerNodeID:     model.WorkerNodeID,
		WorkerGroupID:    model.WorkerGroupID,
		WorkerClusterID:  model.WorkerClusterID,
		NodeIP:           model.NodeIP,
		NodePort:         model.NodePort,
		NodeRegion:       model.NodeRegion,
		NodeZone:         model.NodeZone,
		CPUUsage:         model.CPUUsage,
		MemoryUsage:      model.MemoryUsage,
		NetworkStats:     networkStats,
		PerformanceStats: performanceStats,
		LastError:        model.LastError,
		NodeStatus:       model.NodeStatus,
		NodeOptions:      nodeOptions,
		StartupTime:      model.StartupTime,
		Status:           model.Status,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *workerNodeRepository) convertToModel(dto *dto.WorkerNode) (*worker.WorkerNode, error) {
	if dto == nil {
		return nil, nil
	}

	networkStatsJSON, err := common.FromStruct(dto.NetworkStats)
	if err != nil {
		return nil, fmt.Errorf("转换网络统计失败: %w", err)
	}
	performanceStatsJSON, err := common.FromStruct(dto.PerformanceStats)
	if err != nil {
		return nil, fmt.Errorf("转换性能统计失败: %w", err)
	}
	nodeOptionsJSON, err := common.FromStruct(dto.NodeOptions)
	if err != nil {
		return nil, fmt.Errorf("转换节点选项失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &worker.WorkerNode{
		WorkerNodeID:     dto.WorkerNodeID,
		WorkerGroupID:    dto.WorkerGroupID,
		WorkerClusterID:  dto.WorkerClusterID,
		NodeIP:           dto.NodeIP,
		NodePort:         dto.NodePort,
		NodeRegion:       dto.NodeRegion,
		NodeZone:         dto.NodeZone,
		CPUUsage:         dto.CPUUsage,
		MemoryUsage:      dto.MemoryUsage,
		NetworkStats:     networkStatsJSON,
		PerformanceStats: performanceStatsJSON,
		LastError:        dto.LastError,
		NodeStatus:       dto.NodeStatus,
		NodeOptions:      nodeOptionsJSON,
		StartupTime:      dto.StartupTime,
		Status:           dto.Status,
		CreatedAt:        dto.CreatedAt,
		UpdatedAt:        dto.UpdatedAt,
		DeletedAt:        deletedAt,
	}, nil
}

// Create 创建工作节点
func (r *workerNodeRepository) Create(node *dto.WorkerNode) error {
	model, err := r.convertToModel(node)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新工作节点
func (r *workerNodeRepository) Update(node *dto.WorkerNode) error {
	model, err := r.convertToModel(node)
	if err != nil {
		return err
	}
	return r.db.Model(&worker.WorkerNode{}).Where("worker_node_id = ?", node.WorkerNodeID).Updates(model).Error
}

// Delete 删除工作节点
func (r *workerNodeRepository) Delete(nodeID string) error {
	return r.db.Delete(&worker.WorkerNode{}, "worker_node_id = ?", nodeID).Error
}

// GetByID 根据ID获取工作节点
func (r *workerNodeRepository) GetByID(nodeID string) (*dto.WorkerNode, error) {
	var node worker.WorkerNode
	if err := r.db.Where("worker_node_id = ?", nodeID).First(&node).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&node), nil
}

// List 获取工作节点列表
func (r *workerNodeRepository) List(page, pageSize int) ([]*dto.WorkerNode, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var nodes []worker.WorkerNode

	offset := (page - 1) * pageSize

	if err := r.db.Model(&worker.WorkerNode{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker nodes failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&nodes).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker nodes failed: %w", err)
	}

	dtoList := make([]*dto.WorkerNode, len(nodes))
	for i, n := range nodes {
		dtoList[i] = r.convertToDTO(&n)
	}

	return dtoList, total, nil
}

// ListByGroup 获取指定工作组的节点列表
func (r *workerNodeRepository) ListByGroup(groupID string, page, pageSize int) ([]*dto.WorkerNode, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var nodes []worker.WorkerNode

	offset := (page - 1) * pageSize
	query := r.db.Model(&worker.WorkerNode{}).Where("worker_group_id = ?", groupID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker nodes failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&nodes).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker nodes failed: %w", err)
	}

	dtoList := make([]*dto.WorkerNode, len(nodes))
	for i, n := range nodes {
		dtoList[i] = r.convertToDTO(&n)
	}

	return dtoList, total, nil
}

// ListByCluster 获取指定集群的节点列表
func (r *workerNodeRepository) ListByCluster(clusterID string, page, pageSize int) ([]*dto.WorkerNode, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var nodes []worker.WorkerNode

	offset := (page - 1) * pageSize
	query := r.db.Model(&worker.WorkerNode{}).Where("worker_cluster_id = ?", clusterID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker nodes failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&nodes).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker nodes failed: %w", err)
	}

	dtoList := make([]*dto.WorkerNode, len(nodes))
	for i, n := range nodes {
		dtoList[i] = r.convertToDTO(&n)
	}

	return dtoList, total, nil
}

// ListByStatus 获取指定状态的节点列表
func (r *workerNodeRepository) ListByStatus(status int8, page, pageSize int) ([]*dto.WorkerNode, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var nodes []worker.WorkerNode

	offset := (page - 1) * pageSize
	query := r.db.Model(&worker.WorkerNode{}).Where("status = ?", status)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker nodes failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&nodes).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker nodes failed: %w", err)
	}

	dtoList := make([]*dto.WorkerNode, len(nodes))
	for i, n := range nodes {
		dtoList[i] = r.convertToDTO(&n)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *workerNodeRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行工作节点基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testNode := &dto.WorkerNode{
			WorkerNodeID:    utils.GenerateRandomUUID(12),
			WorkerGroupID:   utils.GenerateRandomUUID(12),
			WorkerClusterID: utils.GenerateRandomUUID(12),
			NodeIP:          fmt.Sprintf("192.168.1.%d", rand.Intn(255)),
			NodePort:        rand.Intn(10000) + 10000,
			NodeRegion:      []string{"cn-east", "cn-south", "cn-north"}[rand.Intn(3)],
			NodeZone:        []string{"zone-a", "zone-b", "zone-c"}[rand.Intn(3)],
			CPUUsage:        rand.Float64() * 100,
			MemoryUsage:     rand.Float64() * 100,
			NetworkStats: dto.NetworkStats{
				NetworkUsage: rand.Float64() * 100,
			},
			PerformanceStats: dto.PerformanceNodeStats{
				CPUUsage:     rand.Float64() * 100,
				MemoryUsage:  rand.Float64() * 100,
				DiskUsage:    rand.Float64() * 100,
				NetworkUsage: rand.Float64() * 100,
			},
			LastError:  "",
			NodeStatus: int8(rand.Intn(5) + 1),
			NodeOptions: dto.WorkerNodeOptions{
				MaxWorkerNodes: rand.Intn(100) + 1,
			},
			StartupTime: utils.MySQLTime(time.Now()),
			Status:      int8(rand.Intn(2) + 1),
		}

		// 创建
		if err := r.Create(testNode); err != nil {
			utils.SysError("创建工作节点失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdNode, err := r.GetByID(testNode.WorkerNodeID)
		if err != nil {
			utils.SysError("获取创建的工作节点失败: " + err.Error())
			return err
		}

		// 更新
		createdNode.Status = int8(rand.Intn(2) + 1)
		if err := r.Update(createdNode); err != nil {
			utils.SysError("更新工作节点失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdNode.WorkerNodeID); err != nil {
			utils.SysError("删除工作节点失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
