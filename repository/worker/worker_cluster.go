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

// WorkerClusterRepository 工作集群仓储接口
type WorkerClusterRepository interface {
	Create(cluster *dto.WorkerCluster) error
	Update(cluster *dto.WorkerCluster) error
	Delete(clusterID string) error
	GetByID(clusterID string) (*dto.WorkerCluster, error)
	List(page, pageSize int) ([]*dto.WorkerCluster, int64, error)
	ListByStatus(status int8, page, pageSize int) ([]*dto.WorkerCluster, int64, error)
	Benchmark(count int) error
}

type workerClusterRepository struct {
	db *gorm.DB
}

// NewWorkerClusterRepository 创建工作集群仓储实例
func NewWorkerClusterRepository(db *gorm.DB) WorkerClusterRepository {
	return &workerClusterRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *workerClusterRepository) convertToDTO(model *worker.WorkerCluster) *dto.WorkerCluster {
	if model == nil {
		return nil
	}

	var workerClusterOptions dto.WorkerClusterOptions
	var resourceUsage dto.ResourceUsage
	var performanceStats dto.PerformanceClusterStats
	var alertOptions dto.AlertOptions
	var maintenanceWindow dto.MaintenanceWindow

	if err := model.WorkerClusterOptions.ToStruct(&workerClusterOptions); err != nil {
		utils.SysError("解析工作集群选项失败: " + err.Error())
	}
	if err := model.ResourceUsage.ToStruct(&resourceUsage); err != nil {
		utils.SysError("解析资源使用情况失败: " + err.Error())
	}
	if err := model.PerformanceStats.ToStruct(&performanceStats); err != nil {
		utils.SysError("解析性能统计失败: " + err.Error())
	}
	if err := model.AlertOptions.ToStruct(&alertOptions); err != nil {
		utils.SysError("解析告警选项失败: " + err.Error())
	}
	if err := model.MaintenanceWindow.ToStruct(&maintenanceWindow); err != nil {
		utils.SysError("解析维护窗口失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.WorkerCluster{
		WorkerClusterID:      model.WorkerClusterID,
		WorkerClusterName:    model.WorkerClusterName,
		TotalWorkerGroups:    model.TotalWorkerGroups,
		ActiveWorkerGroups:   model.ActiveWorkerGroups,
		TotalWorkerNodes:     model.TotalWorkerNodes,
		ActiveWorkerNodes:    model.ActiveWorkerNodes,
		WorkerClusterStatus:  model.WorkerClusterStatus,
		WorkerClusterOptions: workerClusterOptions,
		ResourceUsage:        resourceUsage,
		PerformanceStats:     performanceStats,
		AlertOptions:         alertOptions,
		MaintenanceWindow:    maintenanceWindow,
		Status:               model.Status,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
		DeletedAt:            deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *workerClusterRepository) convertToModel(dto *dto.WorkerCluster) (*worker.WorkerCluster, error) {
	if dto == nil {
		return nil, nil
	}

	workerClusterOptionsJSON, err := common.FromStruct(dto.WorkerClusterOptions)
	if err != nil {
		return nil, fmt.Errorf("转换工作集群选项失败: %w", err)
	}
	resourceUsageJSON, err := common.FromStruct(dto.ResourceUsage)
	if err != nil {
		return nil, fmt.Errorf("转换资源使用情况失败: %w", err)
	}
	performanceStatsJSON, err := common.FromStruct(dto.PerformanceStats)
	if err != nil {
		return nil, fmt.Errorf("转换性能统计失败: %w", err)
	}
	alertOptionsJSON, err := common.FromStruct(dto.AlertOptions)
	if err != nil {
		return nil, fmt.Errorf("转换告警选项失败: %w", err)
	}
	maintenanceWindowJSON, err := common.FromStruct(dto.MaintenanceWindow)
	if err != nil {
		return nil, fmt.Errorf("转换维护窗口失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &worker.WorkerCluster{
		WorkerClusterID:      dto.WorkerClusterID,
		WorkerClusterName:    dto.WorkerClusterName,
		TotalWorkerGroups:    dto.TotalWorkerGroups,
		ActiveWorkerGroups:   dto.ActiveWorkerGroups,
		TotalWorkerNodes:     dto.TotalWorkerNodes,
		ActiveWorkerNodes:    dto.ActiveWorkerNodes,
		WorkerClusterStatus:  dto.WorkerClusterStatus,
		WorkerClusterOptions: workerClusterOptionsJSON,
		ResourceUsage:        resourceUsageJSON,
		PerformanceStats:     performanceStatsJSON,
		AlertOptions:         alertOptionsJSON,
		MaintenanceWindow:    maintenanceWindowJSON,
		Status:               dto.Status,
		CreatedAt:            dto.CreatedAt,
		UpdatedAt:            dto.UpdatedAt,
		DeletedAt:            deletedAt,
	}, nil
}

// Create 创建工作集群
func (r *workerClusterRepository) Create(cluster *dto.WorkerCluster) error {
	model, err := r.convertToModel(cluster)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新工作集群
func (r *workerClusterRepository) Update(cluster *dto.WorkerCluster) error {
	model, err := r.convertToModel(cluster)
	if err != nil {
		return err
	}
	return r.db.Model(&worker.WorkerCluster{}).Where("worker_cluster_id = ?", cluster.WorkerClusterID).Updates(model).Error
}

// Delete 删除工作集群
func (r *workerClusterRepository) Delete(clusterID string) error {
	return r.db.Delete(&worker.WorkerCluster{}, "worker_cluster_id = ?", clusterID).Error
}

// GetByID 根据ID获取工作集群
func (r *workerClusterRepository) GetByID(clusterID string) (*dto.WorkerCluster, error) {
	var cluster worker.WorkerCluster
	if err := r.db.Where("worker_cluster_id = ?", clusterID).First(&cluster).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&cluster), nil
}

// List 获取工作集群列表
func (r *workerClusterRepository) List(page, pageSize int) ([]*dto.WorkerCluster, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var clusters []worker.WorkerCluster

	offset := (page - 1) * pageSize

	if err := r.db.Model(&worker.WorkerCluster{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker clusters failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&clusters).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker clusters failed: %w", err)
	}

	dtoList := make([]*dto.WorkerCluster, len(clusters))
	for i, c := range clusters {
		dtoList[i] = r.convertToDTO(&c)
	}

	return dtoList, total, nil
}

// ListByStatus 获取指定状态的工作集群列表
func (r *workerClusterRepository) ListByStatus(status int8, page, pageSize int) ([]*dto.WorkerCluster, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var clusters []worker.WorkerCluster

	offset := (page - 1) * pageSize
	query := r.db.Model(&worker.WorkerCluster{}).Where("status = ?", status)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker clusters failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&clusters).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker clusters failed: %w", err)
	}

	dtoList := make([]*dto.WorkerCluster, len(clusters))
	for i, c := range clusters {
		dtoList[i] = r.convertToDTO(&c)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *workerClusterRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行工作集群基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testCluster := &dto.WorkerCluster{
			WorkerClusterID:     utils.GenerateRandomUUID(12),
			WorkerClusterName:   fmt.Sprintf("test_cluster_%d", i),
			TotalWorkerGroups:   rand.Intn(10) + 1,
			ActiveWorkerGroups:  rand.Intn(10) + 1,
			TotalWorkerNodes:    rand.Intn(100) + 1,
			ActiveWorkerNodes:   rand.Intn(100) + 1,
			WorkerClusterStatus: int8(rand.Intn(5) + 1),
			WorkerClusterOptions: dto.WorkerClusterOptions{
				MaxWorkerGroups: rand.Intn(10) + 1,
				MaxWorkerNodes:  rand.Intn(100) + 1,
			},
			ResourceUsage: dto.ResourceUsage{
				CPUUsage:     rand.Float64() * 100,
				MemoryUsage:  rand.Float64() * 100,
				DiskUsage:    rand.Float64() * 100,
				NetworkUsage: rand.Float64() * 100,
			},
			PerformanceStats: dto.PerformanceClusterStats{
				CPUUsage:     rand.Float64() * 100,
				MemoryUsage:  rand.Float64() * 100,
				DiskUsage:    rand.Float64() * 100,
				NetworkUsage: rand.Float64() * 100,
			},
			AlertOptions: dto.AlertOptions{
				AlertType:    "email",
				AlertLevel:   "high",
				AlertMessage: "Alert message",
			},
			MaintenanceWindow: dto.MaintenanceWindow{
				StartTime: utils.MySQLTime(time.Now().Add(24 * time.Hour)),
				EndTime:   utils.MySQLTime(time.Now().Add(48 * time.Hour)),
			},
			Status: int8(rand.Intn(2) + 1),
		}

		// 创建
		if err := r.Create(testCluster); err != nil {
			utils.SysError("创建工作集群失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdCluster, err := r.GetByID(testCluster.WorkerClusterID)
		if err != nil {
			utils.SysError("获取创建的工作集群失败: " + err.Error())
			return err
		}

		// 更新
		createdCluster.Status = int8(rand.Intn(2) + 1)
		if err := r.Update(createdCluster); err != nil {
			utils.SysError("更新工作集群失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdCluster.WorkerClusterID); err != nil {
			utils.SysError("删除工作集群失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
