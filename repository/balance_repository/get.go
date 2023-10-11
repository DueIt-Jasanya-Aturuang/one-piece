package balance_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (b *BalanceRepositoryImpl) GetByProfileID(ctx context.Context, profileID string) (*repository.Balance, error) {
	query := `SELECT id, profile_id, total_income_amount, total_spending_amount, balance,
					created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
       FROM m_balance WHERE profile_id = $1`

	db, err := b.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var balance repository.Balance
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
