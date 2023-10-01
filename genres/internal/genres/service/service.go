package service

import (
	"alexdenkk/books/genres/internal/genres"
)

// Service - service layer struct
type Service struct {
	Repository genres.Repository
	SignKey    []byte
}

// New - function for creating new service
func New(repo genres.Repository, key []byte) *Service {
	return &Service{
		Repository: repo,
		SignKey:    key,
	}
}
