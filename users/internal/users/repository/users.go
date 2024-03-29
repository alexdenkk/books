package repository

import (
	"alexdenkk/books/users/model"
	"context"
)

// Get - repository layer function for getting user record by id
func (r *Repository) Get(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	result := r.DB.First(&user, id)
	return user, result.Error
}

// Create - repository layer function for creating user record
func (r *Repository) Create(ctx context.Context, user model.User) error {
	return r.DB.Create(&user).Error
}

// Update - repository layer function for updating user record
func (r *Repository) Update(ctx context.Context, user model.User) error {
	return r.DB.Save(&user).Error
}

// Delete - repository layer function for deleting user record by id
func (r *Repository) Delete(ctx context.Context, id uint) error {
	return r.DB.Delete(&model.User{}, id).Error
}

// GetByLogin - repository layer function for getting user record by login
func (r *Repository) GetByLogin(ctx context.Context, login string) (model.User, error) {
	var user model.User
	result := r.DB.Where("login = ?", login).First(&user)
	return user, result.Error
}
