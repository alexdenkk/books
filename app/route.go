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
		Methods(http.MethodGet)
	sub.HandleFunc("/create/", app.MW.Auth(app.UsersHandler.Create)).
		Methods(http.MethodPost)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.UsersHandler.Update)).
		Methods(http.MethodPut)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.UsersHandler.Delete)).
		Methods(http.MethodDelete)
	sub.HandleFunc("/login/", app.UsersHandler.Login).
		Methods(http.MethodPost)
	sub.HandleFunc("/register/", app.UsersHandler.Register).
		Methods(http.MethodPost)
	sub.HandleFunc("/change-password/", app.MW.Auth(app.UsersHandler.ChangePassword)).
		Methods(http.MethodPost)
}

// RouteBooks - function for routing books service pages
func (app *App) RouteBooks(r *mux.Router) {
	sub := r.PathPrefix("/book").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.BooksHandler.Get).
		Methods(http.MethodGet)
	sub.HandleFunc("/create/", app.MW.Auth(app.BooksHandler.Create)).
		Methods(http.MethodPost)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.BooksHandler.Update)).
		Methods(http.MethodPut)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.BooksHandler.Delete)).
		Methods(http.MethodDelete)
	sub.HandleFunc("/all/", app.BooksHandler.GetAll).
		Methods(http.MethodGet)
}

// RouteReviews - function for routing reviews service pages
func (app *App) RouteReviews(r *mux.Router) {
	sub := r.PathPrefix("/review").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.ReviewsHandler.Get).
		Methods(http.MethodGet)
	sub.HandleFunc("/create/", app.MW.Auth(app.ReviewsHandler.Create)).
		Methods(http.MethodPost)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.ReviewsHandler.Update)).
		Methods(http.MethodPut)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.ReviewsHandler.Delete)).
		Methods(http.MethodDelete)
	sub.HandleFunc("/for/{id}:[0-9]+/", app.ReviewsHandler.GetFor).
		Methods(http.MethodGet)
	sub.HandleFunc("/for/{id}:[0-9]+/", app.MW.Auth(app.ReviewsHandler.DeleteFor)).
		Methods(http.MethodDelete)
}

// RouteGenres - function for routing genres service pages
func (app *App) RouteGenres(r *mux.Router) {
	sub := r.PathPrefix("/genre").Subrouter()

	sub.HandleFunc("/{id:[0-9]+}/", app.GenresHandler.Get).
		Methods(http.MethodGet)
	sub.HandleFunc("/create/", app.MW.Auth(app.GenresHandler.Create)).
		Methods(http.MethodPost)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.GenresHandler.Update)).
		Methods(http.MethodPut)
	sub.HandleFunc("/{id:[0-9]+}/", app.MW.Auth(app.GenresHandler.Delete)).
		Methods(http.MethodDelete)
	sub.HandleFunc("/all/", app.GenresHandler.GetAll).
		Methods(http.MethodGet)
}
