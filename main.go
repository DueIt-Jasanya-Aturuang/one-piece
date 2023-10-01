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
	_repository2 "github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	usecase2 "github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func main() {
	infra.LogInit()
	infra.EnvInit()

	pgConn := infra.NewPostgresConn()
	minioConn := infra.NewMinioConn()

	// repository
	uow := _repository2.NewUnitOfWorkRepositoryImpl(pgConn)
	paymentRepo := _repository2.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository2.NewMinioImpl(minioConn)
	spendingTypeRepo := _repository2.NewSpendingTypeRepositoryImpl(uow)
	spendingHistoryRepo := _repository2.NewSpendingHistoryRepositoryImpl(uow)
	balanceRepo := _repository2.NewBalanceRepositoryImpl(uow)
	incomeTypeRepo := _repository2.NewIncomeTypeRepositoryImpl(uow)

	// usecase
	paymentUsecase := usecase2.NewPaymentUsecaseImpl(paymentRepo, minioRepo)
	spendingTypeUsecase := usecase2.NewSpendingTypeUsecaseImpl(spendingTypeRepo)
	spendingHistoryUsecase := usecase2.NewSpendingHistoryUsecaseImpl(spendingHistoryRepo, spendingTypeRepo, balanceRepo, paymentRepo)
	balanceUsecase := usecase2.NewBalanceUsecaseImpl(balanceRepo)
	incomeTypeUsecase := usecase2.NewIncomeTypeUsecaseImpl(incomeTypeRepo)

	// handler
	paymentHandler := rapi.NewPaymentHandlerImpl(paymentUsecase)
	spendingTypeHandler := rapi.NewSpendingTypeHandlerImpl(spendingTypeUsecase)
	spendingHistoryHandler := rapi.NewSpendingHistoryHandlerImpl(spendingHistoryUsecase)
	balanceHandler := rapi.NewBalanceHandlerImpl(balanceUsecase)
	incomeTypeHandler := rapi.NewIncomeTypeHandlerImpl(incomeTypeUsecase)

	// route
	r := chi.NewRouter()
	r.Use(middlewareChi.Logger)
	r.Use(middlewareChi.Recoverer)
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

		r.Get("/balance", balanceHandler.GetByProfileID)
	})

	log.Info().Msgf("Server is running on port %s", infra.AppAddr)
	err := http.ListenAndServe(infra.AppAddr, r)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed listen server on %s", infra.AppAddr)
	}

}
