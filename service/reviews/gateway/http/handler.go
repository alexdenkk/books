package http

import (
	"alexdenkk/books/service/reviews"
)

type Handler struct {
	Service reviews.Service
}

func New(s reviews.Service) *Handler {
	return &Handler{
		Service: s,
	}
}
