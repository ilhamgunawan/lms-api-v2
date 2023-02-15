package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/db"
	"github.com/ilhamgunawan/lms-api-v2/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Check if env file is exist
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load env file")
	}

	db.Init() // Initialize db connection

	r := gin.Default()

	routes.MakeRoutes(r)

	port := os.Getenv("PORT")

	r.Run(":" + port)
}
