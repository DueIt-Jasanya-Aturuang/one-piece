package middleware

import (
	"net/http"
)

func SetAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newToken := r.Header.Get("Authorization")
		if newToken != "" {
			w.Header().Set("Authorization", newToken)
		}
		next.ServeHTTP(w, r)
	})
}
