package http

import (
	"alexdenkk/books/internal/books"
)

type Handler struct {
	Service books.Service
}

func New(s books.Service) *Handler {
	return &Handler{
		Service: s,
	}
}
