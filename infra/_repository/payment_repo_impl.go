package _repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type PaymentRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewPaymentRepositoryImpl(
	uow domain.UnitOfWorkRepository,
) domain.PaymentRepository {
	return &PaymentRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (p *PaymentRepositoryImpl) CreatePayment(ctx context.Context, payment *domain.Payment) error {
	query := `INSERT INTO m_payment_methods (id, name, description, image, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	tx, err := p.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
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

func (p *PaymentRepositoryImpl) UpdatePayment(ctx context.Context, payment *domain.Payment) error {
	query := `UPDATE m_payment_methods SET name = $1, description = $2, image = $3, updated_at = $4, updated_by = $5 
            	WHERE id = $6`
	tx, err := p.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
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

func (p *PaymentRepositoryImpl) GetPaymentById(ctx context.Context, id string) (*domain.Payment, error) {
	// TODO implement me
	panic("implement me")
}

func (p *PaymentRepositoryImpl) GetPaymentByName(ctx context.Context, name string) (*domain.Payment, error) {
	// TODO implement me
	panic("implement me")
}
