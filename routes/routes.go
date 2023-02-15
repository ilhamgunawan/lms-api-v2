package routes

import (
	"github.com/gin-gonic/gin"
)

func MakeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	AuthRoutes(v1)
	UsersRoutes(v1)
	CourseRoutes(v1)
}
