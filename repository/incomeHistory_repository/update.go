package incomeHistory_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (i *IncomeHistoryRepositoryImpl) Update(ctx context.Context, income *repository.IncomeHistory) error {
	query := `UPDATE t_income_history SET income_type_id=$1, payment_method_id=$2, payment_name=$3, income_amount=$4, description=$5,
                              time_income_history=$6, show_time_income_history=$7, updated_at=$8, updated_by=$9 
                          WHERE id =$10 AND profile_id = $11 AND deleted_at IS NULL`

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
		income.IncomeTypeID,
		income.PaymentMethodID,
		income.PaymentName,
		income.IncomeAmount,
		income.Description,
		income.TimeIncomeHistory,
		income.ShowTimeIncomeHistory,
		income.UpdatedAt,
		income.UpdatedBy,
		income.ID,
		income.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
