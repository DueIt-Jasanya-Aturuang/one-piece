package repository

import (
	"context"
	"database/sql"
	"time"
)

type SpendingHistoryRepository interface {
	Create(ctx context.Context, spendingHistory *SpendingHistory) error
	Update(ctx context.Context, spendingHistory *SpendingHistory) error
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *GetAllSpendingHistoryByFilterWithISD) (*[]SpendingHistoryJoinTable, error)
	GetTotalAmountByPeriode(ctx context.Context, req *GetTotalSpendingHistoryByPeriode) (int, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*SpendingHistoryJoinTable, error)
	UnitOfWorkRepository
}

type SpendingHistory struct {
	ID                      string
	ProfileID               string
	SpendingTypeID          string
	PaymentMethodID         sql.NullString
	PaymentName             sql.NullString
	BeforeBalance           int
	SpendingAmount          int
	AfterBalance            int
	Description             string
	TimeSpendingHistory     time.Time
	ShowTimeSpendingHistory string
	AuditInfo
}

type SpendingHistoryJoinTable struct {
	ID                      string
	ProfileID               string
	SpendingTypeID          string
	SpendingTypeTitle       string
	PaymentMethodID         sql.NullString
	PaymentMethodName       sql.NullString
	PaymentName             sql.NullString
	BeforeBalance           int
	SpendingAmount          int
	AfterBalance            int
	Description             string
	TimeSpendingHistory     time.Time
	ShowTimeSpendingHistory string
	AuditInfo
}

type GetTotalSpendingHistoryByPeriode struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
}

type GetAllSpendingHistoryByFilterWithISD struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
	Type      string
	InfiniteScrollData
}
