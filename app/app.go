package app

import (
	//handlers
	books_handler "alexdenkk/books/service/books/gateway/http"
	genres_handler "alexdenkk/books/service/genres/gateway/http"
	reviews_handler "alexdenkk/books/service/reviews/gateway/http"
	users_handler "alexdenkk/books/service/users/gateway/http"

	"alexdenkk/books/pkg/middleware"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// App - server app struct
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

// Run - function for run app
func (app *App) Run() error {
	app.Route()
	log.Println("app running")
	return app.Server.ListenAndServe()
}

// New - function for create new app
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
