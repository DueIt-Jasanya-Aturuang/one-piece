package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

type IncomeHistoryUsecaseImpl struct {
	incomeTypeRepo    domain.IncomeTypeRepository
	incomeHistoryRepo domain.IncomeHistoryRepository
	paymentRepo       domain.PaymentRepository
	balanceRepo       domain.BalanceRepository
}

func NewIncomeHistoryUsecaseImpl(
	incomeTypeRepo domain.IncomeTypeRepository,
	incomeHistoryRepo domain.IncomeHistoryRepository,
	paymentRepo domain.PaymentRepository,
	balanceRepo domain.BalanceRepository,
) domain.IncomeHistoryUsecase {
	return &IncomeHistoryUsecaseImpl{
		incomeTypeRepo:    incomeTypeRepo,
		incomeHistoryRepo: incomeHistoryRepo,
		paymentRepo:       paymentRepo,
		balanceRepo:       balanceRepo,
	}
}

func (i *IncomeHistoryUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateIncomeHistory) (*domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer i.incomeHistoryRepo.CloseConn()

	err = i.validateIncomeTypeAndPaymend(ctx, &domain.ValidatePaymentAndIncomeTypeID{
		ProfileID:       req.ProfileID,
		IncomeTypeID:    req.IncomeTypeID,
		PaymentMethodID: req.PaymentMethodID,
	})
	if err != nil {
		return nil, err
	}

	balance, err := i.balanceRepo.GetByProfileID(ctx, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	var id string

	err = i.incomeHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		if balance == nil {
			balance = converter.CreateBalanceToModel(req.ProfileID)
			err = i.balanceRepo.Create(ctx, balance)
			if err != nil {
				return err
			}
		}

		incomeHistory := converter.CreateIncomeHistoryToModel(req)
		err = i.incomeHistoryRepo.Create(ctx, incomeHistory)
		if err != nil {
			return err
		}
		id = incomeHistory.ID

		amountSpending := balance.TotalIncomeAmount + req.IncomeAmount
		amountBalance := balance.Balance + req.IncomeAmount
		balance = converter.UpdateBalanceIncomeToModel(balance.ID, req.ProfileID, amountSpending, amountBalance)
		err = i.balanceRepo.UpdateByProfileID(ctx, balance)
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
		return nil, IncomeHistoryNotFound
	}

	resp := converter.IncomeHistoryJoinModelToResponse(incomeHistoryJoin)

	return resp, nil
}

func (i *IncomeHistoryUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateIncomeHistory) (*domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer i.incomeHistoryRepo.CloseConn()

	err = i.validateIncomeTypeAndPaymend(ctx, &domain.ValidatePaymentAndIncomeTypeID{
		ProfileID:       req.ProfileID,
		IncomeTypeID:    req.IncomeTypeID,
		PaymentMethodID: req.PaymentMethodID,
	})
	if err != nil {
		return nil, err
	}

	_, err = i.incomeHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, IncomeHistoryNotFound
	}

	balance, err := i.balanceRepo.GetByProfileID(ctx, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	err = i.incomeHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		if balance == nil {
			balance = converter.CreateBalanceToModel(req.ProfileID)
			err = i.balanceRepo.Create(ctx, balance)
			if err != nil {
				return err
			}
		}

		incomeHistory := converter.UpdateIncomeHistoryToModel(req)
		err = i.incomeHistoryRepo.Update(ctx, incomeHistory)
		if err != nil {
			return err
		}

		amountSpending := balance.TotalIncomeAmount + req.IncomeAmount
		amountBalance := balance.Balance + req.IncomeAmount
		balance = converter.UpdateBalanceIncomeToModel(balance.ID, req.ProfileID, amountSpending, amountBalance)
		err = i.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	incomeHistoryJoin, err := i.incomeHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, IncomeHistoryNotFound
	}

	resp := converter.IncomeHistoryJoinModelToResponse(incomeHistoryJoin)

	return resp, nil
}

func (i *IncomeHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return err
	}
	defer i.incomeHistoryRepo.CloseConn()

	incomeHistoryJoin, err := i.incomeHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return IncomeHistoryNotFound
		}
		return err
	}

	balance, err := i.balanceRepo.GetByProfileID(ctx, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ProfileIDNotFound
		}
		return err
	}

	err = i.incomeHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		incomeAmount := balance.TotalIncomeAmount - incomeHistoryJoin.IncomeAmount
		balanceAmount := balance.Balance - incomeHistoryJoin.IncomeAmount
		balance = converter.UpdateBalanceIncomeToModel(balance.ID, profileID, incomeAmount, balanceAmount)

		err = i.balanceRepo.UpdateByProfileID(ctx, balance)
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

func (i *IncomeHistoryUsecaseImpl) GetAllByTimeAndProfileID(ctx context.Context, req *domain.GetFilteredDataIncomeHistory) (*[]domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer i.incomeHistoryRepo.CloseConn()

	var startTime, endTime time.Time
	if req.Type != "" {
		startTime, endTime, _ = helper.TimeDateByTypeFilter(req.Type)
	} else {
		startTime, endTime, err = helper.FormatDate(req.StartTime, req.EndTime)
		if err != nil {
			return nil, InvalidTimestamp
		}
	}

	incomeHistories, err := i.incomeHistoryRepo.GetAllByTimeAndProfileID(ctx, &domain.GetIncomeHistoryByTimeAndProfileID{
		ProfileID: req.ProfileID,
		StartTime: startTime,
		EndTime:   endTime,
	})
	if err != nil {
		return nil, err
	}

	var incomeHistoryJoinResponses []domain.ResponseIncomeHistory

	for _, incomeHistory := range *incomeHistories {
		incomeHistoryJoinResponse := converter.GetAllIncomeHistoryJoinModelToResponse(incomeHistory)
		incomeHistoryJoinResponses = append(incomeHistoryJoinResponses, incomeHistoryJoinResponse)
	}

	return &incomeHistoryJoinResponses, nil
}

func (i *IncomeHistoryUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer i.incomeHistoryRepo.CloseConn()

	incomeHistory, err := i.incomeHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, IncomeHistoryNotFound
		}
		return nil, err
	}

	incomeHistoryJoinResponse := converter.IncomeHistoryJoinModelToResponse(incomeHistory)

	return incomeHistoryJoinResponse, nil
}

func (i *IncomeHistoryUsecaseImpl) validateIncomeTypeAndPaymend(ctx context.Context, req *domain.ValidatePaymentAndIncomeTypeID) error {
	_, err := i.incomeTypeRepo.GetByIDAndProfileID(ctx, req.IncomeTypeID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return InvalidIncomeTypeID
	}

	if req.PaymentMethodID != "" {
		_, err = i.paymentRepo.GetByIDAndProfileID(ctx, req.PaymentMethodID, req.ProfileID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			return InvalidPaymentMethodID
		}
	}

	return nil

}
