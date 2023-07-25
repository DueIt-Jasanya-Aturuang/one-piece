package exception

import (
	"fmt"
	"net/http"
)

type RespError struct {
	Code    int
	Message any
}

func (r *RespError) Error() string {
	return fmt.Sprintf("%d: %d", r.Code, r.Message)
}

func InternalServerError(msg string) error {
	return &RespError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func RequestTimeOut(msg string) error {
	return &RespError{
		Code:    http.StatusRequestTimeout,
		Message: msg,
	}
}

func BadRequest(msg map[string][]string) error {
	return &RespError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func UnprocessableEntity(msg map[string]map[string]string) error {
	return &RespError{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
	}
}

func NotFound(msg string) error {
	return &RespError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func Unauthorization(msg string) error {
	return &RespError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func Forbidden(msg string) error {
	return &RespError{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}
