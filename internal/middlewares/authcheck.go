package middlewares

import (
	"fmt"
	v1 "gin-layout/api/app/v1"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (m *Middleware) AuthCheck(ctx *gin.Context, operation string) error {
	if operation == v1.OperationAppGreeter {
		// 不需要验证
		return nil
	}
	tokenString := ctx.GetHeader("Authorization")
	if len(tokenString) < 7 {
		return v1.ErrorBadRequest("Token Invalid")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(m.config.Other.JwtKey), nil
	})
	if err != nil {
		return v1.ErrorForbidden(err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if _, ok := claims["user_id"]; !ok {
			return v1.ErrorForbidden("user_id not found")
		}

		ctx.Set("user_id", claims["user_id"])
	} else {
		return v1.ErrorForbidden("Token Invalid")
	}
	return nil
}
