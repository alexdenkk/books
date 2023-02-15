package http

import (
	"alexdenkk/books/internal/genres"
)

type Handler struct {
	Service genres.Service
}

func New(service genres.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
