package app

import (
	handler "alexdenkk/books/genres/internal/genres/gateway/http"
	repository "alexdenkk/books/genres/internal/genres/repository"
	service "alexdenkk/books/genres/internal/genres/service"
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
