package models

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	Action		string		`json:"action"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
}