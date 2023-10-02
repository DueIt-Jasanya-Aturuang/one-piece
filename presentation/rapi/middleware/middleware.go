package middleware

import (
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
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

func CheckApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		if apiKey != infra.AppApiKeyFinance {
			log.Warn().Msgf("invalid api key | key : %s", apiKey)
			helper.ErrorResponseEncode(w, _error.HttpErrString("forbidden", response.CM05))
			return
		}
		next.ServeHTTP(w, r)
	})
}
