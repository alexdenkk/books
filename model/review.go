package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model

	UserID uint   `json:"user_id" gorm:"not null"`
	BookID uint   `json:"book_id" gorm:"not null"`
	Text   string `json:"text" gorm:"not null"`
	Rate   uint8  `json:"rate" gorm:"not null"`
}
