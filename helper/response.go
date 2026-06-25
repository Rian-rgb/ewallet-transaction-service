package helper

import (
	"errors"
	"ewallet-transaction/internal/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessResponse struct {
	Code    string      `json:"code" example:"SUCCESS"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BadRequestResponse struct {
	Code    string            `json:"code" example:"BAD_REQUEST"`
	Message string            `json:"message" example:"validation error"`
	Errors  []errs.FieldError `json:"errors,omitempty"`
}

func SendResponseSuccess(c *gin.Context, message string, data interface{}) {
	resp := SuccessResponse{
		Code:    "SUCCESS",
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func SendResponseError(c *gin.Context, err error) {
	var e *errs.Error
	if errors.As(err, &e) {
		resp := ErrorResponse{
			Code:    string(e.Code),
			Message: e.Message,
		}
		c.JSON(errs.HTTPStatus(e), resp)
		return
	}

	// fallback kalau bukan errs.Error
	resp := ErrorResponse{
		Code:    string(errs.ErrInternal),
		Message: "internal server error",
	}

	c.JSON(http.StatusInternalServerError, resp)
}

func SendResponseBadRequest(c *gin.Context, err error) {
	var e *errs.Error

	if errors.As(err, &e) {
		resp := BadRequestResponse{
			Code:    string(e.Code),
			Message: e.Message,
			Errors:  e.Errors,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// fallback kalau bukan errs.Error
	resp := BadRequestResponse{
		Code:    string(errs.ErrBadRequest),
		Message: err.Error(),
	}

	c.JSON(http.StatusBadRequest, resp)
}
