package service

import (
	"alexdenkk/books/service/genres"
)

// Service - service layer struct
type Service struct {
	Repository genres.Repository
}

// New - function for creating new service
func New(repo genres.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
