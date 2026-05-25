package main

import (
	"os"
	"log"
	
	"github.com/backend/config"
	//"github.com/backend/services"
	"github.com/backend/routes"
	"github.com/backend/middlewares"
	
	"github.com/gin-gonic/gin"
)

func main() {	
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.IARoute(router)

	router.Use(middlewares.CorsMiddleware())

	config.ConnectDatabase()

	go func() {
		log.Println("[Serveur] Démarrage immédiat sur le port " + os.Getenv("BACKEND_PORT"))
		if err := router.Run(":" + os.Getenv("BACKEND_PORT")); err != nil {
			log.Fatal("Erreur serveur :", err)
		}
	}()

	select {}
}