package middleware

import (
	"go-web-starter/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.secret")), nil
		})
		if err != nil || !token.Valid {
			response.Fail(c, "unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
