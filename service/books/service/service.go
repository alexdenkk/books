package service

import (
	"alexdenkk/books/service/books"
)

// Service - books service struct
type Service struct {
	Repository books.Repository
}

// New - function for creating a new books service
func New(repo books.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
