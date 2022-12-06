package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Name   string `json:"name" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Year   uint   `json:"year" gorm:"not null"`
	Genre  string `json:"genre" gorm:"not null"`
}
