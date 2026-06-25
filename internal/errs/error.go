package errs

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code    Code         `json:"code"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors,omitempty"`
	Err     error        `json:"-"`
}

func (e *Error) Error() string {
	return string(e.Code) + ": " + e.Message
}

func New(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// wrap error (untuk logging)
func Wrap(code Code, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewValidation(errors []FieldError) *Error {
	return &Error{
		Code:    ErrBadRequest,
		Message: "validation error",
		Errors:  errors,
	}
}
