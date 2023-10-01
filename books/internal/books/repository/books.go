package repository

import (
	"alexdenkk/books/books/model"
	"context"
)

// GetAll - repository layer function for getting all books records
func (r *Repository) GetAll(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	result := r.DB.Find(&books)
	return books, result.Error
}

// Get - repository layer function for getting book record by id
func (r *Repository) Get(ctx context.Context, id uint) (model.Book, error) {
	var book model.Book
	result := r.DB.First(&book, id)
	return book, result.Error
}

// Create - repository layer function for creating book record
func (r *Repository) Create(ctx context.Context, book model.Book) error {
	result := r.DB.Create(&book)
	return result.Error
}

// Update - repository layer function for updating book record
func (r *Repository) Update(ctx context.Context, book model.Book) error {
	result := r.DB.Save(&book)
	return result.Error
}

// Delete - repository layer function for deleting book record by id
func (r *Repository) Delete(ctx context.Context, id uint) error {
	result := r.DB.Delete(&model.Book{}, id)
	return result.Error
}
