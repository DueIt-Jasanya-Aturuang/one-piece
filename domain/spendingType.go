package domain

import (
	"context"
	"database/sql"
)

// SpendingType spending type entity
type SpendingType struct {
	ID           string
	ProfileID    string
	Title        string
	MaximumLimit int
	CreatedAt    int64
	CreatedBy    string
	UpdatedAt    int64
	UpdatedBy    sql.NullString
	DeletedAt    sql.NullInt64
	DeletedBy    sql.NullString
}

// RequestCreateSpendingType request create spending type
type RequestCreateSpendingType struct {
	ProfileID    string `json:"profile_id"`
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
}

// RequestUpdateSpendingType request update spending type
type RequestUpdateSpendingType struct {
	ID           string
	ProfileID    string
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
}

// ResponseSpendingType response spending type
type ResponseSpendingType struct {
	ID           string `json:"id"`
	ProfileID    string `json:"profile_id"`
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
}

// SpendingTypeRepository spending history repository interface
type SpendingTypeRepository interface {
	Create(ctx context.Context, spendingType *SpendingType) error
	Update(ctx context.Context, spendingType *SpendingType) error
	Delete(ctx context.Context, id string, profileID string) error
	GetByID(ctx context.Context, id string) (*SpendingType, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*SpendingType, error)
	GetAllByProfileID(ctx context.Context, profileID string) (*SpendingType, error)
	UnitOfWorkRepository
}

// SpendingTypeUsecase spending history usecase interface
type SpendingTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingType) (*ResponseSpendingType, error)
	Update(ctx context.Context, req *RequestUpdateSpendingType) (*ResponseSpendingType, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseSpendingType, error)
	GetAllByProfileID(ctx context.Context, profileID string) (*[]ResponseSpendingType, error)
}
