package service

import (
	"alexdenkk/books/books/internal/books"
	"alexdenkk/books/books/model"
	"alexdenkk/books/books/pkg/token"
	"context"
	"strings"
	"time"
)

// GetAll - service layer function for getting all books records
func (s *Service) GetAll(ctx context.Context) ([]model.Book, error) {
	return s.Repository.GetAll(ctx)
}

// Get - service layer function for getting book record by ID
func (s *Service) Get(ctx context.Context, id uint) (model.Book, error) {
	return s.Repository.Get(ctx, id)
}

// Create - service layer function for creating book record
func (s *Service) Create(ctx context.Context, book model.Book, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return books.PermissionsError
	}

	if len(strings.ReplaceAll(book.Name, " ", "")) < 1 {
		return books.ShortFieldError
	}

	if len(strings.ReplaceAll(book.Author, " ", "")) < 2 {
		return books.ShortFieldError
	}

	if book.Year < -10000 || book.Year > time.Now().Year() {
		return books.WrongDateError
	}

	return s.Repository.Create(ctx, book)
}

// Update - service layer function for updating book record
func (s *Service) Update(ctx context.Context, book model.Book, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return books.PermissionsError
	}

	if len(strings.ReplaceAll(book.Name, " ", "")) < 1 {
		return books.ShortFieldError
	}

	if len(strings.ReplaceAll(book.Author, " ", "")) < 2 {
		return books.ShortFieldError
	}

	if book.Year < -10000 || book.Year > time.Now().Year() {
		return books.WrongDateError
	}

	return s.Repository.Update(ctx, book)
}

// Delete - service layer function for deleting book record
func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return books.PermissionsError
	}

	return s.Repository.Delete(ctx, id)
}
