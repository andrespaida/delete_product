package main

import (
	"delete_product/database"
	"delete_product/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env")
	}

	database.ConnectDB()

	r := gin.Default()
	r.DELETE("/products/:id", handlers.DeleteProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4003"
	}

	r.Run(":" + port)
}