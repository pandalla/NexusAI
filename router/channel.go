package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupChannelRouter(server *gin.Engine) {
	channelRouter := server.Group("/channel")
	channelRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is channel router!"})
	})
}
