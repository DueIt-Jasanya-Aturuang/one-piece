package repository

import (
	"context"
	"database/sql"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
	Delete(ctx context.Context, id string, profileID string) error
	CheckData(ctx context.Context, profileID string) (bool, error)
	GetAllByProfileID(ctx context.Context, req *GetAllPaymentWithISD) (*[]Payment, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*Payment, error)
	GetByNameAndProfileID(ctx context.Context, name string, profileID string) (*Payment, error)
	GetDefault(ctx context.Context) (*[]Payment, error)
	UnitOfWorkRepository
}

type Payment struct {
	ID          string
	ProfileID   string
	Name        string
	Description sql.NullString
	Image       string
	AuditInfo
}

type GetAllPaymentWithISD struct {
	ProfileID string
	InfiniteScrollData
}
