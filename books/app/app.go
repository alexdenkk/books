package app

import (
	//handlers
	handler "alexdenkk/books/books/internal/books/gateway/http"

	"alexdenkk/books/books/pkg/middleware"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// App - server service app struct
type App struct {
	Handler *handler.Handler
	Server  *http.Server
	MW      *middleware.Middleware
	SignKey []byte
	DB      *gorm.DB
}

// Run - function for run service app
func (app *App) Run() error {
	app.Route()
	log.Println("books service running")
	return app.Server.ListenAndServe()
}

// New - function for create new service app
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

	app.InitService()

	return app
}
