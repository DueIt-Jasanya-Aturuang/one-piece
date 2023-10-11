package spendingHistory_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingHistoryRepositoryImpl) Update(ctx context.Context, spendingHistory *repository.SpendingHistory) error {
	query := `UPDATE t_spending_history SET spending_type_id = $1, payment_method_id = $2, payment_name = $3, before_balance = $4, 
                              spending_amount = $5, after_balance = $6, description = $7, time_spending_history = $8, 
                              show_time_spending_history = $9, updated_at = $10, updated_by = $11
                          WHERE id = $12 AND profile_id = $13 AND deleted_at IS NULL`

	tx, err := s.GetTx()
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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		spendingHistory.SpendingTypeID,
		spendingHistory.PaymentMethodID,
		spendingHistory.PaymentName,
		spendingHistory.BeforeBalance,
		spendingHistory.SpendingAmount,
		spendingHistory.AfterBalance,
		spendingHistory.Description,
		spendingHistory.TimeSpendingHistory,
		spendingHistory.ShowTimeSpendingHistory,
		spendingHistory.UpdatedAt,
		spendingHistory.UpdatedBy,
		spendingHistory.ID,
		spendingHistory.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
