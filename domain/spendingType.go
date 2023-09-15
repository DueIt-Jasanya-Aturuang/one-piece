package domain

import (
	"context"
	"time"
)

// SpendingType spending type entity
type SpendingType struct {
	ID           string
	ProfileID    string
	Title        string
	MaximumLimit int
	Icon         string
	AuditInfo
}

// SpendingTypeJoin join table
type SpendingTypeJoin struct {
	ID           string
	ProfileID    string
	Title        string
	MaximumLimit int
	Icon         string
	Used         int
	AuditInfo
}

// RequestCreateSpendingType request create spending type
type RequestCreateSpendingType struct {
	ProfileID    string `json:"profile_id"`
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
	Icon         string `json:"icon"`
}

// RequestUpdateSpendingType request update spending type
type RequestUpdateSpendingType struct {
	ID           string
	ProfileID    string
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
	Icon         string `json:"icon"`
}

type RequestGetAllSpendingType struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
}

// ResponseSpendingType response spending type
type ResponseSpendingType struct {
	ID                 string `json:"id"`
	ProfileID          string `json:"profile_id"`
	Title              string `json:"title"`
	MaximumLimit       int    `json:"maximum_limit"`
	FormatMaximumLimit string `json:"format_maximum_limit"`
	Icon               string `json:"icon"`
	Used               int    `json:"used,omiempty"`
	PersentaseUsed     string `json:"persentase_used,omitempty"`
}

// SpendingTypeRepository spending history repository interface
type SpendingTypeRepository interface {
	Create(ctx context.Context, spendingType *SpendingType) error
	Update(ctx context.Context, spendingType *SpendingType) error
	Delete(ctx context.Context, id string, profileID string) error
	CheckData(ctx context.Context, profileID string) (bool, error)
	CheckByTitleAndProfileID(ctx context.Context, profileID string, title string) (bool, error)
	GetDefault(ctx context.Context) (*[]SpendingType, error)
	GetByID(ctx context.Context, id string) (*SpendingType, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*SpendingType, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllSpendingType) (*[]SpendingTypeJoin, error)
	UnitOfWorkRepository
}

// SpendingTypeUsecase spending history usecase interface
type SpendingTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingType) (*ResponseSpendingType, error)
	Update(ctx context.Context, req *RequestUpdateSpendingType) (*ResponseSpendingType, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseSpendingType, error)
	GetAllByProfileID(ctx context.Context, profileID string, periode int) (*[]ResponseSpendingType, error)
}
