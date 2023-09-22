package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
)

func AccountMiddlewareInHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		profileIDUrlParam := chi.URLParam(r, "profile-id")
		profileIDHeader := r.Header.Get("Profile-ID")
		if profileIDHeader != profileIDUrlParam {
			helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile account", response.CM05))
			return
		}

		next.ServeHTTP(w, r)
	})
}
