package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"todolist/internal/types"
)

func ValidateLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path
		if path == "/user/login" {
			return
		}

		// 找到header中的token
		header := c.GetHeader("Authorization")
		split := strings.Split(header, " ")

		if len(split) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.Result{
				Code: 4,
				Msg:  "invalid session",
				Data: nil,
			})
		}

		tokenString := split[1]

		token, err := jwt.ParseWithClaims(tokenString, &types.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(types.UserSignKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.Result{
				Code: 4,
				Msg:  "invalid token",
				Data: nil,
			})
		} else if claims, ok := token.Claims.(*types.UserClaim); ok {
			c.Set("userId", claims.Id)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.Result{
				Code: 4,
				Msg:  "invalid token",
				Data: nil,
			})
		}
	}
}
