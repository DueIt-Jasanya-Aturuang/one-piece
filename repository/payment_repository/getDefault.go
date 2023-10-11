package payment_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *PaymentRepositoryImpl) GetDefault(ctx context.Context) (*[]repository.Payment, error) {
	query := `SELECT id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_default_payment_method WHERE deleted_at IS NULL`

	db, err := p.GetDB()
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
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
		return nil, err
	}

	var payments []repository.Payment
	var payment repository.Payment

	for rows.Next() {
		if err = rows.Scan(
			&payment.ID,
			&payment.Name,
			&payment.Description,
			&payment.Image,
			&payment.CreatedAt,
			&payment.CreatedBy,
			&payment.UpdatedAt,
			&payment.UpdatedBy,
			&payment.DeletedAt,
			&payment.DeletedBy,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
			return nil, err
		}

		payments = append(payments, payment)
	}

	return &payments, nil
}
