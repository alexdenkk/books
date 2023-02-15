package books

import (
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"context"
)

type Service interface {
	Get(ctx context.Context, id uint) (model.Book, error)
	Create(ctx context.Context, book model.Book, act *token.Claims) error
	Update(ctx context.Context, book model.Book, act *token.Claims) error
	Delete(ctx context.Context, id uint, act *token.Claims) error
}
