package spendingHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (s *SpendingHistoryUsecaseImpl) validatePaymentAndSpendingTypeID(ctx context.Context, req *usecase.ValidatePaymentIDAndSpendingTypeID) error {
	_, err := s.spendingTypeRepo.GetByIDAndProfileID(ctx, req.SpendingTypeID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return usecase.InvalidSpendingTypeID
	}

	if req.PaymentMethodID != "" {
		_, err = s.paymentRepo.GetByIDAndProfileID(ctx, req.PaymentMethodID, req.ProfileID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			return usecase.InvalidPaymentMethodID
		}
	}

	return nil
}
