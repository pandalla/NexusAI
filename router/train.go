package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTrainRouter(server *gin.Engine) {
	trainRouter := server.Group("/train")
	trainRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is train router!"})
	})
}
