package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helper"
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

func UpdateBalanceToModel(id string, profileID string, amount int, amountBalance int) *domain.Balance {
	balance := &domain.Balance{
		ID:                  id,
		ProfileID:           profileID,
		TotalSpendingAmount: amount,
		Balance:             amountBalance,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString(profileID),
		},
	}
	return balance
}
