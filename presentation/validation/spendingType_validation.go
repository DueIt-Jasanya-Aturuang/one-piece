package validation

import (
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateSpendingType(req *domain.RequestCreateSpendingType) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}

	if req.Title == "" {
		err["title"] = append(err["title"], required)
	}
	title := maxMinString(req.Title, 3, 22)
	if title != "" {
		err["title"] = append(err["title"], title)
	}

	if req.MaximumLimit == 0 {
		err["maximum_limit"] = append(err["maximum_limit"], required)
	}
	maximumLimit := maxMinNumeric(req.MaximumLimit, 1000, 999999999)
	if maximumLimit != "" {
		err["maximum_limit"] = append(err["maximum_limit"], maximumLimit)
	}

	if req.Icon == "" {
		err["icon"] = append(err["icon"], required)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}

func UpdateSpendingType(req *domain.RequestUpdateSpendingType) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}

	if _, err := ulid.Parse(req.ID); err != nil {
		return _error.HttpErrString("invalid id", response.CM05)
	}

	if req.Title == "" {
		err["title"] = append(err["title"], required)
	}
	title := maxMinString(req.Title, 3, 22)
	if title != "" {
		err["title"] = append(err["title"], title)
	}

	if req.MaximumLimit == 0 {
		err["maximum_limit"] = append(err["maximum_limit"], required)
	}
	maximumLimit := maxMinNumeric(req.MaximumLimit, 1000, 999999999)
	if maximumLimit != "" {
		err["maximum_limit"] = append(err["maximum_limit"], maximumLimit)
	}

	if req.Icon == "" {
		err["icon"] = append(err["icon"], required)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
