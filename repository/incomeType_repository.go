package repository

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
	GetAllByProfileID(ctx context.Context, req *GetAllIncomeTypeWithISD) (*[]IncomeType, error)
	UnitOfWorkRepository
}

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

type GetAllIncomeTypeWithISD struct {
	ProfileID string
	InfiniteScrollData
}
