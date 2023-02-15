package service

import (
	"alexdenkk/books/internal/genres"
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"context"
)

func (s *Service) GetAll(ctx context.Context) ([]model.Genre, error) {
	return s.Repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id uint) (model.Genre, error) {
	return s.Repository.Get(ctx, id)
}

func (s *Service) Create(ctx context.Context, genre model.Genre, act *token.Claims) error {
	if len(genre.Name) < 4 {
		return genres.ShortFieldError
	}

	if act.Role != model.ADMIN {
		return genres.PermissionError
	}

	return s.Repository.Create(ctx, genre)
}

func (s *Service) Update(ctx context.Context, genre model.Genre, act *token.Claims) error {
	if len(genre.Name) < 4 {
		return genres.ShortFieldError
	}

	if act.Role != model.ADMIN {
		return genres.PermissionError
	}

	return s.Repository.Update(ctx, genre)
}

func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return genres.PermissionError
	}

	return s.Repository.Delete(ctx, id)
}
