package repository

import (
	"alexdenkk/books/reviews/model"
	"context"
)

// Get - repository layer function for getting review record by id
func (r *Repository) Get(ctx context.Context, id uint) (model.Review, error) {
	var review model.Review
	result := r.DB.First(&review, id)
	return review, result.Error
}

// Create - repository layer function for creating review record
func (r *Repository) Create(ctx context.Context, review model.Review) error {
	result := r.DB.Create(&review)
	return result.Error
}

// Update - repository layer function for updating review record
func (r *Repository) Update(ctx context.Context, review model.Review) error {
	result := r.DB.Save(&review)
	return result.Error
}

// Delete - repository layer function for deleting review record by id
func (r *Repository) Delete(ctx context.Context, id uint) error {
	result := r.DB.Delete(&model.Review{}, id)
	return result.Error
}

// GetFor - repository layer function for getting all reviews records for book by ID
func (r *Repository) GetFor(ctx context.Context, bookID uint) ([]model.Review, error) {
	var reviews []model.Review
	result := r.DB.Where("book_id = ?", bookID).Find(&reviews)
	return reviews, result.Error
}

// DeleteFor - repository layer function for deleting all reviews records for book by ID
func (r *Repository) DeleteFor(ctx context.Context, bookID uint) error {
	result := r.DB.Where("book_id = ?", bookID).Delete(&model.Review{})
	return result.Error
}
