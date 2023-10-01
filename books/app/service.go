package app

import (
	handler "alexdenkk/books/books/internal/books/gateway/http"
	repository "alexdenkk/books/books/internal/books/repository"
	service "alexdenkk/books/books/internal/books/service"
)

// InitService - function for initializing service
func (app *App) InitService() {
	// repository
	repository := repository.New(app.DB)
	// service
	service := service.New(repository)
	// handler
	h := handler.New(service)
	app.Handler = h
}
