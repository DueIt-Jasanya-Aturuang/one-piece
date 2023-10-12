package rapi

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	middlewareChi "github.com/go-chi/chi/v5/middleware"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/middleware"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type Presenter struct {
	balanceUsecase         usecase.BalanceUsecase
	incomeHistoryUsecase   usecase.IncomeHistoryUsecase
	incomeTypeUsecase      usecase.IncomeTypeUsecase
	paymentUsecase         usecase.PaymentUsecase
	spendingHistoryUsecase usecase.SpendingHistoryUsecase
	spendingTypeUsecase    usecase.SpendingTypeUsecase
}

type Dependency struct {
	BalanceUsecase         usecase.BalanceUsecase
	IncomeHistoryUsecase   usecase.IncomeHistoryUsecase
	IncomeTypeUsecase      usecase.IncomeTypeUsecase
	PaymentUsecase         usecase.PaymentUsecase
	SpendingHistoryUsecase usecase.SpendingHistoryUsecase
	SpendingTypeUsecase    usecase.SpendingTypeUsecase
}

type PresenterConfig struct {
	Dependency *Dependency
}

func NewPresenter(config PresenterConfig) (*http.Server, error) {
	presenter := &Presenter{
		balanceUsecase:         config.Dependency.BalanceUsecase,
		incomeHistoryUsecase:   config.Dependency.IncomeHistoryUsecase,
		incomeTypeUsecase:      config.Dependency.IncomeTypeUsecase,
		paymentUsecase:         config.Dependency.PaymentUsecase,
		spendingHistoryUsecase: config.Dependency.SpendingHistoryUsecase,
		spendingTypeUsecase:    config.Dependency.SpendingTypeUsecase,
	}

	r := chi.NewRouter()
	r.Use(middlewareChi.Logger)
	r.Use(middlewareChi.Recoverer)
	r.Use(middleware.CheckApiKey)
	r.MethodNotAllowed(helper.MethodNotAllowed)

	r.Route("/finance", func(r chi.Router) {
		r.Use(middleware.SetAuthorization)

		r.Get("/payment", presenter.GetAllPayment)
		r.Post("/payment", presenter.CreatePayment)
		r.Put("/payment/{id}", presenter.UpdatePayment)
		r.Delete("/payment/{id}", presenter.DeletePayment)

		r.Get("/spending-type", presenter.GetAllSpendingTypeByProfileID)
		r.Get("/spending-type/{periode}", presenter.GetAllSpendingTypeByPeriodeAndProfileID)
		r.Get("/spending-type/detail/{id}", presenter.GetSpendingTypeByIDAndProfileID)
		r.Post("/spending-type", presenter.CreateSpendingType)
		r.Put("/spending-type/{id}", presenter.UpdateSpendingType)
		r.Delete("/spending-type/{id}", presenter.DeleteSpendingType)

		r.Get("/spending-history", presenter.GetAllSpendingHistoryByProfileID)
		r.Get("/spending-history/{id}", presenter.GetSpendingHistoryByIDAndProfileID)
		r.Post("/spending-history", presenter.CreateSpendingHistory)
		r.Put("/spending-history/{id}", presenter.UpdateSpendingHistory)
		r.Delete("/spending-history/{id}", presenter.DeleteSpendingHistory)

		r.Get("/income-type", presenter.GetAllIncomeTypeByProfileID)
		r.Get("/income-type/detail/{id}", presenter.GetIncomeTypeByIDAndProfileID)
		r.Post("/income-type", presenter.CreateIncomeType)
		r.Put("/income-type/{id}", presenter.UpdateIncomeType)
		r.Delete("/income-type/{id}", presenter.DeleteIncomeType)

		r.Get("/income-history", presenter.GetAllIncomeHistoryByProfileID)
		r.Get("/income-history/{id}", presenter.GetIncomeHistoryByIDAndProfileID)
		r.Post("/income-history", presenter.CreateIncomeHistory)
		r.Put("/income-history/{id}", presenter.UpdateIncomeHistory)
		r.Delete("/income-history/{id}", presenter.DeleteIncomeHistory)

		r.Get("/balance", presenter.GetBalanceByProfileID)
	})

	server := &http.Server{
		Addr:              infra.AppAddr,
		Handler:           r,
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
	}

	return server, nil
}
