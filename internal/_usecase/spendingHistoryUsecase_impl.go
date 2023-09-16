package _usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type SpendingHistoryUsecaseImpl struct {
	spendingHistoryRepo domain.SpendingHistoryRepository
	spendingTypeRepo    domain.SpendingTypeRepository
	balanceRepo         domain.BalanceRepository
}

func NewSpendingHistoryUsecaseImpl(
	spendingHistoryRepo domain.SpendingHistoryRepository,
	spendingTypeRepo domain.SpendingTypeRepository,
	balanceRepo domain.BalanceRepository,
) domain.SpendingHistoryUsecase {
	return &SpendingHistoryUsecaseImpl{
		spendingHistoryRepo: spendingHistoryRepo,
		spendingTypeRepo:    spendingTypeRepo,
		balanceRepo:         balanceRepo,
	}
}

func (s *SpendingHistoryUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateSpendingHistory) (*domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer s.spendingHistoryRepo.CloseConn()

	balance, err := s.balanceRepo.GetByProfileID(ctx, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("", 500)
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
		balance = converter.UpdateBalanceToModel(balance.ID, req.ProfileID, amountSpending, amountBalance)
		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		code := helper.GetCodePqError(err)
		switch code {
		case "23503":
			return nil, util.ErrHTTPString("", 403)
		default:
			return nil, util.ErrHTTPString("", 500)
		}
	}

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("", 500)
		}
		return nil, util.ErrHTTPString("", 404)
	}

	resp := converter.SpendingHistoryModelToResponse(spendingHistoryJoin)

	return resp, nil
}

func (s *SpendingHistoryUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateSpendingHistory) (*domain.ResponseSpendingHistory, error) {
	err := s.spendingHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer s.spendingHistoryRepo.CloseConn()

	balance, err := s.balanceRepo.GetByProfileID(ctx, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("", 500)
		}
	}

	spendingHistoryJoin, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("", 500)
		}
		return nil, util.ErrHTTPString("", 404)
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
		balance = converter.UpdateBalanceToModel(balance.ID, req.ProfileID, amountSpending, amountBalance)

		err = s.balanceRepo.UpdateByProfileID(ctx, balance)
		if err != nil {
			return err
		}

		spendingHistory := converter.UpdateSpendingHistoryToModel(req, beforeBalance)
		err = s.spendingHistoryRepo.Create(ctx, spendingHistory)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		code := helper.GetCodePqError(err)
		switch code {
		case "23503":
			return nil, util.ErrHTTPString("", 403)
		default:
			return nil, util.ErrHTTPString("", 500)
		}
	}

	spendingHistoryJoin, err = s.spendingHistoryRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("", 500)
		}
		return nil, util.ErrHTTPString("", 404)
	}

	resp := converter.SpendingHistoryModelToResponse(spendingHistoryJoin)

	return resp, nil
}

func (s *SpendingHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryUsecaseImpl) GetAllByTimeAndProfileID(ctx context.Context, req *domain.RequestGetFilteredDataSpendingHistory) (*[]domain.ResponseSpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryUsecaseImpl) GetByIDAndProfileID(
	ctx context.Context, id string, profileID string,
) (*domain.ResponseSpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}
