package _usecase

import (
	"context"

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
	// TODO implement me
	panic("implement me")
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
