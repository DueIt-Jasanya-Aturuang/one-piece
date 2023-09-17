package validation

import (
	"github.com/google/uuid"
	errResp "github.com/jasanya-tech/jasanya-response-backend-golang"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateSpendingType(req *domain.RequestCreateSpendingType) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return errResp.HttpErrString("invalid profile-id", errResp.S403)
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
		return errResp.HttpErrMapOfSlices(err, errResp.S400)
	}

	return nil
}

func UpdateSpendingType(req *domain.RequestUpdateSpendingType) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return errResp.HttpErrString("invalid profile-id", errResp.S403)
	}

	if _, err := uuid.Parse(req.ID); err != nil {
		return errResp.HttpErrString("invalid id", errResp.S403)
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
		return errResp.HttpErrMapOfSlices(err, errResp.S400)
	}

	return nil
}
