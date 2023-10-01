package domain

import (
	"context"
	"database/sql"
)

type IncomeTypeRepository interface {
	Create(ctx context.Context, income *IncomeType) error
	Update(ctx context.Context, income *IncomeType) error
	Delete(ctx context.Context, id string, profileID string) error
	CheckByNameAndProfileID(ctx context.Context, profileID string, name string) (bool, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*IncomeType, error)
	GetAllByProfileID(ctx context.Context, profileID string) (*[]IncomeType, error)
	UnitOfWorkRepository
}

type IncomeTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateIncomeType) (*ResponseIncomeType, error)
	Update(ctx context.Context, req *RequestUpdateIncomeType) (*ResponseIncomeType, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseIncomeType, error)
	GetAllByProfileID(ctx context.Context, profileID string) (*[]ResponseIncomeType, error)
}

// IncomeType MODEL
type IncomeType struct {
	ID          string
	ProfileID   string
	Name        string
	Description sql.NullString
	Icon        string
	IncomeType  string
	FixedIncome sql.NullBool
	Periode     sql.NullString
	Amount      sql.NullInt64
	AuditInfo
}

// RequestCreateIncomeType REQUEST create schema for api
type RequestCreateIncomeType struct {
	ProfileID   string
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type RequestUpdateIncomeType struct {
	ID          string
	ProfileID   string
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// ResponseIncomeType response
type ResponseIncomeType struct {
	ID          string  `json:"id"`
	ProfileID   string  `json:"profile_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Icon        string  `json:"icon"`
}
