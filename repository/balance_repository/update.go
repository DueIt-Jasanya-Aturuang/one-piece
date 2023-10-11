package balance_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (b *BalanceRepositoryImpl) UpdateByProfileID(ctx context.Context, balance *repository.Balance) error {
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
