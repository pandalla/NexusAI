package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is health check!"})
	})
	SetupBillingRouter(router)
	SetupChannelRouter(router)
	SetupConfigRouter(router)
	SetupLogRouter(router)
	SetupMessageRouter(router)
	SetupModelRouter(router)
	SetupMonitorRouter(router)
	SetupNotifyRouter(router)
	SetupPayRouter(router)
	SetupRelayRouter(router)
	SetupTestRouter(router)
	SetupTokenRouter(router)
	SetupTrainRouter(router)
	SetupUserRouter(router)
}
