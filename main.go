package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/Rajesh-0907/go-gin-boilerplate/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Fatal(r.Run(":" + port))
}
