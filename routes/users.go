package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/controllers"
)

func UsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	// Get all users
	users.GET("", controllers.GetUsers)

	// Get user by id
	users.GET("/:id", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Get user by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Create user
	users.POST("/create", func(c *gin.Context) {
		data := make(map[string]string)
		data["message"] = "Create user"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Update user by id
	users.PUT("/:id/update", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Update user by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Delete user by id
	users.DELETE("/:id/delete", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Delete user by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}
