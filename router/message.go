package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupMessageRouter(server *gin.Engine) {
	messageRouter := server.Group("/message")
	messageRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is message router!"})
	})
}
