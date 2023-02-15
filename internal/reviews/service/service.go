package service

import (
	"alexdenkk/books/internal/reviews"
)

type Service struct {
	Repository reviews.Repository
}

func New(repo reviews.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
