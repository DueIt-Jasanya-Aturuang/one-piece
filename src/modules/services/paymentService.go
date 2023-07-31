package services

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/config"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/http-protocol/exception"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/utils/validation"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/dto"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/entities"
	satoriuuid "github.com/satori/go.uuid"
)

func (service *PaymentServiceImpl) CreatePayment(ctx context.Context, req *dto.PaymentCreateRequest) (*entities.Payment, error) {
	err := service.Validation.Struct(req)
	if err != nil {
		return nil, err
	}

	var paymentEntity *entities.Payment

	err = service.DbImpl.RunWithTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		payment, err := service.PaymentRepository.GetPaymentByName(ctx, service.DbImpl.DB, req.Name)

		if err != nil {
			return err
		}

		if payment != nil {
			return exception.BadRequest(map[string][]string{
				"message": {"Payment Has Been Created"},
			})
		}

		fileName := service.Minio.GenerateFileName(req.Image, "payment-images/public/", "")
		if req.Image != nil && req.Image.Size > 0 {
			if err := validation.CheckContentType(req.Image, 2097152, req.Image.Size, "image/png", "image/jpeg", "image/jpg"); err != nil {
				return exception.BadRequest(map[string][]string{
					"image": {
						err.Error(),
					},
				})
			}

			if err := service.Minio.UploadFile(ctx, req.Image, fileName); err != nil {
				return err
			}
		}

		payment, err = service.PaymentRepository.CreatePayment(ctx, tx, service.Convert.CreatePaymentToEntity(satoriuuid.NewV4().String()))
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

func (service *PaymentServiceImpl) UpdatePayment(ctx context.Context, req *dto.PaymentUpdateRequest, id string) (*dto.PaymentResponse, error) {
	err := service.Validation.Struct(req)
	if err != nil {
		return nil, err
	}

	err = service.DbImpl.RunWithTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		payment, err := service.PaymentRepository.GetPaymentById(ctx, service.DbImpl.DB, id)
		if err != nil {
			return err
		}

		if payment == nil {
			return exception.NotFound("Payment Not Found")
		}

		service.Convert.PaymentEntity = payment
		oldImage := payment.Image

		fileName := service.Minio.GenerateFileName(req.Image, "user-images/public/", "")
		if req.Image != nil && req.Image.Size > 0 {
			payment.Image = fmt.Sprintf("%s/%s/%s", config.Get().ThirdParty.Minio.Endpoint, config.Get().ThirdParty.Minio.Bucket, fileName)
		}
		service.Convert.PaymentUpdateRequest = req

		if req.Image != nil && req.Image.Size > 0 {
			if err := validation.CheckContentType(req.Image, 2097152, req.Image.Size, "image/png", "image/jpeg", "image/jpg"); err != nil {
				return exception.BadRequest(map[string][]string{
					"image": {
						err.Error(),
					},
				})
			}

			if err := service.Minio.UploadFile(ctx, req.Image, fileName); err != nil {
				return err
			}

			if !strings.Contains(oldImage, "default-male") {
				oldImageSplit := strings.Split(oldImage, "/")
				objectName := fmt.Sprintf("%s/%s/%s", oldImageSplit[2], oldImageSplit[3], oldImageSplit[4])
				if err := service.Minio.DeleteFile(ctx, objectName); err != nil {
					return err
				}
			}
		}

		_, err = service.PaymentRepository.UpdatePayment(ctx, tx, service.Convert.UpdatePaymentToEntity())
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return service.Convert.PaymentEntityToResponse(), nil
}

func (service *PaymentServiceImpl) GetPaymentByName(ctx context.Context, name string) (*dto.PaymentResponse, error) {
	return nil, nil
}
