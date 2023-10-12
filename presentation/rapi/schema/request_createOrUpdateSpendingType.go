package schema

import (
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type RequestCreateOrUpdateSpendingType struct {
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
	Icon         string `json:"icon"`
}

func (req *RequestCreateOrUpdateSpendingType) Validate() error {
	err := map[string][]string{}

	if req.Title == "" {
		err["title"] = append(err["title"], util.Required)
	}
	title := util.MaxMinString(req.Title, 3, 22)
	if title != "" {
		err["title"] = append(err["title"], title)
	}

	if req.MaximumLimit == 0 {
		err["maximum_limit"] = append(err["maximum_limit"], util.Required)
	}
	maximumLimit := util.MaxMinNumeric(req.MaximumLimit, 1000, 999999999)
	if maximumLimit != "" {
		err["maximum_limit"] = append(err["maximum_limit"], maximumLimit)
	}

	if req.Icon == "" {
		err["icon"] = append(err["icon"], util.Required)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
