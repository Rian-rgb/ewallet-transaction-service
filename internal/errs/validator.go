package errs

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func FromValidator(err error) *Error {
	var ve validator.ValidationErrors

	if ok := errors.As(err, &ve); !ok {
		return New(ErrBadRequest, err.Error())
	}

	var fieldErrors []FieldError

	for _, fe := range ve {

		field := fe.Field()

		var msg string

		switch fe.Tag() {

		case "required":
			msg = fmt.Sprintf("%s is required", field)

		case "gt":
			msg = fmt.Sprintf("%s must be greater than %s", field, fe.Param())

		case "oneof":
			msg = fmt.Sprintf("%s must be one of [%s]", field, fe.Param())

		default:
			msg = fmt.Sprintf("%s is invalid", field)
		}

		fieldErrors = append(fieldErrors, FieldError{
			Field:   toSnakeCase(field),
			Message: msg,
		})
	}

	return NewValidation(fieldErrors)
}
