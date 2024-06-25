package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"micro-mini/shared/database"
	"micro-mini/user/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	router := gin.Default()
	routes.UserRoutes(router)

	router.Run(":" + port)
}
