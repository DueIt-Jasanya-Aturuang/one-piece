package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	middlewareChi "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/middleware"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func main() {
	infra.LogInit()
	infra.EnvInit()

	pgConn := infra.NewPostgresConn()
	minioConn := infra.NewMinioConn()

	// repository
	uow := repository.NewUnitOfWorkRepositoryImpl(pgConn)
	paymentRepo := repository.NewPaymentRepositoryImpl(uow)
	minioRepo := repository.NewMinioImpl(minioConn)
	spendingTypeRepo := repository.NewSpendingTypeRepositoryImpl(uow)
	spendingHistoryRepo := repository.NewSpendingHistoryRepositoryImpl(uow)
	balanceRepo := repository.NewBalanceRepositoryImpl(uow)
	incomeTypeRepo := repository.NewIncomeTypeRepositoryImpl(uow)
	incomeHistoryRepo := repository.NewIncomeHistoryRepositoryImpl(uow)

	// usecase
	paymentUsecase := usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)
	spendingTypeUsecase := usecase.NewSpendingTypeUsecaseImpl(spendingTypeRepo)
	spendingHistoryUsecase := usecase.NewSpendingHistoryUsecaseImpl(spendingHistoryRepo, spendingTypeRepo, balanceRepo, paymentRepo)
	balanceUsecase := usecase.NewBalanceUsecaseImpl(balanceRepo)
	incomeTypeUsecase := usecase.NewIncomeTypeUsecaseImpl(incomeTypeRepo)
	incomeHistoryUsecase := usecase.NewIncomeHistoryUsecaseImpl(incomeTypeRepo, incomeHistoryRepo, paymentRepo, balanceRepo)

	// handler
	paymentHandler := rapi.NewPaymentHandlerImpl(paymentUsecase)
	spendingTypeHandler := rapi.NewSpendingTypeHandlerImpl(spendingTypeUsecase)
	spendingHistoryHandler := rapi.NewSpendingHistoryHandlerImpl(spendingHistoryUsecase)
	balanceHandler := rapi.NewBalanceHandlerImpl(balanceUsecase)
	incomeTypeHandler := rapi.NewIncomeTypeHandlerImpl(incomeTypeUsecase)
	incomeHistoryHandler := rapi.NewIncomeHistoryHandlerImpl(incomeHistoryUsecase)

	// route
	r := chi.NewRouter()
	r.Use(middlewareChi.Logger)
	r.Use(middlewareChi.Recoverer)
	r.Use(middleware.CheckApiKey)
	r.MethodNotAllowed(helper.MethodNotAllowed)

	r.Route("/finance", func(r chi.Router) {
		r.Use(middleware.SetAuthorization)

		r.Get("/payment", paymentHandler.GetAll)
		r.Post("/payment", paymentHandler.Create)
		r.Put("/payment/{id}", paymentHandler.Update)
		r.Delete("/payment/{id}", paymentHandler.Delete)

		r.Get("/spending-type", spendingTypeHandler.GetAllByProfileID)
		r.Get("/spending-type/{periode}", spendingTypeHandler.GetAllByPeriodeAndProfileID)
		r.Get("/spending-type/detail/{id}", spendingTypeHandler.GetByIDAndProfileID)
		r.Post("/spending-type", spendingTypeHandler.Create)
		r.Put("/spending-type/{id}", spendingTypeHandler.Update)
		r.Delete("/spending-type/{id}", spendingTypeHandler.Delete)

		r.Get("/spending-history", spendingHistoryHandler.GetAllByProfileID)
		r.Get("/spending-history/{id}", spendingHistoryHandler.GetByIDAndProfileID)
		r.Post("/spending-history", spendingHistoryHandler.Create)
		r.Put("/spending-history/{id}", spendingHistoryHandler.Update)
		r.Delete("/spending-history/{id}", spendingHistoryHandler.Delete)

		r.Get("/income-type", incomeTypeHandler.GetAllByProfileID)
		r.Get("/income-type/detail/{id}", incomeTypeHandler.GetByIDAndProfileID)
		r.Post("/income-type", incomeTypeHandler.Create)
		r.Put("/income-type/{id}", incomeTypeHandler.Update)
		r.Delete("/income-type/{id}", incomeTypeHandler.Delete)

		r.Get("/income-history", incomeHistoryHandler.GetAllByProfileID)
		r.Get("/income-history/{id}", incomeHistoryHandler.GetByIDAndProfileID)
		r.Post("/income-history", incomeHistoryHandler.Create)
		r.Put("/income-history/{id}", incomeHistoryHandler.Update)
		r.Delete("/income-history/{id}", incomeHistoryHandler.Delete)

		r.Get("/balance", balanceHandler.GetByProfileID)
	})

	log.Info().Msgf("Server is running on port %s", infra.AppAddr)
	err := http.ListenAndServe(infra.AppAddr, r)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed listen server on %s", infra.AppAddr)
	}

}
