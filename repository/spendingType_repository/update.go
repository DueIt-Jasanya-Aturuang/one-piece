package spendingType_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingTypeRepositoryImpl) Update(ctx context.Context, spendingType *repository.SpendingType) error {
	query := `UPDATE m_spending_type SET title = $1, maximum_limit = $2, icon = $3, updated_at = $4, updated_by = $5 
                    WHERE id = $6 AND profile_id = $7 AND deleted_at IS NULL`

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
		spendingType.Title,
		spendingType.MaximumLimit,
		spendingType.Icon,
		spendingType.UpdatedAt,
		spendingType.UpdatedBy,
		spendingType.ID,
		spendingType.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
