package http

import (
	"alexdenkk/books/service/reviews"
)

// Handler - gateway layer struct
type Handler struct {
	Service reviews.Service
}

// New - function for creating new handler
func New(s reviews.Service) *Handler {
	return &Handler{
		Service: s,
	}
}
