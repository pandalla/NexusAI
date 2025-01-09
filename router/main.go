package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
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
