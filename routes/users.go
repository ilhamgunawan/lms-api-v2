package routes

import (
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
	users.PUT("/:id/update", controllers.UpdateUserById)

	// Delete user by id
	users.DELETE("/:id/delete", controllers.DeleteUserById)
}
