package _usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	resp "github.com/jasanya-tech/jasanya-response-backend-golang"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helper"
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
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
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
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
	}

	var id string

	err = s.spendingHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		if balance == nil {
			balance = converter.CreateBalanceToModel(req.ProfileID)
			err = s.balanceRepo.Create(ctx, balance)
			if err != nil {
				return resp.HttpErrString(string(resp.S500), resp.S500)
			}
		}

		spendingHistory := converter.CreateSpendingHistoryToModel(req, balance.Balance)
		err = s.spendingHistoryRepo.Create(ctx, spendingHistory)
		if err != nil {
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}
		id = spendingHistory.ID

		amountSpending := balance.TotalSpendingAmount + req.SpendingAmount
		amountBalance := balance.Balance - req.SpendingAmount
		balance = converter.UpdateBalanceToModel(balance.ID, req.ProfileID, amountSpending, amountBalance)
		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
		return nil, resp.HttpErrString(string(resp.S404), resp.S404)
	}

	spendingHistoryJoinResponse := converter.SpendingHistoryJoinModelToResponse(spendingHistoryJoin)

	return spendingHistoryJoinResponse, nil
}

func (s *SpendingHistoryUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateSpendingHistory) (*domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
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
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
	}

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
		return nil, resp.HttpErrString(string(resp.S404), resp.S404)
	}

	err = s.spendingHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		if balance == nil {
			balance = converter.CreateBalanceToModel(req.ProfileID)
			err = s.balanceRepo.Create(ctx, balance)
			if err != nil {
				return resp.HttpErrString(string(resp.S500), resp.S500)
			}
		}

		beforeBalance := balance.Balance + spendingHistoryJoin.SpendingAmount
		amountSpending := balance.TotalSpendingAmount + req.SpendingAmount - spendingHistoryJoin.SpendingAmount
		amountBalance := balance.Balance - req.SpendingAmount + spendingHistoryJoin.SpendingAmount
		balance = converter.UpdateBalanceToModel(balance.ID, req.ProfileID, amountSpending, amountBalance)

		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}

		spendingHistory := converter.UpdateSpendingHistoryToModel(req, beforeBalance)
		err = s.spendingHistoryRepo.Update(ctx, spendingHistory)
		if err != nil {
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	spendingHistoryJoin, err = s.spendingHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
		return nil, resp.HttpErrString(string(resp.S404), resp.S404)
	}

	spendingHistoryJoinResponse := converter.SpendingHistoryJoinModelToResponse(spendingHistoryJoin)

	return spendingHistoryJoinResponse, nil
}

func (s *SpendingHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return resp.HttpErrString(string(resp.S500), resp.S500)
	}
	defer s.spendingHistoryRepo.CloseConn()

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp.HttpErrString(string(resp.S404), resp.S404)
		}
		return resp.HttpErrString(string(resp.S500), resp.S500)
	}

	balance, err := s.balanceRepo.GetByProfileID(ctx, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp.HttpErrString(string(resp.S404), resp.S404)
		}
		return resp.HttpErrString(string(resp.S500), resp.S500)
	}

	err = s.spendingHistoryRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		spendingAmount := balance.TotalSpendingAmount - spendingHistoryJoin.SpendingAmount
		balanceAmount := balance.Balance + spendingHistoryJoin.SpendingAmount
		balance = converter.UpdateBalanceToModel(balance.ID, profileID, spendingAmount, balanceAmount)

		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}

		err = s.spendingHistoryRepo.Delete(ctx, id, profileID)
		if err != nil {
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *SpendingHistoryUsecaseImpl) GetAllByTimeAndProfileID(ctx context.Context, req *domain.RequestGetFilteredDataSpendingHistory) (*[]domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}
	defer s.spendingHistoryRepo.CloseConn()

	var startTime, endTime time.Time
	if req.Type != "" {
		startTime, endTime, err = helper.TimeDateByTypeFilter(req.Type)
		if err != nil {
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
	} else {
		startTime, endTime, err = helper.FormatDate(req.StartTime, req.EndTime)
		if err != nil {
			msg := resp.RegisterErrMapOfSlices("start", "invalid start time query param")
			msg2 := resp.RegisterErrMapOfSlices("start", "invalid start time query param")
			return nil, resp.HttpErrMapOfSlices(resp.MergeMapOfSlices(msg, msg2), resp.S400)
		}
	}
	req.StartTime = startTime
	req.EndTime = endTime

	spendingHistories, err := s.spendingHistoryRepo.GetAllByTimeAndProfileID(ctx, req)
	if err != nil {
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}

	var spendingHistoryJoinResponses []domain.ResponseSpendingHistory

	for _, spendingHistory := range *spendingHistories {
		spendingHistoryJoinResponse := converter.GetAllSpendingHistoryJoinModelToResponse(spendingHistory)
		spendingHistoryJoinResponses = append(spendingHistoryJoinResponses, spendingHistoryJoinResponse)
	}

	return &spendingHistoryJoinResponses, nil
}

func (s *SpendingHistoryUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}
	defer s.spendingHistoryRepo.CloseConn()

	spendingHistory, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, resp.HttpErrString(string(resp.S404), resp.S404)
		}
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}

	spendingHistoryJoinResponse := converter.SpendingHistoryJoinModelToResponse(spendingHistory)

	return spendingHistoryJoinResponse, nil
}

func (s *SpendingHistoryUsecaseImpl) validatePaymentAndSpendingTypeID(ctx context.Context, req *domain.RequestValidatePaymentAndSpendingTypeID) error {
	errBad := map[string][]string{}
	_, err := s.spendingTypeRepo.GetByIDAndProfileID(ctx, req.SpendingTypeID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}
		errBad["spending_type_id"] = append(errBad["spending_type_id"], "invalid spending type id")
	}

	if req.PaymentMethodID != "" {
		_, err = s.paymentRepo.GetByID(ctx, req.PaymentMethodID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return resp.HttpErrString(string(resp.S500), resp.S500)
			}
			errBad["payment_method_id"] = append(errBad["payment_method_id"], "invalid payment method id")
		}
	}

	if len(errBad) != 0 {
		return resp.HttpErrMapOfSlices(errBad, resp.S400)
	}

	return nil
}
