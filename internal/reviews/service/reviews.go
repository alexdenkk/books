package service

import (
	"alexdenkk/books/internal/reviews"
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"context"
)

func (s *Service)  Get(ctx context.Context, id uint) (model.Review, error) {
	return s.Repository.Get(ctx, id)	
}

func (s *Service) Create(ctx context.Context, review model.Review, act *token.Claims) error {
	if len(review.Text) < 4 {
		return reviews.ShortFieldError
	}

	if review.Rate > 10 {
		return reviews.RateLimitError
	}

	review.UserID = act.ID

	return s.Repository.Create(ctx, review)
}

func (s *Service) Update(ctx context.Context, review model.Review, act *token.Claims) error {
	if review.UserID != act.ID {
		return reviews.PermissionError
	}

	if len(review.Text) < 4 {
		return reviews.ShortFieldError
	}

	if review.Rate > 10 {
		return reviews.RateLimitError
	}

	return s.Repository.Update(ctx, review)
}

func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	review , err := s.Repository.Get(ctx, id)

	if err != nil {
		return err
	}

	if review.UserID != act.ID && act.Role != model.ADMIN {
		return reviews.PermissionError
	}

	return s.Repository.Delete(ctx, id)
}

func (s *Service) GetFor(ctx context.Context, bookID uint) ([]model.Review, error) {
	return s.Repository.GetFor(ctx, bookID)
}

func (s *Service) DeleteFor(ctx context.Context, bookID uint, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return reviews.PermissionError
	}

	return s.Repository.DeleteFor(ctx, bookID)
}
