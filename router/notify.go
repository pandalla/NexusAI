package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupNotifyRouter(server *gin.Engine) {
	notifyRouter := server.Group("/notify")
	notifyRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is notify router!"})
	})
}
