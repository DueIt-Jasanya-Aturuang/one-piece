package _usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/config"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
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

func (p *PaymentUsecaseImpl) Create(
	ctx context.Context, req *domain.RequestCreatePayment,
) (resp *domain.ResponsePayment, err error) {
	if err = p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, util.ErrHTTPString("", 500)
	}
	defer p.paymentRepo.CloseConn()

	paymentCheck, err := p.paymentRepo.GetByName(ctx, req.Name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("", 500)
		}
	}
	if paymentCheck != nil {
		return nil, util.ErrHTTPString("nama payment sudah tersedia", 409)
	}

	fileExt := filepath.Ext(req.Image.Filename)
	fileName := p.minioRepo.GenerateFileName(fileExt, "payment-images/public/")
	paymentConv := converter.CreatePaymentReqToModel(req, fmt.Sprintf("/%s/%s", config.MinIoBucket, fileName))

	err = p.paymentRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
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
		return nil, util.ErrHTTPString("", 500)
	}

	resp = converter.PaymentModelToResp(paymentConv)
	return resp, nil
}

func (p *PaymentUsecaseImpl) Update(
	ctx context.Context, req *domain.RequestUpdatePayment,
) (resp *domain.ResponsePayment, err error) {
	if err = p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, util.ErrHTTPString("", 500)
	}
	defer p.paymentRepo.CloseConn()

	payment, err := p.paymentRepo.GetByID(ctx, req.ID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("", 500)
		}
		return nil, util.ErrHTTPString("payment tidak tersedia", 404)
	}
	if payment.Name != req.Name {
		paymentName, err := p.paymentRepo.GetByName(ctx, req.Name)
		if err != nil {
			log.Info().Err(err).Msgf("err")
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, util.ErrHTTPString("", 500)
			}
		}
		if paymentName != nil {
			return nil, util.ErrHTTPString("nama payment sudah tersedia", 409)
		}
	}

	fileName := payment.Image
	reqImageCondition := req.Image != nil && req.Image.Size > 0

	if reqImageCondition {
		fileExt := filepath.Ext(req.Image.Filename)
		fileName = p.minioRepo.GenerateFileName(fileExt, "payment-images/public/")
	}

	paymentConv := converter.UpdatePaymentReqToModel(req, fmt.Sprintf("/%s/%s", config.MinIoBucket, fileName))

	err = p.paymentRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		if err = p.paymentRepo.Update(ctx, paymentConv); err != nil {
			return err
		}

		if reqImageCondition {
			if err = p.minioRepo.UploadFile(ctx, req.Image, fileName); err != nil {
				return err
			}

			imageDelArr := strings.Split(payment.Image, "/")
			imageDel := fmt.Sprintf("/%s/%s/%s", imageDelArr[2], imageDelArr[3], imageDelArr[4])
			if err = p.minioRepo.DeleteFile(ctx, imageDel); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	resp = converter.PaymentModelToResp(paymentConv)
	return resp, nil
}

func (p *PaymentUsecaseImpl) GetAll(ctx context.Context) (*[]domain.ResponsePayment, error) {
	if err := p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, util.ErrHTTPString("", 500)
	}
	defer p.paymentRepo.CloseConn()

	payments, err := p.paymentRepo.GetAll(ctx)
	if err != nil {
		return nil, util.ErrHTTPString("", 500)
	}

	var responses []domain.ResponsePayment

	for _, payment := range *payments {
		paymentConv := converter.PaymentGetAllPaymentModelToResp(payment)
		responses = append(responses, paymentConv)
	}

	return &responses, nil
}