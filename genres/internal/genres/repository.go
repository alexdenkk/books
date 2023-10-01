package genres

import (
	"alexdenkk/books/genres/model"
	"context"
)

// Repository - interface for genres repository layer
type Repository interface {
	GetAll(ctx context.Context) ([]model.Genre, error)
	Get(ctx context.Context, id uint) (model.Genre, error)
	Create(ctx context.Context, genre model.Genre) error
	Update(ctx context.Context, genre model.Genre) error
	Delete(ctx context.Context, id uint) error
}
