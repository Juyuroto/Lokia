package routes

import (
	"github.com/gin-gonic/gin"
	
	"github.com/backend/controllers"
)

func IARoute(router *gin.Engine) {
	router.POST("/chat", controllers.SendMessageController)
	router.GET("/history", controllers.FetchHistoryController)
	router.GET("/history/:id", controllers.FetchHistoryPerIDController)
	router.DELETE("/history/:id", controllers.DeleteHistoryController)
}