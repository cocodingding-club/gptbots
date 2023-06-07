package model

import (
	"gorm.io/gorm"
)

type Chatbot struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:255"`
}
