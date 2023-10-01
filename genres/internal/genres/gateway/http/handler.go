package http

import (
	"alexdenkk/books/genres/internal/genres"
)

// Handler - gateway layer struct
type Handler struct {
	Service genres.Service
}

// New - function for creating new handler
func New(service genres.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
