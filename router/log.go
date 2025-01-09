package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupLogRouter(server *gin.Engine) {
	logRouter := server.Group("/log")
	logRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is log router!"})
	})
}
