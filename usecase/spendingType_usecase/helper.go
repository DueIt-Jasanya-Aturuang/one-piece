package spendingType_usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingTypeUsecaseImpl) createDefaultSpendingType(ctx context.Context, profileID string) error {
	err := s.spendingTypeRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		spendingTypes, err := s.spendingTypeRepo.GetDefault(ctx)
		if err != nil {
			return err
		}

		for _, spendingType := range *spendingTypes {
			spendingType.ProfileID = profileID
			spendingType.ID = util.NewUlid()
			err = s.spendingTypeRepo.Create(ctx, &spendingType)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
