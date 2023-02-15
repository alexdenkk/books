package books

import (
	"errors"
)

var (
	ShortFieldError  = errors.New("too short field")
	PermissionsError = errors.New("no permissions to perform the operation")
	WrongDateError   = errors.New("wrong date")
	WrongGenreError  = errors.New("wrong genre")
)
