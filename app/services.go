package app

import (
	users_handler "alexdenkk/books/internal/users/gateway/http"
	users_service "alexdenkk/books/internal/users/service"
	users_repository "alexdenkk/books/internal/users/repository"
)

func (app *App) InitUsersService() {
	repository := users_repository.New(
		app.DB,
	)

	service := users_service.New(
		repository,
		app.SignKey,
	)

	handler := users_handler.New(
		service,
	)

	app.UsersHandler = handler
}
