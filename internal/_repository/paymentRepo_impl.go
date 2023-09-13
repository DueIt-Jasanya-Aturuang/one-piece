package _repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type PaymentRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewPaymentRepositoryImpl(
	uow domain.UnitOfWorkRepository,
) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (p *PaymentRepositoryImpl) Create(ctx context.Context, payment *domain.Payment) error {
	query := `INSERT INTO m_payment_methods (id, name, description, image, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	tx, err := p.GetTx()
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

	_, err = stmt.ExecContext(
		ctx,
		payment.ID,
		payment.Name,
		payment.Description,
		payment.Image,
		payment.CreatedAt,
		payment.CreatedBy,
		payment.UpdatedAt,
	)
	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (p *PaymentRepositoryImpl) Update(ctx context.Context, payment *domain.Payment) error {
	query := `UPDATE m_payment_methods SET name = $1, description = $2, image = $3, updated_at = $4, updated_by = $5 
            	WHERE id = $6`
	tx, err := p.GetTx()
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

	_, err = stmt.ExecContext(
		ctx,
		payment.Name,
		payment.Description,
		payment.Image,
		payment.UpdatedAt,
		payment.UpdatedBy,
		payment.ID,
	)
	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (p *PaymentRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	query := `SELECT id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_payment_methods WHERE id = $1`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var payment domain.Payment

	if err = stmt.QueryRowContext(ctx, id).Scan(
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
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &payment, nil
}

func (p *PaymentRepositoryImpl) GetByName(ctx context.Context, name string) (*domain.Payment, error) {
	query := `SELECT id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_payment_methods WHERE name = $1`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var payment domain.Payment

	if err = stmt.QueryRowContext(ctx, name).Scan(
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
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &payment, nil
}

func (p *PaymentRepositoryImpl) GetAll(ctx context.Context) (*[]domain.Payment, error) {
	query := `SELECT id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_payment_methods`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
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

	var payments []domain.Payment

	for rows.Next() {
		var payment domain.Payment

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
