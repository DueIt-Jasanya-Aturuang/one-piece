package payment_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *PaymentRepositoryImpl) CheckData(ctx context.Context, profileID string) (bool, error) {
	query := `SELECT EXISTS(SELECT id FROM m_payment_methods WHERE profile_id = $1);`

	db, err := p.GetDB()
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
