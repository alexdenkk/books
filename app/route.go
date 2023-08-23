package app

import (
	"alexdenkk/books/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Route - function for routing services
func (app *App) Route() {
	r := mux.NewRouter()

	app.RouteUsers(r)
	app.RouteBooks(r)
	app.RouteReviews(r)

	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "web/greet.html")
		},
	)

	r.Use(middleware.LoggerMW)

	app.Server.Handler = r
}

// RouteUsers - function for routing users service pages
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

// RouteBooks - function for routing books service pages
func (app *App) RouteBooks(r *mux.Router) {
	sub := r.PathPrefix("/book").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.BooksHandler.Get).
		Methods("GET")
	sub.HandleFunc("/create/", app.MW.Auth(app.BooksHandler.Create)).
		Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.BooksHandler.Update)).
		Methods("PUT")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.BooksHandler.Delete)).
		Methods("DELETE")
	sub.HandleFunc("/all/", app.BooksHandler.GetAll).
		Methods("GET")
}

// RouteReviews - function for routing reviews service pages
func (app *App) RouteReviews(r *mux.Router) {
	sub := r.PathPrefix("/book").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.ReviewsHandler.Get).
		Methods("GET")
	sub.HandleFunc("/create/", app.MW.Auth(app.ReviewsHandler.Create)).
		Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.ReviewsHandler.Update)).
		Methods("PUT")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.ReviewsHandler.Delete)).
		Methods("DELETE")
	sub.HandleFunc("/for/{id}:[0-9]+/", app.ReviewsHandler.GetFor).
		Methods("GET")
	sub.HandleFunc("/for/{id}:[0-9]+/", app.MW.Auth(app.ReviewsHandler.DeleteFor))
}

// RouteGenres - function for routing genres service pages
func (app *App) RouteGenres(r *mux.Router) {
	sub := r.PathPrefix("/genre").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.GenresHandler.Get).
		Methods("GET")
	sub.HandleFunc("/create/", app.MW.Auth(app.GenresHandler.Create)).
		Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.GenresHandler.Update)).
		Methods("PUT")
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.GenresHandler.Delete)).
		Methods("DELETE")
	sub.HandleFunc("/all/", app.GenresHandler.GetAll).
		Methods("GET")
}
