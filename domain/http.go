package domain

import (
	"fmt"
)

type ErrHTTP struct {
	Code    int
	Message any
}

func (e *ErrHTTP) Error() string {
	msg := fmt.Sprintf("error http | code : %d | err : %v", e.Code, e.Message)
	return msg
}

type ResponseErrorHTTP struct {
	Errors any `json:"errors"`
}

type ResponseSuccessHTTP struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Code    int    `json:"-"`
}
