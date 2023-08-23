package http

import (
	"context"
	"encoding/json"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var request AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(context.Background(), "request", r)

	token, err := h.Service.Login(ctx, request.Login, request.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{Token: token})
}
