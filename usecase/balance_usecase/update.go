package balance_usecase

import (
	"context"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (b *BalanceUsecaseImpl) Update(ctx context.Context, req *usecase.RequestUpdateBalance) (*usecase.ResponseBalance, error) {
	balance := &repository.Balance{
		ID:                  req.ID,
		ProfileID:           req.ProfileID,
		TotalIncomeAmount:   req.AmountIncome,
		TotalSpendingAmount: req.AmountSpending,
		Balance:             req.AmountBalance,
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: repository.NewNullString(req.ProfileID),
		},
	}

	err := b.balanceRepo.UpdateByProfileID(ctx, balance)
	if err != nil {
		return nil, err
	}

	resp := &usecase.ResponseBalance{
		ID:                        balance.ID,
		ProfileID:                 balance.ProfileID,
		TotalIncomeAmount:         balance.TotalIncomeAmount,
		TotalIncomeAmountFormat:   usecase.FormatRupiah(balance.TotalIncomeAmount),
		TotalSpendingAmount:       balance.TotalSpendingAmount,
		TotalSpendingAmountFormat: usecase.FormatRupiah(balance.TotalSpendingAmount),
		Balance:                   balance.Balance,
		BalanceFormat:             usecase.FormatRupiah(balance.Balance),
	}

	return resp, nil
}
