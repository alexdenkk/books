package repository

import (
	"gorm.io/gorm"
)

// Repository - books repository struct
type Repository struct {
	DB *gorm.DB
}

// New - function for creating a new books repository
func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
