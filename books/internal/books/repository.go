package books

import (
	"alexdenkk/books/books/model"
	"context"
)

// Repository - interface for books repository layer
type Repository interface {
	GetAll(ctx context.Context) ([]model.Book, error)
	Get(ctx context.Context, id uint) (model.Book, error)
	Create(ctx context.Context, book model.Book) error
	Update(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, id uint) error
}
