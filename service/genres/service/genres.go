package service

import (
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"alexdenkk/books/service/genres"
	"context"
)

// GetAll - service layer function for getting all genres records
func (s *Service) GetAll(ctx context.Context) ([]model.Genre, error) {
	return s.Repository.GetAll(ctx)
}

// Get - service layer function for getting genre record by ID
func (s *Service) Get(ctx context.Context, id uint) (model.Genre, error) {
	return s.Repository.Get(ctx, id)
}

// Create - service layer function for creating a genre record
func (s *Service) Create(ctx context.Context, genre model.Genre, act *token.Claims) error {
	if len(genre.Name) < 4 {
		return genres.ShortFieldError
	}

	if act.Role != model.ADMIN {
		return genres.PermissionError
	}

	return s.Repository.Create(ctx, genre)
}

// Update - service layer function for updating a genre record
func (s *Service) Update(ctx context.Context, genre model.Genre, act *token.Claims) error {
	if len(genre.Name) < 4 {
		return genres.ShortFieldError
	}

	if act.Role != model.ADMIN {
		return genres.PermissionError
	}

	return s.Repository.Update(ctx, genre)
}

// Delete - service layer function for deleting a genre record
func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return genres.PermissionError
	}

	return s.Repository.Delete(ctx, id)
}
