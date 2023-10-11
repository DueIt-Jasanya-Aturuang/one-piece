package balance_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (b *BalanceUsecaseImpl) GetByProfileID(ctx context.Context, profileID string) (*usecase.ResponseBalance, error) {
	balance, err := b.balanceRepo.GetByProfileID(ctx, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msg("balande tidak terseida")
			return nil, usecase.BalanceNotExist
		}
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
