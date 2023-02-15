package reviews

import (
	"errors"
)

var (
	PermissionError = errors.New("no permissions to perform the operation")
	ShortFieldError = errors.New("short field")
	RateLimitError  = errors.New("too high rate")
)
