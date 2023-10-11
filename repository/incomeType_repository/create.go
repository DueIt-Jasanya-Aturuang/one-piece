package incomeType_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (i *IncomeTypeRepositoryImpl) Create(ctx context.Context, income *repository.IncomeType) error {
	query := `INSERT INTO m_income_type (id, profile_id, name, description, icon, income_type, fixed_income, periode,
                             amount, created_at, created_by, updated_at) 
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
		income.Name,
		income.Description,
		income.Icon,
		income.IncomeType,
		income.FixedIncome,
		income.Periode,
		income.Amount,
		income.CreatedAt,
		income.CreatedBy,
		income.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
