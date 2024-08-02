package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func (m *Middleware) AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if len(tokenString) < 7 {
			ctx.Abort()
			ctx.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "Token Invalid",
			})
			return
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
			ctx.JSON(http.StatusForbidden, map[string]interface{}{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if _, ok := claims["user_id"]; !ok {
				ctx.JSON(http.StatusForbidden, map[string]interface{}{
					"message": "Token Invalid",
				})
				ctx.Abort()
				return
			}

			ctx.Set("user_id", claims["user_id"])
		} else {
			ctx.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "Token Invalid",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
