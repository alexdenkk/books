package genres

import (
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]model.Genre, error)
	Get(ctx context.Context, id uint) (model.Genre, error)
	Create(ctx context.Context, genre model.Genre, act *token.Claims) error
	Update(ctx context.Context, genre model.Genre, act *token.Claims) error
	Delete(ctx context.Context, id uint, act *token.Claims) error
}
