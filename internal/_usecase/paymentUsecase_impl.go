package _usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/converter"
)

var NamePaymentExist = errors.New("nama payment sudah tersedia") // create, update
var PaymentNotExist = errors.New("payment tidak tersedia")       // update

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

	paymentCheck, err := p.paymentRepo.GetByName(ctx, req.Name)
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

	payment, err := p.paymentRepo.GetByID(ctx, req.ID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, PaymentNotExist
	}

	if payment.Name != req.Name {
		paymentName, err := p.paymentRepo.GetByName(ctx, req.Name)
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

	if reqImageCondition {
		fileExt := filepath.Ext(req.Image.Filename)
		fileName = p.minioRepo.GenerateFileName(fileExt, "payment-images/public/")
	}

	paymentConv := converter.UpdatePaymentReqToModel(req, fmt.Sprintf("/%s/%s", infra.MinIoBucket, fileName))

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

	paymentResponse := converter.PaymentModelToResp(paymentConv)
	return paymentResponse, nil
}

func (p *PaymentUsecaseImpl) GetAll(ctx context.Context) (*[]domain.ResponsePayment, error) {
	if err := p.paymentRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer p.paymentRepo.CloseConn()

	payments, err := p.paymentRepo.GetAll(ctx)
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
