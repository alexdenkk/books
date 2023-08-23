package service

import (
	"alexdenkk/books/service/reviews"
)

type Service struct {
	Repository reviews.Repository
}

func New(repo reviews.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
