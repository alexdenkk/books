package repository

import (
	"alexdenkk/books/model"
	"context"
)

// GetAll - repository layer function for getting all genres records
func (r *Repository) GetAll(ctx context.Context) ([]model.Genre, error) {
	var genres []model.Genre
	result := r.DB.Find(&genres)
	return genres, result.Error
}

// Get - repository layer function for getting genre record by id
func (r *Repository) Get(ctx context.Context, id uint) (model.Genre, error) {
	var genre model.Genre
	result := r.DB.First(&genre, id)
	return genre, result.Error
}

// Create - repository layer function for creating genre record
func (r *Repository) Create(ctx context.Context, genre model.Genre) error {
	result := r.DB.Create(&genre)
	return result.Error
}

// Update - repository layer function for updating genre record
func (r *Repository) Update(ctx context.Context, genre model.Genre) error {
	result := r.DB.Save(&genre)
	return result.Error
}

// Delete - repository layer function for deleting genre record by id
func (r *Repository) Delete(ctx context.Context, id uint) error {
	result := r.DB.Delete(&model.Genre{}, id)
	return result.Error
}
