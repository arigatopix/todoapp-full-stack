package middlewares

import (
	"net/http"
	"server/pkg/e"
	"server/pkg/utils"
	"server/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Protect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int

		code = e.SUCCESS

		var token string

		cookie, _ := ctx.Cookie("token")
		s := ctx.Request.Header.Get("Authorization")

		if s != "" {
			token = strings.TrimPrefix(s, "Bearer ")
		} else if cookie != "" {
			token = cookie
		}

		if token == "" {
			code = e.ERROR_UNAUTHORIZED
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

		authService := services.User{
			ID: decoded.UserID,
		}

		existed, err := authService.UserExisted()

		if !existed || err != nil {
			code = e.ERROR_USER_NOT_EXIST
		}

		if code != e.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": e.GetMessage(code),
				"data":    nil,
			})

			ctx.Abort()
			return
		}

		ctx.Set("userId", strconv.Itoa(decoded.UserID))

		ctx.Next()
	}
}
