package main

import (
	"fmt"
	"net/http"
	"nexus-ai/constant"
	"nexus-ai/middleware"
	"nexus-ai/model"
	"nexus-ai/mysql"
	"nexus-ai/redis"
	"nexus-ai/test"

	"nexus-ai/router"
	"nexus-ai/utils"

	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.SetupLog()

	// 初始化MySQL
	if err := mysql.Setup(); err != nil {
		utils.FatalLog("MySQL | " + err.Error())
	}
	utils.SysInfo("MySQL setup completed")
	defer func() {
		if err := mysql.Shutdown(); err != nil {
			utils.SysError("MySQL | " + err.Error())
		}
	}()

	// 初始化Gorm
	if err := model.InitGorm(); err != nil {
		utils.FatalLog("Gorm | " + err.Error())
	}
	utils.SysInfo("Gorm setup completed")

	// 执行MySQL基准测试
	test.TestUserRepository()
	test.TestUserGroupRepository()

	// 初始化Redis
	if err := redis.Setup(); err != nil {
		utils.FatalLog("Redis | " + err.Error())
	}
	utils.SysInfo("Redis setup completed")
	defer func() {
		if err := redis.Shutdown(); err != nil {
			utils.SysError("Redis | " + err.Error())
		}
	}()

	// 执行Redis基准测试
	if err := redis.RunBenchmarks(); err != nil {
		utils.SysError("Redis benchmarks failed: " + err.Error())
	} else {
		utils.SysInfo("Redis benchmarks completed successfully")
	}

	// 设置gin模式
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.New()
	server.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		utils.SysError(fmt.Sprintf("panic detected: %v", err)) // 出现panic时，输出错误日志
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"message": fmt.Sprintf("panic detected: %v. Please submit an issue to %s", err, constant.GitRepoURL), // 输出错误信息，并提示用户提交issue
				"type":    constant.ErrorTypeInternalServerError,                                                     // nexus-ai panic error
			},
		})
	}))
	utils.SysInfo("GIN Config setup completed")
	server.Use(middleware.RequestIDMiddleware()) // 添加requestID中间件 生成requestID
	middleware.SetupLog(server)                  // 为gin Engine添加日志服务 记录请求日志
	utils.SysInfo("Middleware setup completed")
	router.SetupRouter(server)
	utils.SysInfo("Router setup completed")
	backendPort, _ := strconv.Atoi(utils.GetEnv("BACKEND_PORT", constant.BackendPort))
	utils.SysInfo("Server starting on port " + strconv.Itoa(backendPort))
	err := server.Run(":" + strconv.Itoa(backendPort))
	if err != nil {
		utils.FatalLog("Failed to start HTTP server: " + err.Error())
	}
}
