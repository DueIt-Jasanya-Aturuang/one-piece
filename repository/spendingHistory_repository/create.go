package spendingHistory_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingHistoryRepositoryImpl) Create(ctx context.Context, spendingHistory *repository.SpendingHistory) error {
	query := `INSERT INTO t_spending_history (id, profile_id, spending_type_id, payment_method_id, payment_name, 
                        before_balance, spending_amount, after_balance, description, time_spending_history, show_time_spending_history, 
                        created_at, created_by, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

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
		spendingHistory.ID,
		spendingHistory.ProfileID,
		spendingHistory.SpendingTypeID,
		spendingHistory.PaymentMethodID,
		spendingHistory.PaymentName,
		spendingHistory.BeforeBalance,
		spendingHistory.SpendingAmount,
		spendingHistory.AfterBalance,
		spendingHistory.Description,
		spendingHistory.TimeSpendingHistory,
		spendingHistory.ShowTimeSpendingHistory,
		spendingHistory.CreatedAt,
		spendingHistory.CreatedBy,
		spendingHistory.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
