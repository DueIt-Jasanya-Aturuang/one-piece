package spendingType_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingTypeRepositoryImpl) GetDefault(ctx context.Context) (*[]repository.SpendingType, error) {
	query := `SELECT id, title, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_default_spending_type WHERE active = true AND deleted_at IS NULL`

	db, err := s.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
		return nil, err
	}

	var spendingTypes []repository.SpendingType
	var spendingType repository.SpendingType

	for rows.Next() {
		if err = rows.Scan(
			&spendingType.ID,
			&spendingType.Title,
			&spendingType.MaximumLimit,
			&spendingType.Icon,
			&spendingType.CreatedAt,
			&spendingType.CreatedBy,
			&spendingType.UpdatedAt,
			&spendingType.UpdatedBy,
			&spendingType.DeletedAt,
			&spendingType.DeletedBy,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
			return nil, err
		}

		spendingTypes = append(spendingTypes, spendingType)
	}

	return &spendingTypes, nil
}
