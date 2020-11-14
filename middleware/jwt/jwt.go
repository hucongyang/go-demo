package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hucongyang/go-demo/pkg/errorCode"
	"github.com/hucongyang/go-demo/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = errorCode.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = errorCode.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = errorCode.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errorCode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != errorCode.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errorCode.GetMessage(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
