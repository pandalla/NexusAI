package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupQuotaRouter(server *gin.Engine) {
	quotaRouter := server.Group("/quota")
	quotaRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is quota router!"})
	})
}
