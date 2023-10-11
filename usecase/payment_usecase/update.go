package payment_usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (p *PaymentUsecaseImpl) Update(ctx context.Context, req *usecase.RequestUpdatePayment) (*usecase.ResponsePayment, error) {
	payment, err := p.paymentRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, usecase.PaymentNotExist
	}

	if payment.Name != req.Name {
		paymentName, err := p.paymentRepo.GetByNameAndProfileID(ctx, req.Name, req.ProfileID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
		}
		if paymentName != nil {
			return nil, usecase.NamePaymentExist
		}
	}

	fileName := payment.Image
	reqImageCondition := req.Image != nil && req.Image.Size > 0
	reqDeleteImageCondition := !strings.Contains(fileName, "default")

	if reqImageCondition {
		fileExt := filepath.Ext(req.Image.Filename)
		fileName = p.minioRepo.GenerateFileName(fileExt, "payment-images/public/")
		fileName = fmt.Sprintf("/%s/%s", infra.MinIoBucket, fileName)
	}

	paymentConv := req.ToModel(fileName)

	err = p.paymentRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		if err = p.paymentRepo.Update(ctx, paymentConv); err != nil {
			return err
		}

		if reqImageCondition {
			if err = p.minioRepo.UploadFile(ctx, req.Image, fileName); err != nil {
				return err
			}

			if reqDeleteImageCondition {
				imageDelArr := strings.Split(payment.Image, "/")
				imageDel := fmt.Sprintf("/%s/%s/%s", imageDelArr[2], imageDelArr[3], imageDelArr[4])
				if err = p.minioRepo.DeleteFile(ctx, imageDel); err != nil {
					return err
				}
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	resp := &usecase.ResponsePayment{
		ID:          paymentConv.ID,
		ProfileID:   paymentConv.ProfileID,
		Name:        paymentConv.Name,
		Description: repository.GetNullString(paymentConv.Description),
		Image:       paymentConv.Image,
	}
	return resp, nil
}
