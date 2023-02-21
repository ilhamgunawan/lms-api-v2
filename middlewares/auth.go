package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/models"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		err := models.VerifyToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		c.Next()
	}
}
