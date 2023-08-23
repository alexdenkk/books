package http

import (
	"alexdenkk/books/service/books"
)

type Handler struct {
	Service books.Service
}

func New(s books.Service) *Handler {
	return &Handler{
		Service: s,
	}
}
