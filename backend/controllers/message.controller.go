package controllers

import (
	"github.com/backend/config"
	"github.com/backend/models"
	"github.com/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/ollama/ollama/api"
)

func SendMessageController(c *gin.Context) {
	var message models.History
	if err := c.BindJSON(&message); err != nil {
		c.JSON(400, gin.H{"error": "Body invalide"})
		return
	}

	message.Role = "user"
	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(500, gin.H{"error": "Impossible de créer le message"})
		return
	}

	var history []models.History
	config.DB.Where("conversation_id = ?", message.ConversationID).Find(&history)

	var ollamaHistory []api.Message
	for _, h := range history {
		ollamaHistory = append(ollamaHistory, api.Message{
			Role:    h.Role,
			Content: h.Content,
		})
	}

	response, err := services.AskIA(message.Content, ollamaHistory)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erreur IA"})
		return
	}

	iaMessage := models.History{
		ConversationID: message.ConversationID,
		Role:           "assistant",
		Content:        response,
	}
	config.DB.Create(&iaMessage)

	c.JSON(200, gin.H{"response": response})
}