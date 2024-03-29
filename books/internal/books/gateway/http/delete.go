package http

import (
	"alexdenkk/books/books/pkg/token"
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Delete - gateway layer function for deleting book
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	err := h.Service.Delete(ctx, uint(id), act)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
