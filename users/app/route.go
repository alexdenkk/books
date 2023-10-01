package app

import (
	"alexdenkk/books/users/pkg/middleware"
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

	sub := r.PathPrefix("/user").Subrouter()

	// Functions routing
	sub.HandleFunc("/{id:[0-9]+}/", app.Handler.Get).
		Methods(http.MethodGet)
	sub.HandleFunc("/create/", app.MW.Auth(app.Handler.Create)).
		Methods(http.MethodPost)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.Handler.Update)).
		Methods(http.MethodPut)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.Handler.Delete)).
		Methods(http.MethodDelete)
	sub.HandleFunc("/login/", app.Handler.Login).
		Methods(http.MethodPost)
	sub.HandleFunc("/register/", app.Handler.Register).
		Methods(http.MethodPost)
	sub.HandleFunc("/change-password/", app.MW.Auth(app.Handler.ChangePassword)).
		Methods(http.MethodPost)

	app.Server.Handler = r
}
