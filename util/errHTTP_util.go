package util

import (
	"net/http"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func ErrHTTP400(msg map[string][]string) error {
	return &domain.ErrHTTP{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func ErrHTTP422(msg map[string][]string) error {
	return &domain.ErrHTTP{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func ErrHTTPString(msg string, code int) error {
	if msg == "" {
		switch code {
		case 404:
			msg = "DATA NOT FOUND"
		case 403:
			msg = "FORBIDDEN"
		case 401:
			msg = "UNAUTHORIZATION"
		case 409:
			msg = "DATA CONFLICT"
		case 500:
			msg = "INTERNAL SERVER ERROR"
		}
	}

	return &domain.ErrHTTP{
		Code:    code,
		Message: msg,
	}
}
