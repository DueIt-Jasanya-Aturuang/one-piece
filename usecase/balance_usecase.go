package usecase

import (
	"context"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type BalanceUsecase interface {
	Update(ctx context.Context, req *RequestUpdateBalance) (*ResponseBalance, error)
	GetByProfileID(ctx context.Context, profileID string) (*ResponseBalance, error)
	GetOrCreateByProfileID(ctx context.Context, profileID string) (*ResponseBalance, error)
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

type RequestUpdateBalance struct {
	ID             string
	ProfileID      string
	AmountSpending int
	AmountIncome   int
	AmountBalance  int
}

func (u *RequestUpdateBalance) UpdateBalanceToModel() *repository.Balance {
	balance := &repository.Balance{
		ID:                  u.ID,
		ProfileID:           u.ProfileID,
		TotalSpendingAmount: u.AmountSpending,
		TotalIncomeAmount:   u.AmountIncome,
		Balance:             u.AmountBalance,
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: repository.NewNullString(u.ProfileID),
		},
	}
	return balance
}

func CreateBalanceToModel(profileID string) *repository.Balance {
	id := util.NewUUID()
	balance := &repository.Balance{
		ID:                  id,
		ProfileID:           profileID,
		TotalIncomeAmount:   0,
		TotalSpendingAmount: 0,
		Balance:             0,
		AuditInfo: repository.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: profileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
	return balance
}
