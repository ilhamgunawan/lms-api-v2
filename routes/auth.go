package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/controllers"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	// Login
	auth.POST("/login", controllers.Login)

	// Validate token
	auth.POST("/validate", controllers.ValidateToken)
}
