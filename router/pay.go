package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupPayRouter(server *gin.Engine) {
	payRouter := server.Group("/pay")
	payRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is pay router!"})
	})
}
