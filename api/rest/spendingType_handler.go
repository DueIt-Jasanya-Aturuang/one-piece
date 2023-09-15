package rest

import (
	"net/http"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type SpendingTypeHandlerImpl struct {
	spendingTypeUsecase domain.SpendingTypeUsecase
}

func NewSpendingTypeHandlerImpl(
	spendingTypeUsecase domain.SpendingTypeUsecase,
) *SpendingTypeHandlerImpl {
	return &SpendingTypeHandlerImpl{
		spendingTypeUsecase: spendingTypeUsecase,
	}
}

func (h *SpendingTypeHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingTypeHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingTypeHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingTypeHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingTypeHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}
