package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupBillingRouter(server *gin.Engine) {
	billingRouter := server.Group("/billing")
	billingRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is billing router!"})
	})
}
