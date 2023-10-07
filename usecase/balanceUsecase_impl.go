package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/converter"
)

type BalanceUsecaseImpl struct {
	balanceRepo domain.BalanceRepository
}

func NewBalanceUsecaseImpl(balanceRepo domain.BalanceRepository) domain.BalanceUsecase {
	return &BalanceUsecaseImpl{
		balanceRepo: balanceRepo,
	}
}

func (b *BalanceUsecaseImpl) GetByProfileID(ctx context.Context, profileID string) (*domain.ResponseBalance, error) {
	balance, err := b.balanceRepo.GetByProfileID(ctx, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msg("balande tidak terseida")
			return nil, BalanceNotExist
		}
		return nil, err
	}

	resp := converter.BalanceModelToResponse(balance)

	return resp, nil
}
