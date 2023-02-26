package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/db"
	"github.com/ilhamgunawan/lms-api-v2/routes"
	"github.com/joho/godotenv"

	// cors "github.com/rs/cors/wrapper/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// Check if env file is exist
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load env file")
	}

	db.Init() // Initialize db connection

	r := gin.Default()

	// Handle CORS
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:       []string{"http://localhost:7001"},
	// 	AllowCredentials:     true,
	// 	AllowedMethods:       []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
	// 	AllowedHeaders:       []string{"*", "Authorization"},
	// 	OptionsPassthrough:   true,
	// 	OptionsSuccessStatus: 204,
	// })
	// r.Use(c)
	c := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:7001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	r.Use(c)

	routes.MakeRoutes(r)

	port := os.Getenv("PORT")

	r.Run(":" + port)
}
