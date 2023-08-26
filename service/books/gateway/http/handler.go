package http

import (
	"alexdenkk/books/service/books"
)

// Handler - gateway layer struct
type Handler struct {
	Service books.Service
}

// New - function for creating new handler
func New(s books.Service) *Handler {
	return &Handler{
		Service: s,
	}
}
