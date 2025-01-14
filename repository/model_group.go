package repository

import (
	"fmt"
	"math/rand"
	"nexus-ai/common"
	dto "nexus-ai/dto/model"
	"nexus-ai/model"
	"nexus-ai/utils"
	"time"

	"gorm.io/gorm"
)

// ModelGroupRepository 模型组仓储接口
type ModelGroupRepository interface {
	Create(modelGroup *dto.ModelGroup) error
	Update(modelGroup *dto.ModelGroup) error
	Delete(modelGroupID string) error
	GetByID(modelGroupID string) (*dto.ModelGroup, error)
	GetByName(name string) (*dto.ModelGroup, error)
	List(page, pageSize int) ([]*dto.ModelGroup, int64, error)
	Benchmark(count int) error
}

type modelGroupRepository struct {
	db *gorm.DB
}

// NewModelGroupRepository 创建模型组仓储实例
func NewModelGroupRepository(db *gorm.DB) ModelGroupRepository {
	return &modelGroupRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *modelGroupRepository) convertToDTO(model *model.ModelGroup) *dto.ModelGroup {
	if model == nil {
		return nil
	}

	var priceFactor dto.ModelGroupPriceFactor
	var options dto.ModelGroupOptions

	if err := model.ModelGroupPriceFactor.ToStruct(&priceFactor); err != nil {
		utils.SysError("解析价格系数失败:" + err.Error())
	}

	if err := model.ModelGroupOptions.ToStruct(&options); err != nil {
		utils.SysError("解析配置选项失败:" + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.ModelGroup{
		ModelGroupID:          model.ModelGroupID,
		ModelGroupName:        model.ModelGroupName,
		ModelGroupDescription: model.ModelGroupDescription,
		ModelGroupPriceFactor: priceFactor,
		ModelGroupOptions:     options,
		CreatedAt:             model.CreatedAt,
		UpdatedAt:             model.UpdatedAt,
		DeletedAt:             deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *modelGroupRepository) convertToModel(dto *dto.ModelGroup) (*model.ModelGroup, error) {
	if dto == nil {
		return nil, nil
	}

	priceFactorJSON, err := common.FromStruct(dto.ModelGroupPriceFactor)
	if err != nil {
		return nil, fmt.Errorf("转换价格系数失败: %w", err)
	}

	optionsJSON, err := common.FromStruct(dto.ModelGroupOptions)
	if err != nil {
		return nil, fmt.Errorf("转换配置选项失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &model.ModelGroup{
		ModelGroupID:          dto.ModelGroupID,
		ModelGroupName:        dto.ModelGroupName,
		ModelGroupDescription: dto.ModelGroupDescription,
		ModelGroupPriceFactor: priceFactorJSON,
		ModelGroupOptions:     optionsJSON,
		CreatedAt:             dto.CreatedAt,
		UpdatedAt:             dto.UpdatedAt,
		DeletedAt:             deletedAt,
	}, nil
}

// Create 创建模型组
func (r *modelGroupRepository) Create(modelGroup *dto.ModelGroup) error {
	model, err := r.convertToModel(modelGroup)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新模型组
func (r *modelGroupRepository) Update(modelGroup *dto.ModelGroup) error {
	modelData, err := r.convertToModel(modelGroup)
	if err != nil {
		return err
	}
	return r.db.Model(&model.ModelGroup{}).Where("model_group_id = ?", modelGroup.ModelGroupID).Updates(modelData).Error
}

// Delete 删除模型组
func (r *modelGroupRepository) Delete(modelGroupID string) error {
	return r.db.Delete(&model.ModelGroup{}, "model_group_id = ?", modelGroupID).Error
}

// GetByID 根据ID获取模型组
func (r *modelGroupRepository) GetByID(modelGroupID string) (*dto.ModelGroup, error) {
	var modelGroup model.ModelGroup
	if err := r.db.Where("model_group_id = ?", modelGroupID).First(&modelGroup).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&modelGroup), nil
}

// GetByName 根据名称获取模型组
func (r *modelGroupRepository) GetByName(name string) (*dto.ModelGroup, error) {
	var modelGroup model.ModelGroup
	if err := r.db.Where("model_group_name = ?", name).First(&modelGroup).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&modelGroup), nil
}

// List 获取模型组列表
func (r *modelGroupRepository) List(page, pageSize int) ([]*dto.ModelGroup, int64, error) {
	var total int64
	var modelGroups []model.ModelGroup

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.ModelGroup{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&modelGroups).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.ModelGroup, len(modelGroups))
	for i, mg := range modelGroups {
		dtoList[i] = r.convertToDTO(&mg)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *modelGroupRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行模型组基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testModelGroup := &dto.ModelGroup{
			ModelGroupID:          utils.GenerateRandomUUID(12),
			ModelGroupName:        fmt.Sprintf("benchmark_model_group_%d", i),
			ModelGroupDescription: fmt.Sprintf("这是第%d个基准测试模型组", i),
			ModelGroupPriceFactor: dto.ModelGroupPriceFactor{
				RequestPriceFactor:    float64(rand.Intn(50)+50) / 100,
				ResponsePriceFactor:   float64(rand.Intn(50)+50) / 100,
				CompletionPriceFactor: float64(rand.Intn(50)+50) / 100,
			},
			ModelGroupOptions: dto.ModelGroupOptions{
				MaxConcurrentRequests: rand.Intn(10) + 1,
				DefaultLevel:          rand.Intn(3) + 1,
				ExtraAllowedModels:    []string{"gpt-3.5-turbo", "gpt-4"},
				APIDiscount:           float64(rand.Intn(50)+50) / 100,
			},
		}

		// 创建
		if err := r.Create(testModelGroup); err != nil {
			utils.SysError("创建模型组失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdModelGroup, err := r.GetByID(testModelGroup.ModelGroupID)
		if err != nil {
			utils.SysError("获取创建的模型组失败: " + err.Error())
			return err
		}

		// 更新
		createdModelGroup.ModelGroupOptions.DefaultLevel = rand.Intn(5) + 1
		if err := r.Update(createdModelGroup); err != nil {
			utils.SysError("更新模型组失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdModelGroup.ModelGroupID); err != nil {
			utils.SysError("删除模型组失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
