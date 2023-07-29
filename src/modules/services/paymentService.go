package services

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/dto"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/entities"
)

func (service *PaymentServiceImpl) CreatePayment(ctx context.Context) (*entities.Payment, error) {
	return nil, nil
}

func (service *PaymentServiceImpl) UpdatePayment(ctx context.Context, req *dto.PaymentUpdateRequest) (*dto.PaymentResponse, error) {
	return nil, nil
}

func (service *PaymentServiceImpl) GetPaymentById(ctx context.Context, id string) (*dto.PaymentResponse, error) {
	return nil, nil
}

func (service *PaymentServiceImpl) GetPaymentByName(ctx context.Context, name string) (*dto.PaymentResponse, error) {
	return nil, nil
}
