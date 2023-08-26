package model

import (
	"gorm.io/gorm"
)

// Genre - genre model struct
type Genre struct {
	gorm.Model

	Name string `json:"name" gorm:"not null"`
}
