package spendingHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (s *SpendingHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return usecase.SpendingHistoryNotFound
		}
		return err
	}

	balance, err := s.balanceUsecase.GetOrCreateByProfileID(ctx, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return usecase.ProfileIDNotFound
		}
		return err
	}

	err = s.spendingHistoryRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		_, err = s.balanceUsecase.Update(ctx, &usecase.RequestUpdateBalance{
			ID:             balance.ID,
			ProfileID:      profileID,
			AmountSpending: balance.TotalSpendingAmount - spendingHistoryJoin.SpendingAmount,
			AmountIncome:   balance.TotalIncomeAmount,
			AmountBalance:  balance.Balance + spendingHistoryJoin.SpendingAmount,
		})
		if err != nil {
			return err
		}

		err = s.spendingHistoryRepo.Delete(ctx, id, profileID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
