package http

import (
	"alexdenkk/books/pkg/token"
	"github.com/gorilla/mux"
	"net/http"
	"context"
	"strconv"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, claims *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	err := h.Service.Delete(ctx, uint(id), claims)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
