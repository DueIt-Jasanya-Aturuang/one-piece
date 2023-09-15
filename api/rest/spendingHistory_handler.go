package rest

import (
	"net/http"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type SpendingHistoryHandlerImpl struct {
	spendingHistoryUsecase domain.SpendingHistoryUsecase
}

func NewSpendingHistoryHandlerImpl(
	spendingHistoryUsecase domain.SpendingHistoryUsecase,
) *SpendingHistoryHandlerImpl {
	return &SpendingHistoryHandlerImpl{
		spendingHistoryUsecase: spendingHistoryUsecase,
	}
}

func (h *SpendingHistoryHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingHistoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingHistoryHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingHistoryHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingHistoryHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}
