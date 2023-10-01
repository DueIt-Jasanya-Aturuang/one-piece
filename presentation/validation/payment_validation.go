package validation

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreatePayment(req *domain.RequestCreatePayment) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}

	if req.Image == nil {
		err["image"] = append(err["image"], required)
	}
	if req.Image != nil && req.Image.Size > 0 {
		if req.Image.Size > 2097152 {
			err["image"] = append(err["image"], fmt.Sprintf(fileSize, 2048, 2))
		}
		if !checkContentType(req.Image.Header.Get("Content-Type"), image) {
			err["image"] = append(err["image"], fmt.Sprintf(fileContent, strings.Join(contentType(image), " or ")))
		}
	}

	if req.Name == "" {
		err["name"] = append(err["name"], required)
	}
	name := maxMinString(req.Name, 3, 32)
	if name != "" {
		err["name"] = append(err["name"], name)
	}

	if req.Description != "" {
		description := maxMinString(req.Description, 3, 50)
		if description != "" {
			err["description"] = append(err["description"], description)
		}
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}
	return nil
}

func UpdatePayment(req *domain.RequestUpdatePayment) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}
	
	if req.Image != nil {
		if req.Image != nil && req.Image.Size > 0 {
			if req.Image.Size > 2097152 {
				err["image"] = append(err["image"], fmt.Sprintf(fileSize, 2048, 2))
			}
			if !checkContentType(req.Image.Header.Get("Content-Type"), image) {
				err["image"] = append(err["image"], fmt.Sprintf(fileContent, strings.Join(contentType(image), " or ")))
			}
		}
	}

	if req.Name == "" {
		err["name"] = append(err["name"], required)
	}
	name := maxMinString(req.Name, 3, 32)
	if name != "" {
		err["name"] = append(err["name"], name)
	}

	if req.Description != "" {
		description := maxMinString(req.Description, 3, 50)
		if description != "" {
			err["description"] = append(err["description"], description)
		}
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}
	return nil
}
