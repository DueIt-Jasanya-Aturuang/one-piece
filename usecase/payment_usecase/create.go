package payment_usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *PaymentUsecaseImpl) Create(ctx context.Context, req *usecase.RequestCreatePayment) (*usecase.ResponsePayment, error) {
	paymentCheck, err := p.paymentRepo.GetByNameAndProfileID(ctx, req.Name, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}
	if paymentCheck != nil {
		return nil, usecase.NamePaymentExist
	}

	fileExt := filepath.Ext(req.Image.Filename)
	fileName := p.minioRepo.GenerateFileName(fileExt, "payment-images/public/")
	payment := req.ToModel(fmt.Sprintf("/%s/%s", infra.MinIoBucket, fileName))

	err = p.paymentRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err = p.paymentRepo.Create(ctx, payment)
		if err != nil {
			return err
		}

		err = p.minioRepo.UploadFile(ctx, req.Image, fileName)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	resp := &usecase.ResponsePayment{
		ID:          payment.ID,
		ProfileID:   payment.ProfileID,
		Name:        payment.Name,
		Description: repository.GetNullString(payment.Description),
		Image:       payment.Image,
	}
	return resp, nil
}

func (p *PaymentUsecaseImpl) createDefaultPayment(ctx context.Context, profileID string) error {
	err := p.paymentRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		payments, err := p.paymentRepo.GetDefault(ctx)
		if err != nil {
			return err
		}

		for _, payment := range *payments {
			payment.ProfileID = profileID
			payment.ID = util.NewUlid()
			err = p.paymentRepo.Create(ctx, &payment)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
