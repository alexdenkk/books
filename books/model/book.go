package model

import (
	"gorm.io/gorm"
)

// Book - book model struct
type Book struct {
	gorm.Model

	Name    string `json:"name" gorm:"not null"`
	Author  string `json:"author" gorm:"not null"`
	Year    int    `json:"year" gorm:"not null"`
	GenreID uint   `json:"genre_id" gorm:"not null"`
}
