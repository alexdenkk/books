package http

import (
	"alexdenkk/books/pkg/token"
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DeleteFor - Gateway layer function for deleting all reviews for a book
func (h *Handler) DeleteFor(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	err := h.Service.DeleteFor(ctx, uint(id), act)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
