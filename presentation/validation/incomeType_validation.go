package validation

import (
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateIncomeType(req *domain.RequestCreateIncomeType) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}

	if req.Name == "" {
		err["name"] = append(err["name"], required)
	}
	name := maxMinString(req.Name, 3, 22)
	if name != "" {
		err["name"] = append(err["name"], name)
	}

	if req.Icon == "" {
		err["icon"] = append(err["icon"], required)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}

func UpdateIncomeType(req *domain.RequestUpdateIncomeType) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}

	if _, err := uuid.Parse(req.ID); err != nil {
		return _error.HttpErrString("not found", response.CM01)
	}

	if req.Name == "" {
		err["name"] = append(err["name"], required)
	}
	name := maxMinString(req.Name, 3, 22)
	if name != "" {
		err["name"] = append(err["name"], name)
	}

	if req.Icon == "" {
		err["icon"] = append(err["icon"], required)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
