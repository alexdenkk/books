package users

import (
	"errors"
)

var (
	ShortFieldError = errors.New("too short field")
	PermissionsError = errors.New("no permissions to perform the operation")
	LoginAlreadyExistsError = errors.New("user with this login already exists")
	ChangePasswordError = errors.New("cannot change password by this function")
	IncorrectPasswordError = errors.New("incorrect password")
)
