package _repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type BalanceRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewBalanceRepositoryImpl(
	uow domain.UnitOfWorkRepository,
) domain.BalanceRepository {
	return &BalanceRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (b *BalanceRepositoryImpl) Create(ctx context.Context, balance *domain.Balance) error {
	query := `INSERT INTO m_balance (id, profile_id, total_income_amount, total_spending_amount, balance, 
                       created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	tx, err := b.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		balance.ID,
		balance.ProfileID,
		balance.TotalIncomeAmount,
		balance.TotalSpendingAmount,
		balance.Balance,
		balance.CreatedAt,
		balance.CreatedBy,
		balance.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (b *BalanceRepositoryImpl) UpdateByProfileID(ctx context.Context, balance *domain.Balance) error {
	query := `UPDATE m_balance SET total_income_amount = $1, total_spending_amount = $2, balance = $3, updated_at = $4,
                     updated_by = $5 WHERE id = $6 AND profile_id = $7 AND deleted_at IS NULL`

	tx, err := b.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		balance.TotalIncomeAmount,
		balance.TotalSpendingAmount,
		balance.Balance,
		balance.UpdatedAt,
		balance.UpdatedBy,
		balance.ID,
		balance.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (b *BalanceRepositoryImpl) GetByProfileID(ctx context.Context, profileID string) (*domain.Balance, error) {
	query := `SELECT id, profile_id, total_income_amount, total_spending_amount, balance,
					created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
       FROM m_balance WHERE profile_id = $1`

	conn, err := b.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var balance domain.Balance
	if err = stmt.QueryRowContext(ctx, profileID).Scan(
		&balance.ID,
		&balance.ProfileID,
		&balance.TotalIncomeAmount,
		&balance.TotalSpendingAmount,
		&balance.Balance,
		&balance.CreatedAt,
		&balance.CreatedBy,
		&balance.UpdatedAt,
		&balance.UpdatedBy,
		&balance.DeletedAt,
		&balance.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &balance, nil
}
