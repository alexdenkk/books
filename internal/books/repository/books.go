package repository

import (
	"alexdenkk/books/model"
	"context"
)

func (r *Repository) Get(ctx context.Context, id uint) (model.Book, error) {
	var book model.Book
	result := r.DB.First(&book, id)
	return book, result.Error
}

func (r *Repository) Create(ctx context.Context, book model.Book) error {
	result := r.DB.Create(&book)
	return result.Error
}

func (r *Repository) Update(ctx context.Context, book model.Book) error {
	result := r.DB.Save(&book)
	return result.Error
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	result := r.DB.Delete(&model.Book{}, id)
	return result.Error
}
