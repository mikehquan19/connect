package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mikehquan19/connect/routes"
	"github.com/mikehquan19/connect/setup"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	setup.ConnectDB(os.Getenv("MONGO_URI"), "workshop")

	r := gin.Default()

	routes.RegisterUserRoutes(r)

	r.Run(":8080")
}
