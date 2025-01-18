package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupChannelGroupRouter(server *gin.Engine) {
	channelGroupRouter := server.Group("/channel_group")
	channelGroupRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is channel_group router!"})
	})
}
