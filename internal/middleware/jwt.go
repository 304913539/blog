package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		token = c.GetHeader("Authorization")

		if token == "" {
			ecode = errcode.UnauthorizedTokenError
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					ecode = errcode.UnauthorizedTokenError
				} else {
					ecode = errcode.UnauthorizedTokenTimeout
				}

			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()

	}
}
