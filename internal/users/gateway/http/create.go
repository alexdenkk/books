package http

import (
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"net/http"
	"context"
	"encoding/json"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, claims *token.Claims) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(context.Background(), "request", r)

	err := h.Service.Create(ctx, user, claims)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
