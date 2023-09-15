package domain

import (
	"context"
	"database/sql"
	"mime/multipart"
)

// Payment entity payment
type Payment struct {
	ID          string
	Name        string
	Description sql.NullString
	Image       string
	AuditInfo
}

// RequestCreatePayment create payment request
type RequestCreatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
}

// RequestUpdatePayment update payment request
type RequestUpdatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	ID          string
}

// ResponsePayment response payment
type ResponsePayment struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
}

// PaymentRepository payment repository interface
//
//counterfeiter:generate -o ./mocks . PaymentRepository
type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
	GetAll(ctx context.Context) (*[]Payment, error)
	GetByID(ctx context.Context, id string) (*Payment, error)
	GetByName(ctx context.Context, name string) (*Payment, error)
	UnitOfWorkRepository
}

// PaymentUsecase payment usecase interface
//
//counterfeiter:generate -o ./mocks . PaymentUsecase
type PaymentUsecase interface {
	Create(ctx context.Context, req *RequestCreatePayment) (*ResponsePayment, error)
	Update(ctx context.Context, req *RequestUpdatePayment) (*ResponsePayment, error)
	GetAll(ctx context.Context) (*[]ResponsePayment, error)
}
