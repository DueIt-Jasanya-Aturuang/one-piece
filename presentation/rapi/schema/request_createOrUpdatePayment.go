package schema

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type RequestCreateOrUpdatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
}

func (req *RequestCreateOrUpdatePayment) ValidateCreate() error {
	err := map[string][]string{}

	if req.Image == nil {
		err["image"] = append(err["image"], util.Required)
	}
	if req.Image != nil && req.Image.Size > 0 {
		if req.Image.Size > 2097152 {
			err["image"] = append(err["image"], fmt.Sprintf(util.FileSize, 2048, 2))
		}
		if !util.CheckContentType(req.Image.Header.Get("Content-Type"), util.Image) {
			err["image"] = append(err["image"], fmt.Sprintf(util.FileContent, strings.Join(util.ContentType(util.Image), " or ")))
		}
	}

	if req.Name == "" {
		err["name"] = append(err["name"], util.Required)
	}
	name := util.MaxMinString(req.Name, 3, 32)
	if name != "" {
		err["name"] = append(err["name"], name)
	}

	if req.Description != "" {
		description := util.MaxMinString(req.Description, 3, 50)
		if description != "" {
			err["description"] = append(err["description"], description)
		}
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}
	return nil
}

func (req *RequestCreateOrUpdatePayment) ValidateUpdate() error {
	err := map[string][]string{}

	if req.Image != nil {
		if req.Image != nil && req.Image.Size > 0 {
			if req.Image.Size > 2097152 {
				err["image"] = append(err["image"], fmt.Sprintf(util.FileSize, 2048, 2))
			}
			if !util.CheckContentType(req.Image.Header.Get("Content-Type"), util.Image) {
				err["image"] = append(err["image"], fmt.Sprintf(util.FileContent, strings.Join(util.ContentType(util.Image), " or ")))
			}
		}
	}

	if req.Name == "" {
		err["name"] = append(err["name"], util.Required)
	}
	name := util.MaxMinString(req.Name, 3, 32)
	if name != "" {
		err["name"] = append(err["name"], name)
	}

	if req.Description != "" {
		description := util.MaxMinString(req.Description, 3, 50)
		if description != "" {
			err["description"] = append(err["description"], description)
		}
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}
	return nil
}
