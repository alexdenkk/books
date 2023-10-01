package http

import (
	"alexdenkk/books/reviews/model"
	"alexdenkk/books/reviews/pkg/token"
	"context"
	"encoding/json"
	"net/http"
)

// Create - Gateway layer function for creating review
func (h *Handler) Create(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	var review model.Review

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(context.Background(), "request", r)

	err := h.Service.Create(ctx, review, act)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
