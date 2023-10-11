package incomeHistory_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (i *IncomeHistoryRepositoryImpl) Create(ctx context.Context, income *repository.IncomeHistory) error {
	query := `INSERT INTO t_income_history(id, profile_id, income_type_id, payment_method_id, payment_name, income_amount,
                               description, time_income_history, show_time_income_history, created_at, created_by, updated_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	tx, err := i.GetTx()
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
		income.ID,
		income.ProfileID,
		income.IncomeTypeID,
		income.PaymentMethodID,
		income.PaymentName,
		income.IncomeAmount,
		income.Description,
		income.TimeIncomeHistory,
		income.ShowTimeIncomeHistory,
		income.CreatedAt,
		income.CreatedBy,
		income.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
