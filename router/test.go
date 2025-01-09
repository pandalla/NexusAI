package router

import (
	"net/http"
	"nexus-ai/utils"

	"github.com/gin-gonic/gin"
)

func SetupTestRouter(router *gin.Engine) {
	testRouter := router.Group("/test")
	testRouter.GET("/info", func(c *gin.Context) {
		ctx := c.Request.Context() // 获取请求上下文

		utils.LogInfo(ctx, "This is an info log.")
		utils.LogWarn(ctx, "This is a warning log.")
		utils.LogError(ctx, "This is an error log.")
		utils.SysInfo("This is a system info log.")
		utils.SysError("This is a system error log.")
		utils.SysWarn("This is a system warning log.")
		// utils.FatalLog("This is a fatal log.")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is test router!"})
	})
}
