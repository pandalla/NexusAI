package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"nexus-ai/dto/model"
	"nexus-ai/middleware"
	"nexus-ai/repository"
	"nexus-ai/utils"
)

// UserController 用户控制器
type UserController struct {
	userRepo repository.UserRepository
}

// NewUserController 创建用户控制器实例
func NewUserController(userRepo repository.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

// RegisterRequest 用户注册请求结构体
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"` // 邮箱，必填，格式需正确
	Password string `json:"password" binding:"required"`    // 密码，必填
}

// LoginRequest 用户登录请求结构体
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"` // 邮箱，必填
	Password string `json:"password" binding:"required"`    // 密码，必填
}

// Register 注册用户
func (uc *UserController) Register(c *gin.Context) {
	var request RegisterRequest

	// 将请求体中的数据绑定到 RegisterRequest 结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 创建 User 实例并填充数据
	user := model.User{
		Email:    request.Email,
		Status:   1, // 假设状态为 1 表示可用
		Phone:    uuid.New().String(),
		Password: utils.HashPassword(request.Password), // 设定哈希密码
	}

	// 调用 UserRepository 创建用户
	if err := uc.userRepo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败: " + err.Error()})
		return
	}

	// 注册成功，返回 201 状态码和创建的用户信息
	c.JSON(http.StatusCreated, gin.H{"message": "注册成功", "email": user.Email})
}

// Login 用户登录
func (uc *UserController) Login(c *gin.Context) {
	var request LoginRequest

	// 将请求体中的数据绑定到 LoginRequest 结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 查询用户
	user, err := uc.userRepo.GetByEmail(request.Email) // 使用邮箱查询用户
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在或密码错误"})
		return
	}

	// 验证密码
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在或密码错误"})
		return
	}

	token, err := middleware.GenerateToken(utils.HashPassword(user.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "生成令牌失败")
		return
	}
	// 返回登录成功的信息和用户的邮箱
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": token})
}
