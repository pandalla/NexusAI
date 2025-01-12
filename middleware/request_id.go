package middleware

import (
	"context"
	"nexus-ai/constant"
	"nexus-ai/utils"

	"github.com/gin-gonic/gin"
)

func RequestIDMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		requestID := utils.GetTimeString() + "-" + utils.GenerateRandomString(12)
		c.Set(string(constant.RequestIDKey), requestID)
		ctx := context.WithValue(c.Request.Context(), constant.RequestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)
		c.Header(string(constant.RequestIDKey), requestID)
		c.Next()
	}
}
