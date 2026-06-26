package errors

import "errors"

var (
	ErrInternalServerError = errors.New("an unexpected errors occurred. Please try again later")
)
