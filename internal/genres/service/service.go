package service

import (
	"alexdenkk/books/internal/genres"
)

type Service struct {
	Repository genres.Repository
}

func New(repo genres.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
