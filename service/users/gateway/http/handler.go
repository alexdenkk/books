package http

import (
	"alexdenkk/books/service/users"
)

type Handler struct {
	Service users.Service
}

func New(service users.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
