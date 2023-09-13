package helper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func DecodeJson(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err == io.EOF {
		log.Warn().Msgf(util.LogErrDecode, err)
		return util.ErrHTTP400(map[string][]string{
			"bad_request": {
				"tidak ada request body",
			},
		})
	}

	if err != nil {
		log.Warn().Msgf(util.LogErrDecode, err)
		return err
	}

	return nil
}
