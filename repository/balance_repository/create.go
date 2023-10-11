package balance_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (b *BalanceRepositoryImpl) Create(ctx context.Context, balance *repository.Balance) error {
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
