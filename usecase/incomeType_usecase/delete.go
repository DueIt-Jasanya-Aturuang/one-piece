package incomeType_usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

func (i *IncomeTypeUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	err := i.incomeTypeRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err := i.incomeTypeRepo.Delete(ctx, id, profileID)
		if err != nil {
			return nil
		}

		return nil
	})

	return err
}
