package validation

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func CreateSpendingType(req *domain.RequestCreateSpendingType) error {
	err := map[string][]string{}

	if req.ProfileID == "" {
		return util.ErrHTTPString("", http.StatusForbidden)
	}
	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return util.ErrHTTPString("", http.StatusForbidden)
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
		return util.ErrHTTP400(err)
	}

	return nil
}

func UpdateSpendingType(req *domain.RequestUpdateSpendingType) error {
	err := map[string][]string{}

	if req.ProfileID == "" {
		return util.ErrHTTPString("", http.StatusForbidden)
	}
	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return util.ErrHTTPString("", http.StatusForbidden)
	}

	if req.ID == "" {
		return util.ErrHTTPString("", http.StatusForbidden)
	}
	if _, err := uuid.Parse(req.ID); err != nil {
		return util.ErrHTTPString("", http.StatusForbidden)
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
		return util.ErrHTTP400(err)
	}

	return nil
}
