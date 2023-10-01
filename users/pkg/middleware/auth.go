package middleware

import (
	"alexdenkk/books/users/pkg/token"
	"net/http"
	"strings"
)

// HandlerFunc type - type for handler functions using Auth Mmiddleware
type HandlerFunc func(http.ResponseWriter, *http.Request, *token.Claims)

// Auth - function for checking user authorization
func (mw *Middleware) Auth(f HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t := strings.Split(r.Header.Get("Authorization"), " ")[1]

		claims, err := token.ParseJWT(t, mw.SignKey)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		f(w, r, claims)
	})
}
