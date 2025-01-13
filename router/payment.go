package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRouter(server *gin.Engine) {
	paymentRouter := server.Group("/payment")
	paymentRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is payment router!"})
	})
}
