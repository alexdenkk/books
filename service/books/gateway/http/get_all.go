package http

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAll - Gateway layer function for getting all books
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "request", r)

	books, err := h.Service.GetAll(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&books)
}
