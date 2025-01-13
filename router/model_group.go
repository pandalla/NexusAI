package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupModelGroupRouter(server *gin.Engine) {
	modelGroupRouter := server.Group("/model_group")
	modelGroupRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is model_group router!"})
	})
}
