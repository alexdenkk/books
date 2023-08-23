package app

import (
	// users
	users_handler "alexdenkk/books/service/users/gateway/http"
	users_repository "alexdenkk/books/service/users/repository"
	users_service "alexdenkk/books/service/users/service"

	// books
	books_handler "alexdenkk/books/service/books/gateway/http"
	books_repository "alexdenkk/books/service/books/repository"
	books_service "alexdenkk/books/service/books/service"

	// reviews
	reviews_handler "alexdenkk/books/service/reviews/gateway/http"
	reviews_repository "alexdenkk/books/service/reviews/repository"
	reviews_service "alexdenkk/books/service/reviews/service"

	// genres
	genres_handler "alexdenkk/books/service/genres/gateway/http"
	genres_repository "alexdenkk/books/service/genres/repository"
	genres_service "alexdenkk/books/service/genres/service"
)

// InitUsersService - function for initialize users service
func (app *App) InitUsersService() {
	// repository
	repository := users_repository.New(app.DB)
	// service
	service := users_service.New(repository, app.SignKey)
	// handler
	handler := users_handler.New(service)
	app.UsersHandler = handler
}

// InitBooksService - function for initialize books service
func (app *App) InitBooksService() {
	// repository
	repository := books_repository.New(app.DB)
	// service
	service := books_service.New(repository)
	// handler
	handler := books_handler.New(service)
	app.BooksHandler = handler
}

// InitReviewsService - function for initialize reviews service
func (app *App) InitReviewsService() {
	// repository
	repository := reviews_repository.New(app.DB)
	// service
	service := reviews_service.New(repository)
	// handler
	handler := reviews_handler.New(service)
	app.ReviewsHandler = handler
}

// InitGenresService - function for initialize genres service
func (app *App) InitGenresService() {
	// repository
	repository := genres_repository.New(app.DB)
	// service
	service := genres_service.New(repository)
	// handler
	handler := genres_handler.New(service)
	app.GenresHandler = handler
}
