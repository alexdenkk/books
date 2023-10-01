package reviews

import (
	"alexdenkk/books/reviews/model"
	"alexdenkk/books/reviews/pkg/token"
	"context"
)

// Service - interface for reviews service layer
type Service interface {
	Get(ctx context.Context, id uint) (model.Review, error)
	Create(ctx context.Context, review model.Review, act *token.Claims) error
	Update(ctx context.Context, review model.Review, act *token.Claims) error
	Delete(ctx context.Context, id uint, act *token.Claims) error
	GetFor(ctx context.Context, bookID uint) ([]model.Review, error)
	DeleteFor(ctx context.Context, bookID uint, act *token.Claims) error
}
