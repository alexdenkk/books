package reviews

import (
	"alexdenkk/books/model"
	"context"
)

type Repository interface {
	Get(ctx context.Context, id uint) (model.Review, error)
	Create(ctx context.Context, review model.Review) error
	Update(ctx context.Context, review model.Review) error
	Delete(ctx context.Context, id uint) error
	GetFor(ctx context.Context, bookID uint) ([]model.Review, error)
	DeleteFor(ctx context.Context, bookID uint) error
}
