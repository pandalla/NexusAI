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

// UserGroupRepository 用户组仓储接口
type UserGroupRepository interface {
	Create(userGroup *dto.UserGroup) error
	Update(userGroup *dto.UserGroup) error
	Delete(userGroupID string) error
	GetByID(userGroupID string) (*dto.UserGroup, error)
	GetByName(name string) (*dto.UserGroup, error)
	List(page, pageSize int) ([]*dto.UserGroup, int64, error)
	ListByPriceFactor(factor dto.UserGroupPriceFactor, page, pageSize int) ([]*dto.UserGroup, int64, error)
	ListByOptions(options dto.UserGroupOptions, page, pageSize int) ([]*dto.UserGroup, int64, error)
	Benchmark(count int) error
}

type userGroupRepository struct {
	db *gorm.DB
}

// NewUserGroupRepository 创建用户组仓储实例
func NewUserGroupRepository(db *gorm.DB) UserGroupRepository {
	return &userGroupRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *userGroupRepository) convertToDTO(model *model.UserGroup) *dto.UserGroup {
	if model == nil {
		return nil
	}

	var priceFactor dto.UserGroupPriceFactor
	var options dto.UserGroupOptions

	if err := model.UserGroupPriceFactor.ToStruct(&priceFactor); err != nil {
		utils.SysError("解析价格系数失败:" + err.Error())
	}

	if err := model.UserGroupOptions.ToStruct(&options); err != nil {
		utils.SysError("解析配置选项失败:" + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.UserGroup{
		UserGroupID:          model.UserGroupID,
		UserGroupName:        model.UserGroupName,
		UserGroupDescription: model.UserGroupDescription,
		UserGroupPriceFactor: priceFactor,
		UserGroupOptions:     options,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
		DeletedAt:            deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *userGroupRepository) convertToModel(dto *dto.UserGroup) (*model.UserGroup, error) {
	if dto == nil {
		return nil, nil
	}

	priceFactorJSON, err := common.FromStruct(dto.UserGroupPriceFactor)
	if err != nil {
		return nil, fmt.Errorf("转换价格系数失败: %w", err)
	}

	optionsJSON, err := common.FromStruct(dto.UserGroupOptions)
	if err != nil {
		return nil, fmt.Errorf("转换配置选项失败: %w", err)
	}

	return &model.UserGroup{
		UserGroupID:          dto.UserGroupID,
		UserGroupName:        dto.UserGroupName,
		UserGroupDescription: dto.UserGroupDescription,
		UserGroupPriceFactor: priceFactorJSON,
		UserGroupOptions:     optionsJSON,
	}, nil
}

// Create 创建用户组
func (r *userGroupRepository) Create(userGroup *dto.UserGroup) error {
	model, err := r.convertToModel(userGroup)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新用户组
func (r *userGroupRepository) Update(userGroup *dto.UserGroup) error {
	modelData, err := r.convertToModel(userGroup)
	if err != nil {
		return err
	}
	return r.db.Model(&model.UserGroup{}).Where("user_group_id = ?", userGroup.UserGroupID).Updates(modelData).Error
}

// Delete 删除用户组
func (r *userGroupRepository) Delete(userGroupID string) error {
	return r.db.Delete(&model.UserGroup{}, "user_group_id = ?", userGroupID).Error
}

// GetByID 根据ID获取用户组
func (r *userGroupRepository) GetByID(userGroupID string) (*dto.UserGroup, error) {
	var userGroup model.UserGroup
	if err := r.db.Where("user_group_id = ?", userGroupID).First(&userGroup).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&userGroup), nil
}

// GetByName 根据名称获取用户组
func (r *userGroupRepository) GetByName(name string) (*dto.UserGroup, error) {
	var userGroup model.UserGroup
	if err := r.db.Where("user_group_name = ?", name).First(&userGroup).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&userGroup), nil
}

// List 获取用户组列表
func (r *userGroupRepository) List(page, pageSize int) ([]*dto.UserGroup, int64, error) {
	var total int64
	var userGroups []model.UserGroup

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.UserGroup{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&userGroups).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.UserGroup, len(userGroups))
	for i, ug := range userGroups {
		dtoList[i] = r.convertToDTO(&ug)
	}

	return dtoList, total, nil
}

// ListByPriceFactor 根据价格系数筛选用户组
func (r *userGroupRepository) ListByPriceFactor(factor dto.UserGroupPriceFactor, page, pageSize int) ([]*dto.UserGroup, int64, error) {
	var total int64
	var userGroups []model.UserGroup

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.UserGroup{})

	// 构建查询条件
	factorJSON, err := common.FromStruct(factor)
	if err != nil {
		return nil, 0, fmt.Errorf("转换价格系数失败: %w", err)
	}

	query = query.Where("user_group_price_factor @> ?", factorJSON)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&userGroups).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.UserGroup, len(userGroups))
	for i, ug := range userGroups {
		dtoList[i] = r.convertToDTO(&ug)
	}

	return dtoList, total, nil
}

// ListByOptions 根据配置选项筛选用户组
func (r *userGroupRepository) ListByOptions(options dto.UserGroupOptions, page, pageSize int) ([]*dto.UserGroup, int64, error) {
	var total int64
	var userGroups []model.UserGroup

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.UserGroup{})

	// 构建查询条件
	optionsJSON, err := common.FromStruct(options)
	if err != nil {
		return nil, 0, fmt.Errorf("转换配置选项失败: %w", err)
	}

	query = query.Where("user_group_options @> ?", optionsJSON)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&userGroups).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.UserGroup, len(userGroups))
	for i, ug := range userGroups {
		dtoList[i] = r.convertToDTO(&ug)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *userGroupRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行用户组基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		// 创建测试用户组
		testGroup := &dto.UserGroup{
			UserGroupName:        fmt.Sprintf("benchmark_group_%d", i),
			UserGroupDescription: "基准测试用户组",
			UserGroupPriceFactor: dto.UserGroupPriceFactor{
				RequestPriceFactor:    float64(rand.Intn(50)+50) / 100,
				ResponsePriceFactor:   float64(rand.Intn(50)+50) / 100,
				CompletionPriceFactor: float64(rand.Intn(50)+50) / 100,
				CachePriceFactor:      float64(rand.Intn(50)+50) / 100,
			},
			UserGroupOptions: dto.UserGroupOptions{
				MaxConcurrentRequests: rand.Intn(10) + 1,
				DefaultLevel:          rand.Intn(3) + 1,
				ExtraAllowedModels:    []string{"gpt-3.5-turbo", "gpt-4"},
				ExtraAllowedChannels:  []string{"openai", "anthropic"},
				APIDiscount:           float64(rand.Intn(50)+50) / 100,
			},
		}

		// 创建
		if err := r.Create(testGroup); err != nil {
			utils.SysError("创建用户组失败: " + err.Error())
			return err
		}

		// 更新
		testGroup.UserGroupDescription = "已更新的基准测试用户组"
		if err := r.Update(testGroup); err != nil {
			utils.SysError("更新用户组失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(testGroup.UserGroupID); err != nil {
			utils.SysError("删除用户组失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
