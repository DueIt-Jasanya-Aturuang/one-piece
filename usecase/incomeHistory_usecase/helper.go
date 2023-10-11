package incomeHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (i *IncomeHistoryUsecaseImpl) validateIncomeTypeAndPaymend(ctx context.Context, req *usecase.ValidatePaymentIDAndIncomeTypeID) error {
	_, err := i.incomeTypeRepo.GetByIDAndProfileID(ctx, req.IncomeTypeID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return usecase.InvalidIncomeTypeID
	}

	if req.PaymentMethodID != "" {
		_, err = i.paymentRepo.GetByIDAndProfileID(ctx, req.PaymentMethodID, req.ProfileID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			return usecase.InvalidPaymentMethodID
		}
	}

	return nil

}
