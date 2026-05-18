package main

import (
	"os"
	
	"github.com/backend/config"
	"github.com/backend/services"
	"github.com/backend/routes"
	"github.com/backend/middlewares"
	
	"github.com/gin-gonic/gin"
)

func main() {
	host := os.Getenv("BACKEND_PORT")
	
	router := gin.Default()

	config.ConnectDatabase()

	router.Use(middlewares.CorsMiddleware())

	services.LoadIA()

	routes.IARoute(router)

	router.Run(host)
}