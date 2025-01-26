package middleware

import (
	"net/http"
	"nexus-ai/constant"
	dto "nexus-ai/dto/model"
	"nexus-ai/repository"
	"nexus-ai/utils"

	"github.com/gin-gonic/gin"
)

func QuotaVerifyMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.MustGet(string(constant.TokenKey)).(*dto.Token)
		tokenQuotaLeft := token.TokenQuotaLeft
		userID := token.UserID
		ok, err := repository.QuotaVerify(userID, tokenQuotaLeft)
		if !ok {
			utils.AbortWhenQuotaVerifyFailed(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.Next()
	}
}
