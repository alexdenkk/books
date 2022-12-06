package service

import (
	"alexdenkk/books/internal/users"
)

type Service struct {
	SignKey []byte
	Repository users.Repository
}

func New(repo users.Repository, key []byte) *Service {
	return &Service{
		Repository: repo,
		SignKey: key,
	}
}
