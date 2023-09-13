package helper

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func ErrorResponseEncode(w http.ResponseWriter, err error) {
	var (
		errHTTP          *domain.ErrHTTP
		errUnmarshalType *json.UnmarshalTypeError
		errSyntak        *json.SyntaxError
	)

	switch {
	case errors.As(err, &errUnmarshalType):
		err = util.ErrHTTP422(map[string][]string{
			errUnmarshalType.Field: {
				"invalid tipe input, tipe harus tipe data ", errUnmarshalType.Type.String(),
			},
		})
	case errors.As(err, &errSyntak):
		err = util.ErrHTTP400(map[string][]string{
			"unexpected": {
				"unexpected end of json input",
				errSyntak.Error(),
			},
		})
	case errors.Is(err, context.DeadlineExceeded):
		err = util.ErrHTTPString("", http.StatusRequestTimeout)
	}

	ok := errors.As(err, &errHTTP)
	if !ok {
		err = util.ErrHTTPString("", http.StatusInternalServerError)
		errors.As(err, &errHTTP)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errHTTP.Code)
	resp := &domain.ResponseErrorHTTP{
		Errors: errHTTP.Message,
	}

	if errEncode := json.NewEncoder(w).Encode(resp); errEncode != nil {
		log.Err(errEncode).Msgf(util.LogErrEncode, resp, errEncode)
	}
}

func SuccessResponseEncode(w http.ResponseWriter, data domain.ResponseSuccessHTTP) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)
	if errEncode := json.NewEncoder(w).Encode(data); errEncode != nil {
		log.Err(errEncode).Msgf(util.LogErrEncode, data, errEncode)
	}
}
