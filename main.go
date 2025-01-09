package main

import (
	"fmt"
	"net/http"
	"nexus-ai/constant"
	"nexus-ai/middleware"
	"nexus-ai/router"
	"nexus-ai/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.SetupLog()
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
	server.Use(middleware.RequestIDMiddleware()) // 添加requestID中间件 生成requestID
	middleware.SetupLog(server)                  // 为gin Engine添加日志服务 记录请求日志
	utils.SysInfo("Middleware setup completed")
	router.SetupRouter(server)
	utils.SysInfo("Router setup completed")
	var port = os.Getenv("PORT")
	if port == "" { // 如果PORT环境变量未设置，则使用配置文件中的backend_port
		port = constant.BackendPort
	}
	utils.SysInfo("Starting server on port " + port)
	err := server.Run(":" + port)
	if err != nil { // 如果启动失败，则输出错误日志
		utils.FatalLog("Failed to start HTTP server: " + err.Error())
	}
}
