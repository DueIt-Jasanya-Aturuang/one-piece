package domain

import (
	"context"
	"database/sql"
	"mime/multipart"
)

type Payment struct {
	ID          string
	Name        string
	Description sql.NullString
	Image       string
	CreatedAt   int64
	CreatedBy   string
	UpdatedAt   int64
	UpdatedBy   sql.NullString
	DeletedAt   sql.NullInt64
	DeletedBy   sql.NullString
}

type PaymentUsecase interface {
	CreatePayment(ctx context.Context, req *RequestCreatePayment) (*ResponsePayment, error)
	UpdatePayment(ctx context.Context, req *RequestUpdatePayment, id string) (*ResponsePayment, error)
	GetPaymentByName(ctx context.Context, name string) (*ResponsePayment, error)
	DeletePayment(ctx context.Context, id string) (bool, error)
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *Payment) error
	UpdatePayment(ctx context.Context, payment *Payment) error
	GetPaymentByID(ctx context.Context, id string) (*Payment, error)
	GetPaymentByName(ctx context.Context, name string) (*Payment, error)
	UnitOfWorkRepository
}

type RequestCreatePayment struct {
	Id          string
	Name        string                `json:"name" form:"name" validation:"required,min=3,max=32"`
	Description string                `json:"description" form:"description"`
	Image       *multipart.FileHeader `json:"image" form:"image" validation:"required"`
}

type RequestUpdatePayment struct {
	Name        string                `json:"name" form:"name" validation:"required,min=3,max=32"`
	Description string                `json:"description" form:"description"`
	Image       *multipart.FileHeader `json:"image" form:"image" validation:"required"`
}

type ResponsePayment struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
}
