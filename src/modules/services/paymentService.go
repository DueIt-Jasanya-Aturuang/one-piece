package services

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/dto"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/entities"
	satoriuuid "github.com/satori/go.uuid"
)

func (service *PaymentServiceImpl) CreatePayment(ctx context.Context) (*entities.Payment, error) {
	var paymentEntity *entities.Payment

	err := service.DbImpl.RunWithTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		payment, err := service.PaymentRepository.CreatePayment(ctx, tx, service.Convert.CreatePaymentToEntity(satoriuuid.NewV4().String()))
		if err != nil {
			return err
		}

		paymentEntity = payment

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paymentEntity, nil
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
