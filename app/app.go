package app

import (
	//handlers
	books_handler "alexdenkk/books/internal/books/gateway/http"
	genres_handler "alexdenkk/books/internal/genres/gateway/http"
	reviews_handler "alexdenkk/books/internal/reviews/gateway/http"
	users_handler "alexdenkk/books/internal/users/gateway/http"

	"alexdenkk/books/pkg/middleware"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type App struct {
	UsersHandler   *users_handler.Handler
	BooksHandler   *books_handler.Handler
	ReviewsHandler *reviews_handler.Handler
	GenresHandler  *genres_handler.Handler
	Server         *http.Server
	MW             *middleware.Middleware
	SignKey        []byte
	DB             *gorm.DB
}

func (app *App) Run() error {
	app.Route()
	log.Println("app running")
	return app.Server.ListenAndServe()
}

func New(db *gorm.DB, key []byte, addr string) *App {
	app := &App{
		DB:      db,
		SignKey: key,
	}

	app.Server = &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.MW = middleware.New(app.SignKey)

	app.InitUsersService()
	app.InitBooksService()
	app.InitReviewsService()
	app.InitGenresService()

	return app
}
