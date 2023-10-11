package repository

import (
	"context"
	"database/sql"
	"time"
)

type IncomeHistoryRepository interface {
	Create(ctx context.Context, income *IncomeHistory) error
	Update(ctx context.Context, income *IncomeHistory) error
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *GetAllIncomeHistoryByTimeFilterWithISD) (*[]IncomeHistoryJoinTable, error)
	GetTotalAmountByPeriode(ctx context.Context, req *GetTotalIncomeHistoryByPeriode) (int, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*IncomeHistoryJoinTable, error)
	UnitOfWorkRepository
}

type IncomeHistory struct {
	ID                    string
	ProfileID             string
	IncomeTypeID          string
	PaymentMethodID       sql.NullString
	PaymentName           sql.NullString
	IncomeAmount          int
	Description           string
	TimeIncomeHistory     time.Time
	ShowTimeIncomeHistory string
	AuditInfo
}

type IncomeHistoryJoinTable struct {
	ID                    string
	ProfileID             string
	IncomeTypeID          string
	IncomeTypeTitle       string
	PaymentMethodID       sql.NullString
	PaymentMethodName     sql.NullString
	PaymentName           sql.NullString
	IncomeAmount          int
	Description           string
	TimeIncomeHistory     time.Time
	ShowTimeIncomeHistory string
	AuditInfo
}

type GetAllIncomeHistoryByTimeFilterWithISD struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
	InfiniteScrollData
}

type GetTotalIncomeHistoryByPeriode struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
}
