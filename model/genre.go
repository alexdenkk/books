package model

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model

	Name string `json:"name" gorm:"not null"`
}
