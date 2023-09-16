package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func ErrorResponseEncode(w http.ResponseWriter, err error) {
	var (
		errHTTP          *domain.ErrHTTP
		errUnmarshalType *json.UnmarshalTypeError
		errSyntak        *json.SyntaxError
		errPQ            *pq.Error
	)

	switch {
	case errors.As(err, &errUnmarshalType):
		msg := fmt.Sprintf("UnprocessableEntity : %s", err.Error())
		err = util.ErrHTTPString(msg, http.StatusUnprocessableEntity)
	case errors.As(err, &errSyntak):
		err = util.ErrHTTP400(map[string][]string{
			"unexpected": {
				"unexpected end of json input",
				errSyntak.Error(),
			},
		})
	case errors.As(err, &errPQ):
		log.Warn().Msgf("pqerror | err : %v", err)
		if errPQ.Code == "23503" {
			err = util.ErrHTTPString("", http.StatusForbidden)
		} else {
			err = util.ErrHTTPString("", http.StatusInternalServerError)
		}
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
