package spendingType_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingTypeRepositoryImpl) CheckData(ctx context.Context, profileID string) (bool, error) {
	query := `SELECT EXISTS(SELECT id FROM m_spending_type WHERE profile_id = $1);`

	db, err := s.GetDB()
	if err != nil {
		return false, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var exist bool
	if err = stmt.QueryRowContext(ctx, profileID).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	return exist, nil
}

func (s *SpendingTypeRepositoryImpl) CheckByTitleAndProfileID(ctx context.Context, profileID string, title string) (bool, error) {
	query := `SELECT EXISTS(SELECT id FROM m_spending_type WHERE profile_id = $1 AND title = $2 AND deleted_at IS NULL);`

	db, err := s.GetDB()
	if err != nil {
		return false, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var exist bool
	if err = stmt.QueryRowContext(ctx, profileID, title).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	return exist, nil
}
