package users

import (
	"alexdenkk/books/model"
	"context"
)

// Repository - interface for user repository layer
type Repository interface {
	Get(ctx context.Context, id uint) (model.User, error)
	GetByLogin(ctx context.Context, login string) (model.User, error)
	Create(ctx context.Context, user model.User) error
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id uint) error
}
