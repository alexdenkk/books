package app

import (
	"alexdenkk/books/reviews/pkg/middleware"
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

	sub := r.PathPrefix("/review").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.Handler.Get).
		Methods(http.MethodGet)
	sub.HandleFunc("/create/", app.MW.Auth(app.Handler.Create)).
		Methods(http.MethodPost)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.Handler.Update)).
		Methods(http.MethodPut)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.Handler.Delete)).
		Methods(http.MethodDelete)
	sub.HandleFunc("/for/{id:[0-9]+}/", app.Handler.GetFor).
		Methods(http.MethodGet)
	sub.HandleFunc("/for/{id:[0-9]+}/", app.MW.Auth(app.Handler.DeleteFor)).
		Methods(http.MethodDelete)

	app.Server.Handler = r
}
