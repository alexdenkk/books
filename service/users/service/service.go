package service

import (
	"alexdenkk/books/service/users"
)

// Service - service layer struct
type Service struct {
	SignKey    []byte
	Repository users.Repository
}

// New - function for creating new service
func New(repo users.Repository, key []byte) *Service {
	return &Service{
		Repository: repo,
		SignKey:    key,
	}
}
