package middlewares

import (
	"fmt"
	"server/pkg/e"
	"server/pkg/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Protect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int

		code = e.SUCCESS

		s := ctx.Request.Header.Get("Authorization")

		token := strings.TrimPrefix(s, "Bearer ")

		if s == "" || token == "" {
			code = e.INVALID_PARAMS
		}

		decoded, err := utils.ParseToken(token)

		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != e.SUCCESS {
			ctx.JSON(code, gin.H{
				"code":    code,
				"message": e.GetMessage(code),
				"data":    nil,
			})

			ctx.Abort()
		}

		fmt.Println(decoded.UserId)

		ctx.Set("userId", strconv.Itoa(decoded.UserId))

		ctx.Next()
	}
}
