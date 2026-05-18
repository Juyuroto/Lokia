package controllers

import (
	"github.com/backend/config"
	"github.com/backend/models"

	"github.com/gin-gonic/gin"
)

func FetchConversationsController(c *gin.Context) {
	conversation := []models.Conversation{}
	if err := config.DB.Find(&conversation).Error; err != nil  {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, &conversation)
}

func FetchConversationByIDController(c *gin.Context) {
	id := c.Param("id")
    conversation := []models.Conversation{}
    if err := config.DB.Where("id = ?", id).Find(&conversation).Error; err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, &conversation)
}

func CreateConversationController(c *gin.Context) {
	var conversation models.Conversation

	 c.BindJSON(&conversation)
		
	if err := config.DB.Create(&conversation).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, &conversation)
}

func DeleteConversationController(c *gin.Context) {
	var conversation models.Conversation
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&conversation).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, &conversation)
}