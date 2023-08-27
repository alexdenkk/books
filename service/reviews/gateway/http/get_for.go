package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetFor - Gateway layer function for getting all reviews for a book
func (h *Handler) GetFor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	reviews, err := h.Service.GetFor(ctx, uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reviews)
}
