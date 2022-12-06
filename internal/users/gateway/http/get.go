package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"context"
	"strconv"
	"encoding/json"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	user, err := h.Service.Get(ctx, uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&user)
}
