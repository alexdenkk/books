package service

import (
	"alexdenkk/books/internal/books"
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"context"
	"strings"
	"time"
)

func (s *Service) Get(ctx context.Context, id uint) (model.Book, error) {
	return s.Repository.Get(ctx, id)
}

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

func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return books.PermissionsError
	}

	return s.Repository.Delete(ctx, id)
}
