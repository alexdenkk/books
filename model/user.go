package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Login    string `json:"login"`
	Password string `json:"password"` // hash
	Role     uint8  `json:"role"`
}
