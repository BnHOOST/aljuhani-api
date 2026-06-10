package main

import (
	"aljuhani-api/configs"
	"aljuhani-api/middlewares"
	"aljuhani-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	configs.ConnectDatabase()

	router := gin.Default()

	router.Use(
		middlewares.LocalizationMiddleware(),
	)

	router.Static("/public/imgs", "./public/imgs")

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}