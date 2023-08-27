package service

import (
	"alexdenkk/books/service/reviews"
)

// Service - service layer struct
type Service struct {
	Repository reviews.Repository
}

// New - function for creating new service
func New(repo reviews.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
