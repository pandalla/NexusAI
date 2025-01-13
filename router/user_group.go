package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupUserGroupRouter(server *gin.Engine) {
	userGroupRouter := server.Group("/user_group")
	userGroupRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is user_group router!"})
	})
}
