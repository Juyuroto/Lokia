package routes

import (
	"github.com/gin-gonic/gin"
	
	"github.com/backend/controllers"
)

func IARoute(router *gin.Engine) {
	
    router.POST("/chat", controllers.SendMessageController)
    router.GET("/conversations", controllers.FetchConversationsController)
    router.GET("/conversations/:id", controllers.FetchConversationByIDController)
    router.POST("/conversations", controllers.CreateConversationController)
    router.DELETE("/conversations/:id", controllers.DeleteConversationController)
}