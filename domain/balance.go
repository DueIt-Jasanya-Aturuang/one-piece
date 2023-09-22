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

type ResponseBalance struct {
	ID                        string `json:"id"`
	ProfileID                 string `json:"profile_id"`
	TotalIncomeAmount         int    `json:"total_income_amount"`
	TotalIncomeAmountFormat   string `json:"total_income_amount_format"`
	TotalSpendingAmount       int    `json:"total_spending_amount"`
	TotalSpendingAmountFormat string `json:"total_spending_amount_format"`
	Balance                   int    `json:"balance"`
	BalanceFormat             string `json:"balance_format"`
}

type BalanceRepository interface {
	Create(ctx context.Context, balance *Balance) error
	UpdateByProfileID(ctx context.Context, balance *Balance) error
	GetByProfileID(ctx context.Context, profileID string) (*Balance, error)
	UnitOfWorkRepository
}

type BalanceUsecase interface {
	GetByProfileID(ctx context.Context, profileID string) (*ResponseBalance, error)
}
