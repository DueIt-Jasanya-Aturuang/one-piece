package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	resp "github.com/jasanya-tech/jasanya-response-backend-golang"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func ErrorResponseEncode(w http.ResponseWriter, err error) {
	var (
		errHttp          *resp.HttpError
		errUnmarshalType *json.UnmarshalTypeError
		errSyntak        *json.SyntaxError
		errPQ            *pq.Error
	)

	switch {
	case errors.As(err, &errUnmarshalType):
		msg := fmt.Sprintf("UnprocessableEntity : %s", err.Error())
		err = resp.HttpErrString(msg, resp.S422)
	case errors.As(err, &errSyntak):
		msg := fmt.Sprintf("unexpected end of json input : %s", err.Error())
		err = resp.HttpErrString(msg, resp.S422)
	case errors.As(err, &errPQ):
		log.Warn().Msgf("pqerror | err : %v", err)
		if errPQ.Code == "23503" {
			err = resp.HttpErrString(string(resp.S403), resp.S403)
		} else {
			err = resp.HttpErrString(string(resp.S500), resp.S500)
		}
	case errors.Is(err, context.DeadlineExceeded):
		err = resp.HttpErrString("request time out", resp.S408)
	}

	ok := errors.As(err, &errHttp)
	if !ok {
		err = resp.HttpErrString(string(resp.S500), resp.S500)
		errors.As(err, &errHttp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errHttp.Code)
	response := resp.ToResponseError(err)

	if errEncode := json.NewEncoder(w).Encode(response); errEncode != nil {
		log.Err(errEncode).Msgf(util.LogErrEncode, response, errEncode)
	}
}

func SuccessResponseEncode(w http.ResponseWriter, data any, message string) {
	response := resp.ToResponseSuccess(data, message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	
	if errEncode := json.NewEncoder(w).Encode(response); errEncode != nil {
		log.Err(errEncode).Msgf(util.LogErrEncode, response, errEncode)
	}
}
