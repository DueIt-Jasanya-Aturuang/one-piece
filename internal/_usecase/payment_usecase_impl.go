package _usecase

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type PaymentUsecaseImpl struct {
	paymentRepo domain.PaymentRepository
}

func NewPaymentUsecaseImpl(
	paymentRepo domain.PaymentRepository,
) domain.PaymentUsecase {
	return &PaymentUsecaseImpl{
		paymentRepo: paymentRepo,
	}
}

func (p *PaymentUsecaseImpl) CreatePayment(
	ctx context.Context, req *domain.RequestCreatePayment,
) (resp *domain.ResponsePayment, err error) {
	if err = p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer p.paymentRepo.CloseConn()

	err = p.paymentRepo.StartTx(ctx, &sql.TxOptions{}, func() error {
		return nil
	})

	return nil, err
}

func (p *PaymentUsecaseImpl) UpdatePayment(ctx context.Context, req *domain.RequestUpdatePayment, id string) (*domain.ResponsePayment, error) {
	// TODO implement me
	panic("implement me")
}

func (p *PaymentUsecaseImpl) GetPaymentByName(ctx context.Context, name string) (*domain.ResponsePayment, error) {
	// TODO implement me
	panic("implement me")
}

func (p *PaymentUsecaseImpl) DeletePayment(ctx context.Context, id string) (bool, error) {
	// TODO implement me
	panic("implement me")
}
