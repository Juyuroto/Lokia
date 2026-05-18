package main

import (
	"github.com/backend/config"
	"github.com/backend/services"
	"github.com/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	routes.IARoute(router)

	services.LoadIA()

	router.Run(":5000")
}