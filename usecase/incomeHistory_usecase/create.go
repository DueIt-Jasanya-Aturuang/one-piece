package incomeHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (i *IncomeHistoryUsecaseImpl) Create(ctx context.Context, req *usecase.RequestCreateIncomeHistory) (*usecase.ResponseIncomeHistory, error) {
	err := i.validateIncomeTypeAndPaymend(ctx, &usecase.ValidatePaymentIDAndIncomeTypeID{
		ProfileID:       req.ProfileID,
		IncomeTypeID:    req.IncomeTypeID,
		PaymentMethodID: req.PaymentMethodID,
	})
	if err != nil {
		return nil, err
	}

	balance, err := i.balanceUsecase.GetOrCreateByProfileID(ctx, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	var id string

	err = i.incomeHistoryRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		incomeHistory := req.ToModel()
		err = i.incomeHistoryRepo.Create(ctx, incomeHistory)
		if err != nil {
			return err
		}
		id = incomeHistory.ID

		balanceUpdate := &usecase.RequestUpdateBalance{
			ID:             balance.ID,
			ProfileID:      req.ProfileID,
			AmountSpending: balance.TotalSpendingAmount,
			AmountIncome:   balance.TotalIncomeAmount + req.IncomeAmount,
			AmountBalance:  balance.Balance + req.IncomeAmount,
		}
		balance, err = i.balanceUsecase.Update(ctx, balanceUpdate)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	incomeHistoryJoin, err := i.incomeHistoryRepo.GetByIDAndProfileID(ctx, id, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, usecase.IncomeHistoryNotFound
	}

	resp := &usecase.ResponseIncomeHistory{
		ID:                    incomeHistoryJoin.ID,
		ProfileID:             incomeHistoryJoin.ProfileID,
		IncomeTypeID:          incomeHistoryJoin.IncomeTypeID,
		IncomeTypeTitle:       incomeHistoryJoin.IncomeTypeTitle,
		PaymentMethodID:       repository.GetNullString(incomeHistoryJoin.PaymentMethodID),
		PaymentMethodName:     repository.GetNullString(incomeHistoryJoin.PaymentMethodName),
		PaymentName:           repository.GetNullString(incomeHistoryJoin.PaymentName),
		IncomeAmount:          incomeHistoryJoin.IncomeAmount,
		FormatIncomeAmount:    usecase.FormatRupiah(incomeHistoryJoin.IncomeAmount),
		Description:           incomeHistoryJoin.Description,
		TimeIncomeHistory:     incomeHistoryJoin.TimeIncomeHistory,
		ShowTimeIncomeHistory: incomeHistoryJoin.ShowTimeIncomeHistory,
	}

	return resp, nil
}
