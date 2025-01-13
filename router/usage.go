package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupUsageRouter(server *gin.Engine) {
	usageRouter := server.Group("/usage")
	usageRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is usage router!"})
	})
}
