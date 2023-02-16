package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/controllers"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	// Login
	auth.POST("/login", controllers.Login)

	// Validate token
	auth.POST("/validate", func(c *gin.Context) {
		data := make(map[string]string)
		data["message"] = "Validate token"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}
