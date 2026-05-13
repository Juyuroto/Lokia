package main

import (
	"github.com/backend/config"
	"github.com/backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	services.LocalIA()

	router.Run(":5000")
}