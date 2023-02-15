package books

import (
	"alexdenkk/books/model"
	"context"
)

type Repository interface {
	Get(ctx context.Context, id uint) (model.Book, error)
	Create(ctx context.Context, book model.Book) error
	Update(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, id uint) error
}
