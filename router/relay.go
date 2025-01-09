package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRelayRouter(server *gin.Engine) {
	relayRouter := server.Group("/relay")
	relayRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is relay router!"})
	})
}
