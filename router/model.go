package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupModelRouter(server *gin.Engine) {
	modelRouter := server.Group("/model")
	modelRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is model router!"})
	})
}
