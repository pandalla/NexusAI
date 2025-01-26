package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserVerifyMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 继续处理请求
		c.Next()
	}
}
