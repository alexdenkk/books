package genres

import (
	"alexdenkk/books/model"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]model.Genre, error)
	Get(ctx context.Context, id uint) (model.Genre, error)
	Create(ctx context.Context, genre model.Genre) error
	Update(ctx context.Context, genre model.Genre) error
	Delete(ctx context.Context, id uint) error
}
