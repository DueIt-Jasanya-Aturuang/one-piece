package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

type SpendingHistoryUsecaseImpl struct {
	spendingHistoryRepo domain.SpendingHistoryRepository
	spendingTypeRepo    domain.SpendingTypeRepository
	balanceRepo         domain.BalanceRepository
	paymentRepo         domain.PaymentRepository
}

func NewSpendingHistoryUsecaseImpl(
	spendingHistoryRepo domain.SpendingHistoryRepository,
	spendingTypeRepo domain.SpendingTypeRepository,
	balanceRepo domain.BalanceRepository,
	paymentRepo domain.PaymentRepository,
) domain.SpendingHistoryUsecase {
	return &SpendingHistoryUsecaseImpl{
		spendingHistoryRepo: spendingHistoryRepo,
		spendingTypeRepo:    spendingTypeRepo,
		balanceRepo:         balanceRepo,
		paymentRepo:         paymentRepo,
	}
}

func (s *SpendingHistoryUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateSpendingHistory) (*domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer s.spendingHistoryRepo.CloseConn()

	err = s.validatePaymentAndSpendingTypeID(ctx, &domain.RequestValidatePaymentAndSpendingTypeID{
		ProfileID:       req.ProfileID,
		SpendingTypeID:  req.SpendingTypeID,
		PaymentMethodID: req.PaymentMethodID,
	})
	if err != nil {
		return nil, err
	}

	balance, err := s.balanceRepo.GetByProfileID(ctx, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	var id string

	err = s.spendingHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		if balance == nil {
			balance = converter.CreateBalanceToModel(req.ProfileID)
			err = s.balanceRepo.Create(ctx, balance)
			if err != nil {
				return err
			}
		}

		spendingHistory := converter.CreateSpendingHistoryToModel(req, balance.Balance)
		err = s.spendingHistoryRepo.Create(ctx, spendingHistory)
		if err != nil {
			return err
		}
		id = spendingHistory.ID

		amountSpending := balance.TotalSpendingAmount + req.SpendingAmount
		amountBalance := balance.Balance - req.SpendingAmount
		balance = converter.UpdateBalanceToModel(balance.ID, req.ProfileID, amountSpending, balance.TotalIncomeAmount, amountBalance)
		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, SpendingHistoryNotFound
	}

	resp := converter.SpendingHistoryJoinModelToResponse(spendingHistoryJoin)

	return resp, nil
}

func (s *SpendingHistoryUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateSpendingHistory) (*domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer s.spendingHistoryRepo.CloseConn()

	err = s.validatePaymentAndSpendingTypeID(ctx, &domain.RequestValidatePaymentAndSpendingTypeID{
		ProfileID:       req.ProfileID,
		SpendingTypeID:  req.SpendingTypeID,
		PaymentMethodID: req.PaymentMethodID,
	})
	if err != nil {
		return nil, err
	}

	balance, err := s.balanceRepo.GetByProfileID(ctx, req.ProfileID)
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
		return nil, SpendingHistoryNotFound
	}

	err = s.spendingHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		if balance == nil {
			balance = converter.CreateBalanceToModel(req.ProfileID)
			err = s.balanceRepo.Create(ctx, balance)
			if err != nil {
				return err
			}
		}

		beforeBalance := balance.Balance + spendingHistoryJoin.SpendingAmount
		amountSpending := balance.TotalSpendingAmount + req.SpendingAmount - spendingHistoryJoin.SpendingAmount
		amountBalance := balance.Balance - req.SpendingAmount + spendingHistoryJoin.SpendingAmount
		balance = converter.UpdateBalanceToModel(balance.ID, req.ProfileID, amountSpending, balance.TotalIncomeAmount, amountBalance)

		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return err
		}

		spendingHistory := converter.UpdateSpendingHistoryToModel(req, beforeBalance)
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
		return nil, SpendingHistoryNotFound
	}

	resp := converter.SpendingHistoryJoinModelToResponse(spendingHistoryJoin)

	return resp, nil
}

func (s *SpendingHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return err
	}
	defer s.spendingHistoryRepo.CloseConn()

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return SpendingHistoryNotFound
		}
		return err
	}

	balance, err := s.balanceRepo.GetByProfileID(ctx, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ProfileIDNotFound
		}
		return err
	}

	err = s.spendingHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		spendingAmount := balance.TotalSpendingAmount - spendingHistoryJoin.SpendingAmount
		balanceAmount := balance.Balance + spendingHistoryJoin.SpendingAmount
		balance = converter.UpdateBalanceToModel(balance.ID, profileID, spendingAmount, balance.TotalIncomeAmount, balanceAmount)

		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
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

func (s *SpendingHistoryUsecaseImpl) GetAllByTimeAndProfileID(ctx context.Context, req *domain.GetFilteredDataSpendingHistory) (*[]domain.ResponseSpendingHistory, string, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, "", err
	}
	defer s.spendingHistoryRepo.CloseConn()

	if req.Type != "" {
		req.StartTime, req.EndTime, _ = helper.TimeDateByTypeFilter(req.Type)
	} else {
		req.StartTime, req.EndTime, err = helper.FormatDate(req.StartTime, req.EndTime)
		if err != nil {
			return nil, "", InvalidTimestamp
		}
	}

	spendingHistories, err := s.spendingHistoryRepo.GetAllByTimeAndProfileID(ctx, req)
	if err != nil {
		return nil, "", err
	}

	var spendingHistoryJoinResponses []domain.ResponseSpendingHistory
	var cursor string
	for _, spendingHistory := range *spendingHistories {
		spendingHistoryJoinResponse := converter.GetAllSpendingHistoryJoinModelToResponse(spendingHistory)
		spendingHistoryJoinResponses = append(spendingHistoryJoinResponses, spendingHistoryJoinResponse)

		cursor = spendingHistory.ID
	}

	return &spendingHistoryJoinResponses, cursor, nil
}

func (s *SpendingHistoryUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer s.spendingHistoryRepo.CloseConn()

	spendingHistory, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, SpendingHistoryNotFound
		}
		return nil, err
	}

	spendingHistoryJoinResponse := converter.SpendingHistoryJoinModelToResponse(spendingHistory)

	return spendingHistoryJoinResponse, nil
}

func (s *SpendingHistoryUsecaseImpl) validatePaymentAndSpendingTypeID(ctx context.Context, req *domain.RequestValidatePaymentAndSpendingTypeID) error {
	_, err := s.spendingTypeRepo.GetByIDAndProfileID(ctx, req.SpendingTypeID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return InvalidSpendingTypeID
	}

	if req.PaymentMethodID != "" {
		_, err = s.paymentRepo.GetByIDAndProfileID(ctx, req.PaymentMethodID, req.ProfileID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			return InvalidPaymentMethodID
		}
	}

	return nil
}
