package incomeType_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type IncomeTypeUsecaseImpl struct {
	incomeTypeRepo repository.IncomeTypeRepository
}

func NewIncomeTypeUsecaseImpl(
	incomeTypeRepo repository.IncomeTypeRepository,
) usecase.IncomeTypeUsecase {
	return &IncomeTypeUsecaseImpl{
		incomeTypeRepo: incomeTypeRepo,
	}
}
