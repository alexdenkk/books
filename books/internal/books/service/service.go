package service

import (
	"alexdenkk/books/books/internal/books"
)

// Service - service layer struct
type Service struct {
	Repository books.Repository
}

// New - function for creating new service
func New(repo books.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
