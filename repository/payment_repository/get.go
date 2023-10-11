package payment_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *PaymentRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*repository.Payment, error) {
	query := `SELECT id, profile_id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_payment_methods WHERE id = $1 AND profile_id = $2 AND deleted_at IS NULL`

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

	var payment repository.Payment

	if err = stmt.QueryRowContext(ctx, id, profileID).Scan(
		&payment.ID,
		&payment.ProfileID,
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
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &payment, nil
}

func (p *PaymentRepositoryImpl) GetByNameAndProfileID(ctx context.Context, name string, profileID string) (*repository.Payment, error) {
	query := `SELECT id, profile_id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_payment_methods WHERE name = $1 AND profile_id = $2 AND deleted_at IS NULL`

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

	var payment repository.Payment

	if err = stmt.QueryRowContext(ctx, name, profileID).Scan(
		&payment.ID,
		&payment.ProfileID,
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
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &payment, nil
}

func (p *PaymentRepositoryImpl) GetAllByProfileID(ctx context.Context, req *repository.GetAllPaymentWithISD) (*[]repository.Payment, error) {
	query := `SELECT id, profile_id, name, description, image, created_at, created_by, 
       			updated_at, updated_by, deleted_at, deleted_by
          FROM m_payment_methods WHERE profile_id = $1 AND deleted_at IS NULL `
	if req.ID != "" {
		query += `AND id ` + req.Operation + ` $2 `
	}
	query += `ORDER BY id ` + req.Order + ` LIMIT 5`

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

	var rows *sql.Rows
	if req.ID != "" {
		rows, err = stmt.QueryContext(ctx, req.ProfileID, req.ID)
	} else {
		rows, err = stmt.QueryContext(ctx, req.ProfileID)
	}

	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
		return nil, err
	}

	var payments []repository.Payment
	var payment repository.Payment

	for rows.Next() {
		if err = rows.Scan(
			&payment.ID,
			&payment.ProfileID,
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
