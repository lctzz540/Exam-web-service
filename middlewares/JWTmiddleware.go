package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lctzz540/Exam-web-service/helpers"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := helpers.ValidateToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		email := claims.Email

		c.Set("contextEmail", email)
		c.Next()
	}
}
