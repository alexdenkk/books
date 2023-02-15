package http

import (
	"alexdenkk/books/pkg/token"
	"context"
	"encoding/json"
	"net/http"
)

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	var request ChangePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(context.Background(), "request", r)

	err := h.Service.ChangePassword(ctx, request.New, request.Old, act)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
