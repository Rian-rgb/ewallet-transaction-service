package errs

import "net/http"

func HTTPStatus(err *Error) int {
	switch err.Code {

	case ErrBadRequest:
		return http.StatusBadRequest

	case ErrUnauthorized:
		return http.StatusUnauthorized

	case ErrNotFound, ErrTransactionNotFound:
		return http.StatusNotFound

	case ErrInvalidStatusTransition:
		return http.StatusUnprocessableEntity

	default:
		return http.StatusInternalServerError
	}
}
