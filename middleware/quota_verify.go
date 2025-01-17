package middleware

import "github.com/gin-gonic/gin"

func QuotaVerifyMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
	}
}
