package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupMonitorRouter(server *gin.Engine) {
	monitorRouter := server.Group("/monitor")
	monitorRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is monitor router!"})
	})
}
