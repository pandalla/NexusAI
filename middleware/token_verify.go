package middleware

import (
	"net/http"
	"nexus-ai/constant"
	"nexus-ai/repository"
	"nexus-ai/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenVerifyMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var tokenParts []string
		tokenKey := c.Request.Header.Get("Authorization")  // 获取Authorization头
		tokenKey = strings.TrimPrefix(tokenKey, "Bearer ") // 去掉Bearer
		if tokenKey == "midjourney-proxy" {
			tokenKey = c.Request.Header.Get("mj-api-secret")
			tokenKey = strings.TrimPrefix(tokenKey, "Bearer ")
			tokenKey = strings.TrimPrefix(tokenKey, "sk-")
			tokenParts = strings.Split(tokenKey, "-")
			tokenKey = tokenParts[0]
		} else {
			tokenKey = strings.TrimPrefix(tokenKey, "sk-")
			tokenParts = strings.Split(tokenKey, "-")
			tokenKey = tokenParts[0]
		}
		token, err := repository.TokenVerify(tokenKey)
		if err != nil {
			utils.AbortWhenTokenVerifyFailed(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.Set(string(constant.TokenKey), token)
		c.Next()
	}
}
