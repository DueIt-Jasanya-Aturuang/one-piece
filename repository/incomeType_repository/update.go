package incomeType_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (i *IncomeTypeRepositoryImpl) Update(ctx context.Context, income *repository.IncomeType) error {
	query := `UPDATE m_income_type SET name = $1, description = $2, icon = $3, income_type = $4, fixed_income = $5,
                         periode = $6, amount = $7, updated_at = $8, updated_by = $9 
                    WHERE id = $10 AND profile_id = $11 AND deleted_at IS NULL`

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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		income.Name,
		income.Description,
		income.Icon,
		income.IncomeType,
		income.FixedIncome,
		income.Periode,
		income.Amount,
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
