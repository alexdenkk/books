package app

import (
	handler "alexdenkk/books/users/internal/users/gateway/http"
	repository "alexdenkk/books/users/internal/users/repository"
	service "alexdenkk/books/users/internal/users/service"
)

// InitService - function for initializing service
func (app *App) InitService() {
	// repository
	repository := repository.New(app.DB)
	// service
	service := service.New(repository, app.SignKey)
	// handler
	h := handler.New(service)
	app.Handler = h
}
