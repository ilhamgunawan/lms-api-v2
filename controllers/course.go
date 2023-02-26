package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/db"
	"github.com/ilhamgunawan/lms-api-v2/models"
)

type CourseController struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func CreateCourse(c *gin.Context) {
	userId := c.GetString("user_id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	body := CourseController{}

	err := c.ShouldBindJSON(&body)

	if err != nil || body.Title == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing fields"})
		return
	}

	course, err := models.CreateCourse(db.Course{
		Title:       body.Title,
		Description: body.Description,
		AuthorId:    userId,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func UpdateCourse(c *gin.Context) {
	courseId := c.Param("id")
	body := CourseController{}

	err := c.ShouldBindJSON(&body)

	if err != nil || courseId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing fields"})
		return
	}

	updateCourse := db.Course{
		ID:          courseId,
		Title:       body.Title,
		Description: body.Description,
		Status:      body.Status,
	}

	course, err := models.UpdateCourse(updateCourse)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}
