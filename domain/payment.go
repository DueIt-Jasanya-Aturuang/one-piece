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

//counterfeiter:generate -o ./mocks . PaymentRepository
type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *Payment) error
	UpdatePayment(ctx context.Context, payment *Payment) error
	GetAllPayment(ctx context.Context) (*[]Payment, error)
	GetPaymentByID(ctx context.Context, id string) (*Payment, error)
	GetPaymentByName(ctx context.Context, name string) (*Payment, error)
	UnitOfWorkRepository
}

//counterfeiter:generate -o ./mocks . PaymentUsecase
type PaymentUsecase interface {
	CreatePayment(ctx context.Context, req *RequestCreatePayment) (*ResponsePayment, error)
	UpdatePayment(ctx context.Context, req *RequestUpdatePayment) (*ResponsePayment, error)
	GetAllPayment(ctx context.Context) (*[]ResponsePayment, error)
}

type RequestCreatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
}

type RequestUpdatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	ID          string
}

type ResponsePayment struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
