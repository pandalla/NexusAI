package middleware

import "github.com/gin-gonic/gin"

func TokenVerifyMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
	}
}
