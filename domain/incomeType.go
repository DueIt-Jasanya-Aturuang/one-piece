package domain

import (
	"context"
	"time"
)

type IncomeType struct {
	ID string
	AuditInfo
}

type RequestCreateIncomeType struct {
	Name string
}

type RequestUpdateIncomeType struct {
	ID string
}

type RequestGetAllIncomeType struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
	Type      string
}

type ResponseIncomeType struct {
	ID string
}

type IncomeTypeRepository interface {
	Create(ctx context.Context, income *IncomeType) error
	Update(ctx context.Context, income *IncomeType) error
	Delete(ctx context.Context, id string, profileID string) error
	CheckByTitleAndProfileID(ctx context.Context, profileID string, title string) (bool, error)
	GetByID(ctx context.Context, id string) (*IncomeType, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*IncomeType, error)
	GetAllByProfileID(ctx context.Context, profileID string) (*[]IncomeType, error)
	UnitOfWorkRepository
}

type IncomeTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateIncomeType) (*ResponseIncomeType, error)
	Update(ctx context.Context, req *RequestUpdateIncomeHistory) (*ResponseIncomeType, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseIncomeType, error)
	GetAllByProfileID(ctx context.Context, profileID string) (*[]ResponseIncomeType, error)
}
