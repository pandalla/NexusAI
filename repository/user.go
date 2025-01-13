package repository

import (
	"context"
	"errors"
	"fmt"
	"nexus-ai/model"
	"nexus-ai/utils"

	"gorm.io/gorm"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	// 创建用户
	Create(ctx context.Context, user *model.User) error
	// 根据用户ID获取用户
	GetByID(ctx context.Context, userID string) (*model.User, error)
	// 根据邮箱获取用户
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	// 根据手机号获取用户
	GetByPhone(ctx context.Context, phone string) (*model.User, error)
	// 根据用户名获取用户
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	// 更新用户信息
	Update(ctx context.Context, user *model.User) error
	// 删除用户（软删除）
	Delete(ctx context.Context, userID string) error
	// 批量获取用户
	List(ctx context.Context, offset, limit int) ([]*model.User, int64, error)
	// 创建测试用户
	CreateTestUser(ctx context.Context) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create 创建用户
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GetByID 根据用户ID获取用户
func (r *userRepository) GetByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByPhone 根据手机号获取用户
func (r *userRepository) GetByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete 删除用户（软删除）
func (r *userRepository) Delete(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&model.User{}).Error
}

// List 批量获取用户
func (r *userRepository) List(ctx context.Context, offset, limit int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	err := r.db.WithContext(ctx).Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateTestUser 创建测试用户
func (r *userRepository) CreateTestUser(ctx context.Context) error {
	// 检查是否已存在测试用户
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("username LIKE ?", "test_user_%").Count(&count).Error; err != nil {
		return fmt.Errorf("检查测试用户失败: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("测试用户已存在")
	}

	// 生成随机用户名
	randomUsername := fmt.Sprintf("test_user_%s", utils.GenerateRandomString(8))
	// 生成随机邮箱
	randomEmail := fmt.Sprintf("test_%s@example.com", utils.GenerateRandomString(8))
	// 生成随机手机号
	randomPhone := fmt.Sprintf("1%s", utils.GenerateRandomNumber(10))

	testUser := &model.User{
		UserID:   utils.GenerateRandomUUID(12),
		Username: randomUsername,
		Email:    randomEmail,
		Phone:    randomPhone,
		Password: utils.HashPassword("test123456"), // 默认密码
	}

	return r.Create(ctx, testUser)
}
