package helper

import (
	"net/http"
)

func MethodNotAllowed(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
	return
}
