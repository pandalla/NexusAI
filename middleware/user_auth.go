package middleware

import "github.com/gin-gonic/gin"

func UserAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
	}
}
