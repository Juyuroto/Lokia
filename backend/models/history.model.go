package models

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	ConversationID		uint		`json:"conversation_id" gorm:"index"`
	Role				string		`json:"role"`
	Content				string		`json:"content"`
}