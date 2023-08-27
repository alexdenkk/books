package http

import (
	"alexdenkk/books/pkg/token"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Update - Gateway layer function for updating user
func (h *Handler) Update(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	user, err := h.Service.Get(ctx, uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.Update(ctx, user, act)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
