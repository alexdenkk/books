package genres

import (
	"errors"
)

var (
	ShortFieldError = errors.New("short field")
	PermissionError = errors.New("no permissions to perform the operation")
)
