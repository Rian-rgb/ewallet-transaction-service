package errs

type Code string

const (
	// general
	ErrInternal     Code = "INTERNAL_ERROR"
	ErrBadRequest   Code = "BAD_REQUEST"
	ErrUnauthorized Code = "UNAUTHORIZED"
	ErrNotFound     Code = "NOT_FOUND"

	// transaction
	ErrInvalidStatusTransition Code = "INVALID_STATUS_TRANSITION"
	ErrTransactionNotFound     Code = "TRANSACTION_NOT_FOUND"
)
