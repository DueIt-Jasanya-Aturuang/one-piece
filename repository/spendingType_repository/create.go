package spendingType_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingTypeRepositoryImpl) Create(ctx context.Context, spendingType *repository.SpendingType) error {
	query := `INSERT INTO m_spending_type (id, profile_id, title, maximum_limit, icon, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

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
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		spendingType.ID,
		spendingType.ProfileID,
		spendingType.Title,
		spendingType.MaximumLimit,
		spendingType.Icon,
		spendingType.CreatedAt,
		spendingType.CreatedBy,
		spendingType.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
