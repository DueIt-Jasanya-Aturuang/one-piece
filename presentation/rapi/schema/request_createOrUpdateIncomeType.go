package schema

import (
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type RequestCreateOrUpdateIncomeType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (req *RequestCreateOrUpdateIncomeType) Validate() error {
	err := map[string][]string{}

	if req.Name == "" {
		err["name"] = append(err["name"], util.Required)
	}
	name := util.MaxMinString(req.Name, 3, 22)
	if name != "" {
		err["name"] = append(err["name"], name)
	}

	if req.Icon == "" {
		err["icon"] = append(err["icon"], util.Required)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
