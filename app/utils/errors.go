package utils

import (
	"errors"
	"fmt"
	"net/http"
	"sort"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/gommon/log"
)

var (
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnauthorized        = errors.New("unauthorized")
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HttpErr interface {
	Status() int
	Error() string
	Details() interface{}
}

// HttpError struct
type HttpError struct {
	ErrStatus  int         `json:"status"`
	ErrError   string      `json:"error"`
	ErrDetails interface{} `json:"details"`
}

// Error  Error() interface method
func (e HttpError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - details: %v", e.ErrStatus, e.ErrError, e.ErrDetails)
}

// Error status
func (e HttpError) Status() int {
	return e.ErrStatus
}

// HttpError Details
func (e HttpError) Details() interface{} {
	return e.ErrDetails
}

// New Internal Server Error
func NewInternalServerError(details interface{}) HttpErr {
	log.Error(details.(error).Error())

	return HttpError{
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   ErrInternalServerError.Error(),
		ErrDetails: nil,
	}
}

func Respond(code int, data interface{}, message string) (response Response) {
	return Response{
		Status:  code,
		Message: message,
		Data:    data,
	}
}

func NewUnauthorizedError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   ErrUnauthorized.Error(),
		ErrDetails: details,
	}
}

func NewBadRequestError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusBadRequest,
		ErrError:   ErrBadRequest.Error(),
		ErrDetails: details,
	}
}

// New Unprocessable Entity Error
func NewUnprocessableEntityError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   ErrUnprocessableEntity.Error(),
		ErrDetails: details,
	}
}
func ParseHttpError(err error) (int, interface{}) {
	if httpErr, ok := err.(HttpErr); ok {
		return httpErr.Status(), httpErr
	}
	return http.StatusInternalServerError, NewInternalServerError(err)
}

func NewInvalidInputError(errs validation.Errors) HttpErr {
	type invalidField struct {
		Field string `json:"field"`
		Error string `json:"error"`
	}

	var details []invalidField
	var fields []string
	for field := range errs {
		fields = append(fields, field)
	}

	sort.Strings(fields)
	for _, field := range fields {
		details = append(details, invalidField{
			Field: field,
			Error: errs[field].Error(),
		})
	}

	return HttpError{
		ErrStatus:  http.StatusBadRequest,
		ErrError:   ErrBadRequest.Error(),
		ErrDetails: details,
	}
}
