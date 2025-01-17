package middleware

import "github.com/gin-gonic/gin"

func DistributeChannelMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
	}
}
