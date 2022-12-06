package app

import (
	//handlers
	users_handler "alexdenkk/books/internal/users/gateway/http"

	"alexdenkk/books/pkg/middleware"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type App struct {
	UsersHandler *users_handler.Handler
	Server       *http.Server
	MW       	 *middleware.Middleware
	SignKey      []byte
	DB           *gorm.DB
}

func (app *App) Run() error {
	app.Route()
	return app.Server.ListenAndServe()
}

func New(db *gorm.DB, key []byte, addr string) *App {
	app := &App{
		DB: db,
		SignKey: key,
	}

	app.Server = &http.Server{
		Addr: addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.MW = middleware.New(app.SignKey)

	app.InitUsersService()

	return app
}
