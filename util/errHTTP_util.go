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

func ErrHTTPString(msg string, code int) error {
	if msg == "" {
		switch code {
		case 404:
			msg = "Data Not Found"
		case 403:
			msg = "Forbidden"
		case 401:
			msg = "Unauthorization"
		case 408:
			msg = "Request Time Out"
		case 409:
			msg = "Data Conflict"
		case 422:
			msg = "Unprocessable Entity"
		case 500:
			msg = "Internal Server Error"
		}
	}

	return &domain.ErrHTTP{
		Code:    code,
		Message: msg,
	}
}
