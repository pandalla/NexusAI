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

// WorkerGroupRepository 工作组仓储接口
type WorkerGroupRepository interface {
	Create(group *dto.WorkerGroup) error
	Update(group *dto.WorkerGroup) error
	Delete(groupID string) error
	GetByID(groupID string) (*dto.WorkerGroup, error)
	List(page, pageSize int) ([]*dto.WorkerGroup, int64, error)
	ListByCluster(clusterID string, page, pageSize int) ([]*dto.WorkerGroup, int64, error)
	ListByStatus(status int8, page, pageSize int) ([]*dto.WorkerGroup, int64, error)
	Benchmark(count int) error
}

type workerGroupRepository struct {
	db *gorm.DB
}

// NewWorkerGroupRepository 创建工作组仓储实例
func NewWorkerGroupRepository(db *gorm.DB) WorkerGroupRepository {
	return &workerGroupRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *workerGroupRepository) convertToDTO(model *worker.WorkerGroup) *dto.WorkerGroup {
	if model == nil {
		return nil
	}

	var workerGroupOptions dto.WorkerGroupOptions
	var resourceLimits dto.WorkerGroupResourceLimits
	var scalingRules dto.WorkerGroupScalingRules

	if err := model.WorkerGroupOptions.ToStruct(&workerGroupOptions); err != nil {
		utils.SysError("解析工作组选项失败: " + err.Error())
	}
	if err := model.WorkerGroupResourceLimits.ToStruct(&resourceLimits); err != nil {
		utils.SysError("解析资源限制失败: " + err.Error())
	}
	if err := model.WorkerGroupScalingRules.ToStruct(&scalingRules); err != nil {
		utils.SysError("解析扩缩规则失败: " + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.WorkerGroup{
		WorkerGroupID:               model.WorkerGroupID,
		WorkerClusterID:             model.WorkerClusterID,
		WorkerGroupName:             model.WorkerGroupName,
		WorkerGroupDescription:      model.WorkerGroupDescription,
		WorkerGroupType:             model.WorkerGroupType,
		WorkerGroupRole:             model.WorkerGroupRole,
		WorkerGroupPriority:         model.WorkerGroupPriority,
		WorkerGroupOptions:          workerGroupOptions,
		WorkerGroupStatus:           model.WorkerGroupStatus,
		WorkerGroupLoadBalance:      model.WorkerGroupLoadBalance,
		WorkerGroupMaxInstances:     model.WorkerGroupMaxInstances,
		WorkerGroupMinInstances:     model.WorkerGroupMinInstances,
		WorkerGroupCurrentInstances: model.WorkerGroupCurrentInstances,
		WorkerGroupResourceLimits:   resourceLimits,
		WorkerGroupScalingRules:     scalingRules,
		Status:                      model.Status,
		CreatedAt:                   model.CreatedAt,
		UpdatedAt:                   model.UpdatedAt,
		DeletedAt:                   deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *workerGroupRepository) convertToModel(dto *dto.WorkerGroup) (*worker.WorkerGroup, error) {
	if dto == nil {
		return nil, nil
	}

	optionsJSON, err := common.FromStruct(dto.WorkerGroupOptions)
	if err != nil {
		return nil, fmt.Errorf("转换工作组选项失败: %w", err)
	}
	resourceLimitsJSON, err := common.FromStruct(dto.WorkerGroupResourceLimits)
	if err != nil {
		return nil, fmt.Errorf("转换资源限制失败: %w", err)
	}
	scalingRulesJSON, err := common.FromStruct(dto.WorkerGroupScalingRules)
	if err != nil {
		return nil, fmt.Errorf("转换扩缩规则失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &worker.WorkerGroup{
		WorkerGroupID:               dto.WorkerGroupID,
		WorkerClusterID:             dto.WorkerClusterID,
		WorkerGroupName:             dto.WorkerGroupName,
		WorkerGroupDescription:      dto.WorkerGroupDescription,
		WorkerGroupType:             dto.WorkerGroupType,
		WorkerGroupRole:             dto.WorkerGroupRole,
		WorkerGroupPriority:         dto.WorkerGroupPriority,
		WorkerGroupOptions:          optionsJSON,
		WorkerGroupStatus:           dto.WorkerGroupStatus,
		WorkerGroupLoadBalance:      dto.WorkerGroupLoadBalance,
		WorkerGroupMaxInstances:     dto.WorkerGroupMaxInstances,
		WorkerGroupMinInstances:     dto.WorkerGroupMinInstances,
		WorkerGroupCurrentInstances: dto.WorkerGroupCurrentInstances,
		WorkerGroupResourceLimits:   resourceLimitsJSON,
		WorkerGroupScalingRules:     scalingRulesJSON,
		Status:                      dto.Status,
		CreatedAt:                   dto.CreatedAt,
		UpdatedAt:                   dto.UpdatedAt,
		DeletedAt:                   deletedAt,
	}, nil
}

// Create 创建工作组
func (r *workerGroupRepository) Create(group *dto.WorkerGroup) error {
	model, err := r.convertToModel(group)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新工作组
func (r *workerGroupRepository) Update(group *dto.WorkerGroup) error {
	model, err := r.convertToModel(group)
	if err != nil {
		return err
	}
	return r.db.Model(&worker.WorkerGroup{}).Where("worker_group_id = ?", group.WorkerGroupID).Updates(model).Error
}

// Delete 删除工作组
func (r *workerGroupRepository) Delete(groupID string) error {
	return r.db.Delete(&worker.WorkerGroup{}, "worker_group_id = ?", groupID).Error
}

// GetByID 根据ID获取工作组
func (r *workerGroupRepository) GetByID(groupID string) (*dto.WorkerGroup, error) {
	var group worker.WorkerGroup
	if err := r.db.Where("worker_group_id = ?", groupID).First(&group).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&group), nil
}

// List 获取工作组列表
func (r *workerGroupRepository) List(page, pageSize int) ([]*dto.WorkerGroup, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var groups []worker.WorkerGroup

	offset := (page - 1) * pageSize

	if err := r.db.Model(&worker.WorkerGroup{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker groups failed: %w", err)
	}

	if err := r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&groups).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker groups failed: %w", err)
	}

	dtoList := make([]*dto.WorkerGroup, len(groups))
	for i, g := range groups {
		dtoList[i] = r.convertToDTO(&g)
	}

	return dtoList, total, nil
}

// ListByCluster 获取指定集群的工作组列表
func (r *workerGroupRepository) ListByCluster(clusterID string, page, pageSize int) ([]*dto.WorkerGroup, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var groups []worker.WorkerGroup

	offset := (page - 1) * pageSize
	query := r.db.Model(&worker.WorkerGroup{}).Where("worker_cluster_id = ?", clusterID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker groups failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&groups).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker groups failed: %w", err)
	}

	dtoList := make([]*dto.WorkerGroup, len(groups))
	for i, g := range groups {
		dtoList[i] = r.convertToDTO(&g)
	}

	return dtoList, total, nil
}

// ListByStatus 获取指定状态的工作组列表
func (r *workerGroupRepository) ListByStatus(status int8, page, pageSize int) ([]*dto.WorkerGroup, int64, error) {
	if page < 1 || pageSize < 1 {
		return nil, 0, fmt.Errorf("invalid pagination parameters: page=%d, pageSize=%d", page, pageSize)
	}

	var total int64
	var groups []worker.WorkerGroup

	offset := (page - 1) * pageSize
	query := r.db.Model(&worker.WorkerGroup{}).Where("status = ?", status)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count worker groups failed: %w", err)
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&groups).Error; err != nil {
		return nil, 0, fmt.Errorf("find worker groups failed: %w", err)
	}

	dtoList := make([]*dto.WorkerGroup, len(groups))
	for i, g := range groups {
		dtoList[i] = r.convertToDTO(&g)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *workerGroupRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行工作组基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testGroup := &dto.WorkerGroup{
			WorkerGroupID:          utils.GenerateRandomUUID(12),
			WorkerClusterID:        utils.GenerateRandomUUID(12),
			WorkerGroupName:        fmt.Sprintf("test_group_%d", i),
			WorkerGroupDescription: fmt.Sprintf("测试工作组 %d", i),
			WorkerGroupType:        []string{"compute", "storage", "network"}[rand.Intn(3)],
			WorkerGroupRole:        []string{"master", "worker"}[rand.Intn(2)],
			WorkerGroupPriority:    rand.Intn(10),
			WorkerGroupOptions: dto.WorkerGroupOptions{
				MaxWorkerNodes: rand.Intn(100) + 1,
			},
			WorkerGroupStatus:           int8(rand.Intn(5) + 1),
			WorkerGroupLoadBalance:      rand.Intn(100),
			WorkerGroupMaxInstances:     rand.Intn(100) + 10,
			WorkerGroupMinInstances:     rand.Intn(10),
			WorkerGroupCurrentInstances: rand.Intn(50) + 1,
			WorkerGroupResourceLimits: dto.WorkerGroupResourceLimits{
				CPUUsage:     rand.Float64() * 100,
				MemoryUsage:  rand.Float64() * 100,
				DiskUsage:    rand.Float64() * 100,
				NetworkUsage: rand.Float64() * 100,
			},
			WorkerGroupScalingRules: dto.WorkerGroupScalingRules{
				ScalingType:  []string{"cpu", "memory", "request"}[rand.Intn(3)],
				ScalingValue: rand.Intn(100),
			},
			Status: int8(rand.Intn(2) + 1),
		}

		// 创建
		if err := r.Create(testGroup); err != nil {
			utils.SysError("创建工作组失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdGroup, err := r.GetByID(testGroup.WorkerGroupID)
		if err != nil {
			utils.SysError("获取创建的工作组失败: " + err.Error())
			return err
		}

		// 更新
		createdGroup.Status = int8(rand.Intn(2) + 1)
		if err := r.Update(createdGroup); err != nil {
			utils.SysError("更新工作组失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdGroup.WorkerGroupID); err != nil {
			utils.SysError("删除工作组失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
