package balance_usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (b *BalanceUsecaseImpl) GetOrCreateByProfileID(ctx context.Context, profileID string) (*usecase.ResponseBalance, error) {
	balance, err := b.balanceRepo.GetByProfileID(ctx, profileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	if balance == nil {
		balance = &repository.Balance{
			ID:                  util.NewUUID(),
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
		
		err = b.balanceRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
			err = b.balanceRepo.Create(ctx, balance)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
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
