package middleware

import (
	"bookstore/store/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthMiddleware struct {
	auth *jwt.JWT
}

func New(auth *jwt.JWT) *AuthMiddleware {
	return &AuthMiddleware{auth: auth}
}

func (m *AuthMiddleware) Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		bearerToken := context.GetHeader("Authorization")
		if bearerToken == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		var tokenString string
		if len(strings.Split(bearerToken, " ")) == 2 {
			tokenString = strings.Split(bearerToken, " ")[1]
		} else {
			context.JSON(401, gin.H{"error": "invalid bearer token"})
			context.Abort()
			return
		}
		email, err := m.auth.Validate(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set("email", email)
		context.Next()
	}
}
