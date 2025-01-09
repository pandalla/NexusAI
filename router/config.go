package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupConfigRouter(server *gin.Engine) {
	configRouter := server.Group("/config")
	configRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is config router!"})
	})
}
