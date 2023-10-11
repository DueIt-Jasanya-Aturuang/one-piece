package incomeHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (i *IncomeHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	incomeHistoryJoin, err := i.incomeHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return usecase.IncomeHistoryNotFound
		}
		return err
	}

	balance, err := i.balanceUsecase.GetOrCreateByProfileID(ctx, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return usecase.ProfileIDNotFound
		}
		return err
	}

	err = i.incomeHistoryRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		_, err = i.balanceUsecase.Update(ctx, &usecase.RequestUpdateBalance{
			ID:             balance.ID,
			ProfileID:      profileID,
			AmountSpending: balance.TotalSpendingAmount,
			AmountIncome:   balance.TotalIncomeAmount - incomeHistoryJoin.IncomeAmount,
			AmountBalance:  balance.Balance - incomeHistoryJoin.IncomeAmount,
		})
		if err != nil {
			return err
		}

		err = i.incomeHistoryRepo.Delete(ctx, id, profileID)
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
