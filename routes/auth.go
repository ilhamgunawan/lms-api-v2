package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	// Login
	auth.POST("/login", func(c *gin.Context) {
		data := make(map[string]string)
		data["message"] = "Log in"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Validate token
	auth.POST("/validate", func(c *gin.Context) {
		data := make(map[string]string)
		data["message"] = "Validate token"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}
