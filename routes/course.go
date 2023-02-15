package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CourseRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/courses")

	// Get all courses
	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Get all courses"})
	})

	// Get course by id
	users.GET("/:id", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Get course by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Create course
	users.POST("/create", func(c *gin.Context) {
		data := make(map[string]string)
		data["message"] = "Create course"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Update course by id
	users.PUT("/:id/update", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Update course by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Delete course by id
	users.DELETE("/:id/delete", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Delete course by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}
