package utils

import (
	"fmt"
	"nexus-ai/constant"

	"github.com/gin-gonic/gin"
)

func AbortWhenTokenVerifyFailed(c *gin.Context, statusCode int, message string) {
	requestID := c.GetString(string(constant.RequestIDKey))
	c.AbortWithStatusJSON(statusCode, gin.H{
		"success": false,
		"error": gin.H{
			"message": message,
			"type":    constant.ErrorTypeTokenVerifyFailed,
		},
	})
	LogError(c.Request.Context(), fmt.Sprintf("requestID: %s, type: %s, message: %s", requestID, constant.ErrorTypeTokenVerifyFailed, message))
}

func AbortWhenQuotaVerifyFailed(c *gin.Context, statusCode int, message string) {
	requestID := c.GetString(string(constant.RequestIDKey))
	c.AbortWithStatusJSON(statusCode, gin.H{
		"success": false,
		"error": gin.H{
			"message": message,
			"type":    constant.ErrorTypeQuotaVerifyFailed,
		},
	})
	LogError(c.Request.Context(), fmt.Sprintf("requestID: %s, type: %s, message: %s", requestID, constant.ErrorTypeQuotaVerifyFailed, message))
}

func AbortWhenChannelDistributeFailed(c *gin.Context, statusCode int, message string) {
	requestID := c.GetString(string(constant.RequestIDKey))
	c.AbortWithStatusJSON(statusCode, gin.H{
		"success": false,
		"error": gin.H{
			"message": message,
			"type":    constant.ErrorTypeChannelDistributeFailed,
		},
	})
	LogError(c.Request.Context(), fmt.Sprintf("requestID: %s, type: %s, message: %s", requestID, constant.ErrorTypeChannelDistributeFailed, message))
}
