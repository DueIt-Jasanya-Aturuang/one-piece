package dto

import "mime/multipart"

type PaymentCreateRequest struct {
	Id          string
	Name        string                `json:"name" form:"name" validate:"required,min=3,max=32"`
	Description string                `json:"description" form:"description"`
	Image       *multipart.FileHeader `json:"image" form:"image" swaggerignore:"true" validate:"required"`
}

type PaymentUpdateRequest struct {
	Name        string                `json:"name" form:"name" validate:"required,min=3,max=32"`
	Description string                `json:"description" form:"description"`
	Image       *multipart.FileHeader `json:"image" form:"image" swaggerignore:"true" validate:"required"`
}
