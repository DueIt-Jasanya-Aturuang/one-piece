package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	helper2 "github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

func CreateBalanceToModel(profileID string) *domain.Balance {
	id := uuid.NewV4().String()
	balance := &domain.Balance{
		ID:                  id,
		ProfileID:           profileID,
		TotalIncomeAmount:   0,
		TotalSpendingAmount: 0,
		Balance:             0,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: profileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
	return balance
}

func UpdateBalanceSpendingToModel(id string, profileID string, amount int, amountBalance int) *domain.Balance {
	balance := &domain.Balance{
		ID:                  id,
		ProfileID:           profileID,
		TotalSpendingAmount: amount,
		Balance:             amountBalance,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper2.NewNullString(profileID),
		},
	}
	return balance
}

func UpdateBalanceIncomeToModel(id string, profileID string, amount int, amountBalance int) *domain.Balance {
	balance := &domain.Balance{
		ID:                id,
		ProfileID:         profileID,
		TotalIncomeAmount: amount,
		Balance:           amountBalance,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper2.NewNullString(profileID),
		},
	}
	return balance
}

func BalanceModelToResponse(b *domain.Balance) *domain.ResponseBalance {
	return &domain.ResponseBalance{
		ID:                        b.ID,
		ProfileID:                 b.ProfileID,
		TotalIncomeAmount:         b.TotalIncomeAmount,
		TotalIncomeAmountFormat:   helper2.FormatRupiah(b.TotalIncomeAmount),
		TotalSpendingAmount:       b.TotalSpendingAmount,
		TotalSpendingAmountFormat: helper2.FormatRupiah(b.TotalSpendingAmount),
		Balance:                   b.Balance,
		BalanceFormat:             helper2.FormatRupiah(b.Balance),
	}
}
