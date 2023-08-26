package http

import (
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/token"
	"context"
	"encoding/json"
	"net/http"
)

// Create - Gateway layer function for creating book
func (h *Handler) Create(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(context.Background(), "request", r)

	err := h.Service.Create(ctx, book, act)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
