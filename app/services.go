package app

import (
	// users
	users_handler "alexdenkk/books/internal/users/gateway/http"
	users_repository "alexdenkk/books/internal/users/repository"
	users_service "alexdenkk/books/internal/users/service"

	// books
	books_handler "alexdenkk/books/internal/books/gateway/http"
	books_repository "alexdenkk/books/internal/books/repository"
	books_service "alexdenkk/books/internal/books/service"

	// reviews
	reviews_handler "alexdenkk/books/internal/reviews/gateway/http"
	reviews_repository "alexdenkk/books/internal/reviews/repository"
	reviews_service "alexdenkk/books/internal/reviews/service"

	// genres
	genres_handler "alexdenkk/books/internal/genres/gateway/http"
	genres_repository "alexdenkk/books/internal/genres/repository"
	genres_service "alexdenkk/books/internal/genres/service"
)

// users
func (app *App) InitUsersService() {
	// repository
	repository := users_repository.New(app.DB)
	// service
	service := users_service.New(repository, app.SignKey)
	// handler
	handler := users_handler.New(service)
	app.UsersHandler = handler
}

func (app *App) InitBooksService() {
	// repository
	repository := books_repository.New(app.DB)
	// service
	service := books_service.New(repository)
	// handler
	handler := books_handler.New(service)
	app.BooksHandler = handler
}

func (app *App) InitReviewsService() {
	// repository
	repository := reviews_repository.New(app.DB)
	// service
	service := reviews_service.New(repository)
	// handler
	handler := reviews_handler.New(service)
	app.ReviewsHandler = handler
}

func (app *App) InitGenresService() {
	// repository
	repository := genres_repository.New(app.DB)
	// service
	service := genres_service.New(repository)
	// handler
	handler := genres_handler.New(service)
	app.GenresHandler = handler
}
