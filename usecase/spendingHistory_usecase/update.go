package spendingHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (s *SpendingHistoryUsecaseImpl) Update(ctx context.Context, req *usecase.RequestUpdateSpendingHistory) (*usecase.ResponseSpendingHistory, error) {
	err := s.validatePaymentAndSpendingTypeID(ctx, &usecase.ValidatePaymentIDAndSpendingTypeID{
		ProfileID:       req.ProfileID,
		SpendingTypeID:  req.SpendingTypeID,
		PaymentMethodID: req.PaymentMethodID,
	})
	if err != nil {
		return nil, err
	}

	balance, err := s.balanceUsecase.GetOrCreateByProfileID(ctx, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, usecase.SpendingHistoryNotFound
	}

	err = s.spendingHistoryRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		_, err = s.balanceUsecase.Update(ctx, &usecase.RequestUpdateBalance{
			ID:             balance.ID,
			ProfileID:      req.ProfileID,
			AmountSpending: balance.TotalSpendingAmount + req.SpendingAmount - spendingHistoryJoin.SpendingAmount,
			AmountIncome:   balance.TotalIncomeAmount,
			AmountBalance:  balance.Balance - req.SpendingAmount + spendingHistoryJoin.SpendingAmount,
		})
		if err != nil {
			return err
		}

		beforeBalance := balance.Balance + spendingHistoryJoin.SpendingAmount
		spendingHistory := req.ToModel(beforeBalance)

		err = s.spendingHistoryRepo.Update(ctx, spendingHistory)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	spendingHistoryJoin, err = s.spendingHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, usecase.SpendingHistoryNotFound
	}

	resp := &usecase.ResponseSpendingHistory{
		ID:                      spendingHistoryJoin.ID,
		ProfileID:               spendingHistoryJoin.ProfileID,
		SpendingTypeID:          spendingHistoryJoin.SpendingTypeID,
		SpendingTypeTitle:       spendingHistoryJoin.SpendingTypeTitle,
		PaymentMethodID:         repository.GetNullString(spendingHistoryJoin.PaymentMethodID),
		PaymentMethodName:       repository.GetNullString(spendingHistoryJoin.PaymentMethodName),
		PaymentName:             repository.GetNullString(spendingHistoryJoin.PaymentName),
		BeforeBalance:           spendingHistoryJoin.BeforeBalance,
		SpendingAmount:          spendingHistoryJoin.SpendingAmount,
		FormatSpendingAmount:    usecase.FormatRupiah(spendingHistoryJoin.SpendingAmount),
		AfterBalance:            spendingHistoryJoin.AfterBalance,
		Description:             spendingHistoryJoin.Description,
		TimeSpendingHistory:     spendingHistoryJoin.TimeSpendingHistory,
		ShowTimeSpendingHistory: spendingHistoryJoin.ShowTimeSpendingHistory,
	}

	return resp, nil
}
