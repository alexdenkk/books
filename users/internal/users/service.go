package users

import (
	"alexdenkk/books/users/model"
	"alexdenkk/books/users/pkg/token"
	"context"
)

// Service - interface for user service layer
type Service interface {
	Get(ctx context.Context, id uint) (model.User, error)
	Create(ctx context.Context, user model.User, act *token.Claims) error
	Update(ctx context.Context, user model.User, act *token.Claims) error
	Delete(ctx context.Context, id uint, act *token.Claims) error

	Login(ctx context.Context, login, password string) (string, error)
	Register(ctx context.Context, login, password string) error
	ChangePassword(ctx context.Context, newPswd, oldPswd string, act *token.Claims) error
}
