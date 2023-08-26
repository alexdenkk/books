package books

import (
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"context"
)

// Service - interface for books service layer
type Service interface {
	GetAll(ctx context.Context) ([]model.Book, error)
	Get(ctx context.Context, id uint) (model.Book, error)
	Create(ctx context.Context, book model.Book, act *token.Claims) error
	Update(ctx context.Context, book model.Book, act *token.Claims) error
	Delete(ctx context.Context, id uint, act *token.Claims) error
}
