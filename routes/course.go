package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/middlewares"
)

func CourseRoutes(rg *gin.RouterGroup) {
	course := rg.Group("/courses")

	// Use validate token middleware for router group
	course.Use(middlewares.ValidateToken())

	// Get all courses
	course.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Get all courses"})
	})

	// Get course by id
	course.GET("/:id", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Get course by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Create course
	course.POST("/create", func(c *gin.Context) {
		data := make(map[string]string)
		data["message"] = "Create course"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Update course by id
	course.PUT("/:id/update", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Update course by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Delete course by id
	course.DELETE("/:id/delete", func(c *gin.Context) {
		data := make(map[string]string)
		data["id"] = c.Param("id")
		data["message"] = "Delete course by id"

		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}
