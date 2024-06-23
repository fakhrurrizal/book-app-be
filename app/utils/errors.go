package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
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
