package payment_usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (p *PaymentUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	payment, err := p.paymentRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msgf("payment tidak tersedia saat delete | id : %s", id)
			return usecase.PaymentNotExist
		}
		return err
	}
	reqDeleteImageCondition := !strings.Contains(payment.Image, "default")

	err = p.paymentRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
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
