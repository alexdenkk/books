package http

import (
	"alexdenkk/books/service/genres"
)

type Handler struct {
	Service genres.Service
}

func New(service genres.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
