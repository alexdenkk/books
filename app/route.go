package app

import (
	"alexdenkk/books/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *App) Route() {
	r := mux.NewRouter()

	app.RouteUsers(r)

	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "web/greet.html")
		},
	)

	r.Use(middleware.LoggerMW)

	app.Server.Handler = r
}

func (app *App) RouteUsers(r *mux.Router) {
	sub := r.PathPrefix("/user").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.UsersHandler.Get).
		Methods("GET")
	sub.HandleFunc("/create/", app.MW.Auth(app.UsersHandler.Create)).
		Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.UsersHandler.Update)).
		Methods("PUT")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.UsersHandler.Delete)).
		Methods("DELETE")
	sub.HandleFunc("/login/", app.UsersHandler.Login).
		Methods("POST")
	sub.HandleFunc("/register/", app.UsersHandler.Register).
		Methods("POST")
	sub.HandleFunc("/change-password/", app.MW.Auth(app.UsersHandler.ChangePassword)).
		Methods("POST")
}
