package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

type PaymentUsecaseImpl struct {
	paymentRepo domain.PaymentRepository
	minioRepo   domain.MinioRepo
}

func NewPaymentUsecaseImpl(
	paymentRepo domain.PaymentRepository,
	minioRepo domain.MinioRepo,
) *PaymentUsecaseImpl {
	return &PaymentUsecaseImpl{
		paymentRepo: paymentRepo,
		minioRepo:   minioRepo,
	}
}

func (p *PaymentUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreatePayment) (*domain.ResponsePayment, error) {
	if err := p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer p.paymentRepo.CloseConn()

	paymentCheck, err := p.paymentRepo.GetByNameAndProfileID(ctx, req.Name, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}
	if paymentCheck != nil {
		return nil, NamePaymentExist
	}

	fileExt := filepath.Ext(req.Image.Filename)
	fileName := p.minioRepo.GenerateFileName(fileExt, "payment-images/public/")
	paymentConv := converter.CreatePaymentReqToModel(req, fmt.Sprintf("/%s/%s", infra.MinIoBucket, fileName))

	err = p.paymentRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		err = p.paymentRepo.Create(ctx, paymentConv)
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

	paymentResponse := converter.PaymentModelToResp(paymentConv)
	return paymentResponse, nil
}

func (p *PaymentUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdatePayment) (*domain.ResponsePayment, error) {
	if err := p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer p.paymentRepo.CloseConn()

	payment, err := p.paymentRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, PaymentNotExist
	}

	if payment.Name != req.Name {
		paymentName, err := p.paymentRepo.GetByNameAndProfileID(ctx, req.Name, req.ProfileID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
		}
		if paymentName != nil {
			return nil, NamePaymentExist
		}
	}

	fileName := payment.Image
	reqImageCondition := req.Image != nil && req.Image.Size > 0
	reqDeleteImageCondition := !strings.Contains(fileName, "default")

	if reqImageCondition {
		fileExt := filepath.Ext(req.Image.Filename)
		fileName = p.minioRepo.GenerateFileName(fileExt, "payment-images/public/")
	}

	paymentConv := converter.UpdatePaymentReqToModel(req, fmt.Sprintf("/%s/%s", infra.MinIoBucket, fileName))

	err = p.paymentRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
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

	paymentResponse := converter.PaymentModelToResp(paymentConv)
	return paymentResponse, nil
}

func (p *PaymentUsecaseImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.ResponsePayment, error) {
	if err := p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer p.paymentRepo.CloseConn()

	exist, err := p.paymentRepo.CheckData(ctx, profileID)
	if err != nil {
		return nil, err
	}

	if !exist {
		err = p.createDefaultPayment(ctx, profileID)
		if err != nil {
			return nil, err
		}
	}

	payments, err := p.paymentRepo.GetAllByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	var responses []domain.ResponsePayment

	for _, payment := range *payments {
		paymentConv := converter.PaymentGetAllPaymentModelToResp(payment)
		responses = append(responses, paymentConv)
	}

	return &responses, nil
}

func (p *PaymentUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	if err := p.paymentRepo.OpenConn(ctx); err != nil {
		return err
	}
	defer p.paymentRepo.CloseConn()

	payment, err := p.paymentRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msgf("payment tidak tersedia saat delete | id : %s", id)
			return PaymentNotExist
		}
		return err
	}
	reqDeleteImageCondition := !strings.Contains(payment.Image, "default")

	err = p.paymentRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		if err = p.paymentRepo.Delete(ctx, id, profileID); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	if reqDeleteImageCondition {
		imageDelArr := strings.Split(payment.Image, "/")
		imageDel := fmt.Sprintf("/%s/%s/%s", imageDelArr[2], imageDelArr[3], imageDelArr[4])
		if err = p.minioRepo.DeleteFile(ctx, imageDel); err != nil {
			return err
		}
	}

	return nil
}

func (p *PaymentUsecaseImpl) createDefaultPayment(ctx context.Context, profileID string) error {
	err := p.paymentRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		payments, err := p.paymentRepo.GetDefault(ctx)
		if err != nil {
			return err
		}

		for _, payment := range *payments {
			payment.ProfileID = profileID
			payment.ID = uuid.NewV4().String()
			err = p.paymentRepo.Create(ctx, &payment)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
