package payment_usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (p *PaymentUsecaseImpl) GetAllByProfileID(ctx context.Context, req *usecase.RequestGetAllByProfileIDWithISD) (*[]usecase.ResponsePayment, string, error) {
	exist, err := p.paymentRepo.CheckData(ctx, req.ProfileID)
	if err != nil {
		return nil, "", err
	}

	if !exist {
		err = p.createDefaultPayment(ctx, req.ProfileID)
		if err != nil {
			return nil, "", err
		}
	}

	order, operation := usecase.GetOrder(req.Order)
	payments, err := p.paymentRepo.GetAllByProfileID(ctx, &repository.GetAllPaymentWithISD{
		ProfileID: req.ProfileID,
		InfiniteScrollData: repository.InfiniteScrollData{
			ID:        req.ID,
			Order:     order,
			Operation: operation,
		},
	})
	if err != nil {
		return nil, "", err
	}

	if len(*payments) < 1 {
		return nil, "", nil
	}

	var responses []usecase.ResponsePayment
	var response *usecase.ResponsePayment

	for _, payment := range *payments {
		response = &usecase.ResponsePayment{
			ID:          payment.ID,
			ProfileID:   payment.ProfileID,
			Name:        payment.Name,
			Description: repository.GetNullString(payment.Description),
			Image:       payment.Image,
		}

		responses = append(responses, *response)

	}

	cursor := (*payments)[len(*payments)-1].ID
	return &responses, cursor, nil
}
