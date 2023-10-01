package http

import (
	"alexdenkk/books/users/internal/users"
)

// Handler - gateway layer struct
type Handler struct {
	Service users.Service
}

// New - function for creating new handler
func New(service users.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
