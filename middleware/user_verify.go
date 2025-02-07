package middleware

import (
	"net/http"
	"nexus-ai/constant"
	"nexus-ai/repository"
	"nexus-ai/utils"

	"github.com/gin-gonic/gin"
)

func UserVerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(string(constant.AccessTokenKey))
		refreshToken := c.GetHeader(string(constant.RefreshTokenKey))
		userID := c.GetHeader(string(constant.UserIDKey))
		user, err := repository.UserVerify(accessToken, refreshToken, userID)
		if err != nil {
			utils.AbortWhenUserVerifyFailed(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.Set(string(constant.UserIDKey), user.UserID)
		c.Set(string(constant.UserKey), user)
		// 继续处理请求
		c.Next()
	}
}
