package domain

import (
	"context"
	"time"
)

// SpendingTypeRepository spending history repository interface
type SpendingTypeRepository interface {
	Create(ctx context.Context, spendingType *SpendingType) error
	Update(ctx context.Context, spendingType *SpendingType) error
	Delete(ctx context.Context, id string, profileID string) error
	CheckData(ctx context.Context, profileID string) (bool, error)
	CheckByTitleAndProfileID(ctx context.Context, profileID string, title string) (bool, error)
	GetDefault(ctx context.Context) (*[]SpendingType, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*SpendingType, error)
	GetAllByTimeAndProfileID(ctx context.Context, req *RequestGetAllSpendingTypeByTime) (*[]SpendingTypeJoin, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllPaginate) (*[]SpendingType, error)
	UnitOfWorkRepository
}

// SpendingTypeUsecase spending history usecase interface
type SpendingTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingType) (*ResponseSpendingType, error)
	Update(ctx context.Context, req *RequestUpdateSpendingType) (*ResponseSpendingType, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseSpendingType, error)
	GetAllByPeriodeAndProfileID(ctx context.Context, req *RequestGetAllSpendingTypeByTime) (*ResponseAllSpendingType, string, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllPaginate) (*[]ResponseSpendingType, string, error)
}

type RequestGetAllSpendingTypeByTime struct {
	ProfileID string
	Periode   int
	StartTime time.Time
	EndTime   time.Time
	RequestGetAllPaginate
}

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
	ProfileID    string
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
	Icon         string `json:"icon"`
}

// RequestUpdateSpendingType request update spending type
type RequestUpdateSpendingType struct {
	ID           string // ID get in url parameter
	ProfileID    string
	Title        string `json:"title"`
	MaximumLimit int    `json:"maximum_limit"`
	Icon         string `json:"icon"`
}

// ResponseSpendingType response spending type
type ResponseSpendingType struct {
	ID                 string `json:"id"`
	ProfileID          string `json:"profile_id"`
	Title              string `json:"title"`
	MaximumLimit       int    `json:"maximum_limit"`
	FormatMaximumLimit string `json:"format_maximum_limit"`
	Icon               string `json:"icon"`
}

type ResponseSpendingTypeJoin struct {
	ID                 string `json:"id"`
	ProfileID          string `json:"profile_id"`
	Title              string `json:"title"`
	MaximumLimit       int    `json:"maximum_limit"`
	FormatMaximumLimit string `json:"format_maximum_limit"`
	Icon               string `json:"icon"`
	Used               int    `json:"used"`
	FormatUsed         string `json:"format_used"`
	PersentaseUsed     string `json:"persentase_used"`
}

// ResponseAllSpendingType response get all spending type per periode
type ResponseAllSpendingType struct {
	ResponseSpendingType *[]ResponseSpendingTypeJoin `json:"spending_type"`
	BudgetAmount         int                         `json:"budget_amount"`
	FormatBudgetAmount   string                      `json:"format_budget_amount"`
}
