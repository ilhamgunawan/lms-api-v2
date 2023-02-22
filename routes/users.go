package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/controllers"
	"github.com/ilhamgunawan/lms-api-v2/middlewares"
)

func UsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	// Use validate token middleware for router group
	users.Use(middlewares.ValidateToken())

	// Get all users
	users.GET("", controllers.GetUsers)

	// Get user by id
	users.GET("/:id", controllers.GetUserById)

	// Create user
	users.POST("/create", controllers.CreateUser)

	// Update user by id
	users.PUT("/:id/update", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Update user by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Delete user by id
	users.DELETE("/:id/delete", controllers.DeleteUserById)
}
