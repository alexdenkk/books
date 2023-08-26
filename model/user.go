package model

import (
	"gorm.io/gorm"
)

// User - user model struct
type User struct {
	gorm.Model

	Login    string `json:"login" gorm:"not null"`
	Password string `json:"password" gorm:"not null"` // hash
	Role     uint8  `json:"role" gorm:"not null"`
}
