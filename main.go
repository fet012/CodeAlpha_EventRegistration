package main

import (
	"EventRegistration/config"
	"EventRegistration/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	db := config.DbConnect()
	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":" + os.Getenv("PORT"))
}