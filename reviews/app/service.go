package app

import (
	handler "alexdenkk/books/reviews/internal/reviews/gateway/http"
	repository "alexdenkk/books/reviews/internal/reviews/repository"
	service "alexdenkk/books/reviews/internal/reviews/service"
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
