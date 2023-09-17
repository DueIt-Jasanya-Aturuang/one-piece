package _usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	resp "github.com/jasanya-tech/jasanya-response-backend-golang"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/config"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/converter"
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
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}
	defer p.paymentRepo.CloseConn()

	paymentCheck, err := p.paymentRepo.GetByName(ctx, req.Name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
	}
	if paymentCheck != nil {
		msg := resp.RegisterErrMapOfSlices("name", "nama payment sudah tersedia")
		return nil, resp.HttpErrMapOfSlices(msg, resp.S400)
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
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}

		err = p.minioRepo.UploadFile(ctx, req.Image, fileName)
		if err != nil {
			return resp.HttpErrString(string(resp.S500), resp.S500)
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
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}
	defer p.paymentRepo.CloseConn()

	payment, err := p.paymentRepo.GetByID(ctx, req.ID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, resp.HttpErrString(string(resp.S500), resp.S500)
		}
		return nil, resp.HttpErrString("payment tidak tersedia", resp.S404)
	}

	if payment.Name != req.Name {
		paymentName, err := p.paymentRepo.GetByName(ctx, req.Name)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, resp.HttpErrString(string(resp.S500), resp.S500)
			}
		}
		if paymentName != nil {
			msg := resp.RegisterErrMapOfSlices("name", "nama payment sudah tersedia")
			return nil, resp.HttpErrMapOfSlices(msg, resp.S400)
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
			return resp.HttpErrString(string(resp.S500), resp.S500)
		}

		if reqImageCondition {
			if err = p.minioRepo.UploadFile(ctx, req.Image, fileName); err != nil {
				return resp.HttpErrString(string(resp.S500), resp.S500)
			}

			imageDelArr := strings.Split(payment.Image, "/")
			imageDel := fmt.Sprintf("/%s/%s/%s", imageDelArr[2], imageDelArr[3], imageDelArr[4])
			if err = p.minioRepo.DeleteFile(ctx, imageDel); err != nil {
				return resp.HttpErrString(string(resp.S500), resp.S500)
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

func (p *PaymentUsecaseImpl) GetAll(ctx context.Context) (*[]domain.ResponsePayment, error) {
	if err := p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}
	defer p.paymentRepo.CloseConn()

	payments, err := p.paymentRepo.GetAll(ctx)
	if err != nil {
		return nil, resp.HttpErrString(string(resp.S500), resp.S500)
	}

	var responses []domain.ResponsePayment

	for _, payment := range *payments {
		paymentConv := converter.PaymentGetAllPaymentModelToResp(payment)
		responses = append(responses, paymentConv)
	}

	return &responses, nil
}
