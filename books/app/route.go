package app

import (
	"alexdenkk/books/books/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Route - function for routing
func (app *App) Route() {
	r := mux.NewRouter()

	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "web/greet.html")
		},
	)

	r.Use(middleware.LoggerMW)

	sub := r.PathPrefix("/book").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.Handler.Get).
		Methods(http.MethodGet)
	sub.HandleFunc("/create/", app.MW.Auth(app.Handler.Create)).
		Methods(http.MethodPost)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.Handler.Update)).
		Methods(http.MethodPut)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.Handler.Delete)).
		Methods(http.MethodDelete)
	sub.HandleFunc("/all/", app.Handler.GetAll).
		Methods(http.MethodGet)

	app.Server.Handler = r
}
