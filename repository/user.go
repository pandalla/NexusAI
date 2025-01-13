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

// UserRepository 用户仓储接口
type UserRepository interface {
	Create(user *dto.User) error
	Update(user *dto.User) error
	Delete(userID string) error
	GetByID(userID string) (*dto.User, error)
	GetByEmail(email string) (*dto.User, error)
	GetByPhone(phone string) (*dto.User, error)
	GetByUsername(username string) (*dto.User, error)
	List(page, pageSize int) ([]*dto.User, int64, error)
	ListByQuota(quota dto.UserQuota, page, pageSize int) ([]*dto.User, int64, error)
	ListByOptions(options dto.UserOptions, page, pageSize int) ([]*dto.User, int64, error)
	Benchmark(count int) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *userRepository) convertToDTO(model *model.User) *dto.User {
	if model == nil {
		return nil
	}

	var oauthInfo dto.OAuthInfo
	var userQuota dto.UserQuota
	var userOptions dto.UserOptions

	if err := model.OAuthInfo.ToStruct(&oauthInfo); err != nil {
		utils.SysError("解析OAuth信息失败:" + err.Error())
	}

	if err := model.UserQuota.ToStruct(&userQuota); err != nil {
		utils.SysError("解析配额信息失败:" + err.Error())
	}

	if err := model.UserOptions.ToStruct(&userOptions); err != nil {
		utils.SysError("解析用户选项失败:" + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.User{
		UserID:        model.UserID,
		UserGroupID:   model.UserGroupID,
		Username:      model.Username,
		Password:      model.Password,
		Email:         model.Email,
		Phone:         model.Phone,
		OAuthInfo:     oauthInfo,
		UserQuota:     userQuota,
		UserOptions:   userOptions,
		LastLoginTime: model.LastLoginTime,
		LastLoginIP:   model.LastLoginIP,
		Status:        model.Status,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *userRepository) convertToModel(dto *dto.User) (*model.User, error) {
	if dto == nil {
		return nil, nil
	}

	oauthInfoJSON, err := common.FromStruct(dto.OAuthInfo)
	if err != nil {
		return nil, fmt.Errorf("转换OAuth信息失败: %w", err)
	}

	userQuotaJSON, err := common.FromStruct(dto.UserQuota)
	if err != nil {
		return nil, fmt.Errorf("转换配额信息失败: %w", err)
	}

	userOptionsJSON, err := common.FromStruct(dto.UserOptions)
	if err != nil {
		return nil, fmt.Errorf("转换用户选项失败: %w", err)
	}

	return &model.User{
		UserID:        dto.UserID,
		UserGroupID:   dto.UserGroupID,
		Username:      dto.Username,
		Password:      dto.Password,
		Email:         dto.Email,
		Phone:         dto.Phone,
		OAuthInfo:     oauthInfoJSON,
		UserQuota:     userQuotaJSON,
		UserOptions:   userOptionsJSON,
		LastLoginTime: dto.LastLoginTime,
		LastLoginIP:   dto.LastLoginIP,
		Status:        dto.Status,
	}, nil
}

// Create 创建用户
func (r *userRepository) Create(user *dto.User) error {
	model, err := r.convertToModel(user)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新用户
func (r *userRepository) Update(user *dto.User) error {
	modelData, err := r.convertToModel(user)
	if err != nil {
		return err
	}
	return r.db.Model(&model.User{}).Where("user_id = ?", user.UserID).Updates(modelData).Error
}

// Delete 删除用户
func (r *userRepository) Delete(userID string) error {
	return r.db.Delete(&model.User{}, "user_id = ?", userID).Error
}

// GetByID 根据ID获取用户
func (r *userRepository) GetByID(userID string) (*dto.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&user), nil
}

// GetByEmail 根据邮箱获取用户
func (r *userRepository) GetByEmail(email string) (*dto.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&user), nil
}

// GetByPhone 根据手机号获取用户
func (r *userRepository) GetByPhone(phone string) (*dto.User, error) {
	var user model.User
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&user), nil
}

// GetByUsername 根据用户名获取用户
func (r *userRepository) GetByUsername(username string) (*dto.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&user), nil
}

// List 获取用户列表
func (r *userRepository) List(page, pageSize int) ([]*dto.User, int64, error) {
	var total int64
	var users []model.User

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.User, len(users))
	for i, u := range users {
		dtoList[i] = r.convertToDTO(&u)
	}

	return dtoList, total, nil
}

// ListByQuota 根据配额筛选用户
func (r *userRepository) ListByQuota(quota dto.UserQuota, page, pageSize int) ([]*dto.User, int64, error) {
	var total int64
	var users []model.User

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.User{})

	quotaJSON, err := common.FromStruct(quota)
	if err != nil {
		return nil, 0, fmt.Errorf("转换配额信息失败: %w", err)
	}

	query = query.Where("user_quota @> ?", quotaJSON)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.User, len(users))
	for i, u := range users {
		dtoList[i] = r.convertToDTO(&u)
	}

	return dtoList, total, nil
}

// ListByOptions 根据用户选项筛选用户
func (r *userRepository) ListByOptions(options dto.UserOptions, page, pageSize int) ([]*dto.User, int64, error) {
	var total int64
	var users []model.User

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.User{})

	optionsJSON, err := common.FromStruct(options)
	if err != nil {
		return nil, 0, fmt.Errorf("转换用户选项失败: %w", err)
	}

	query = query.Where("user_options @> ?", optionsJSON)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.User, len(users))
	for i, u := range users {
		dtoList[i] = r.convertToDTO(&u)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *userRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行用户基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testUser := &dto.User{
			Username: fmt.Sprintf("benchmark_user_%d", i),
			Email:    fmt.Sprintf("benchmark_%d@example.com", i),
			Phone:    fmt.Sprintf("1%010d", rand.Intn(10000000000)),
			Password: utils.HashPassword(fmt.Sprintf("test%d", i)),
			UserQuota: dto.UserQuota{
				TotalQuota:  float64(rand.Intn(1000)),
				FrozenQuota: float64(rand.Intn(100)),
				GiftQuota:   float64(rand.Intn(100)),
			},
			UserOptions: dto.UserOptions{
				MaxConcurrentRequests: rand.Intn(10) + 1,
				DefaultLevel:          rand.Intn(3) + 1,
				APIDiscount:           float64(rand.Intn(50)+50) / 100,
			},
			Status: 1,
		}

		// 创建
		if err := r.Create(testUser); err != nil {
			utils.SysError("创建用户失败: " + err.Error())
			return err
		}

		// 更新
		testUser.UserOptions.DefaultLevel = rand.Intn(5) + 1
		if err := r.Update(testUser); err != nil {
			utils.SysError("更新用户失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(testUser.UserID); err != nil {
			utils.SysError("删除用户失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
