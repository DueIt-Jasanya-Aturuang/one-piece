package domain

import (
	"context"
	"database/sql"
	"mime/multipart"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
	Delete(ctx context.Context, id string, profileID string) error
	CheckData(ctx context.Context, profileID string) (bool, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllPaymentPaginate) (*[]Payment, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*Payment, error)
	GetByNameAndProfileID(ctx context.Context, name string, profileID string) (*Payment, error)
	GetDefault(ctx context.Context) (*[]Payment, error)
	UnitOfWorkRepository
}

type PaymentUsecase interface {
	Create(ctx context.Context, req *RequestCreatePayment) (*ResponsePayment, error)
	Update(ctx context.Context, req *RequestUpdatePayment) (*ResponsePayment, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllPaymentPaginate) (*[]ResponsePayment, string, error)
	Delete(ctx context.Context, id string, profileID string) error
}

type RequestGetAllPaymentPaginate struct {
	ProfileID string
	ID        string
	Operation string
	Order     string
}

// Payment entity payment
type Payment struct {
	ID          string
	ProfileID   string
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
	ProfileID   string
}

// RequestUpdatePayment update payment request
type RequestUpdatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	ProfileID   string
	ID          string
}

// ResponsePayment response payment
type ResponsePayment struct {
	ID          string  `json:"id"`
	ProfileID   string  `json:"profile_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
}
