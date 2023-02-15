package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/routes"
)

func main() {
	r := gin.Default()

	routes.MakeRoutes(r)

	r.Run(":7002")
}
