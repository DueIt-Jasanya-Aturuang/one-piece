package domain

import (
	"context"
)

type Balance struct {
	ID                  string
	ProfileID           string
	TotalIncomeAmount   int
	TotalSpendingAmount int
	Balance             int
	AuditInfo
}

type BalanceRepository interface {
	Create(ctx context.Context, balance *Balance) error
	UpdateByProfileID(ctx context.Context, balance *Balance) error
	GetByProfileID(ctx context.Context, profileID string) (*Balance, error)
	UnitOfWorkRepository
}
