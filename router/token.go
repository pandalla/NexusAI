package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTokenRouter(server *gin.Engine) {
	tokenRouter := server.Group("/token")
	tokenRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is token router!"})
	})
}
