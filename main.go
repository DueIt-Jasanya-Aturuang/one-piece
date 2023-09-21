package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/_usecase"
)

func main() {
	infra.LogInit()
	infra.EnvInit()

	pgConn := infra.NewPostgresConn()
	minioConn := infra.NewMinioConn()

	// repository
	uow := _repository.NewUnitOfWorkRepositoryImpl(pgConn)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioConn)
	spendingTypeRepo := _repository.NewSpendingTypeRepositoryImpl(uow)
	spendingHistoryRepo := _repository.NewSpendingHistoryRepositoryImpl(uow)
	balanceRepo := _repository.NewBalanceRepositoryImpl(uow)

	// usecase
	paymentUsecase := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)
	spendingTypeUsecase := _usecase.NewSpendingTypeUsecaseImpl(spendingTypeRepo)
	spendingHistoryUsecase := _usecase.NewSpendingHistoryUsecaseImpl(spendingHistoryRepo, spendingTypeRepo, balanceRepo, paymentRepo)
	balanceUsecase := _usecase.NewBalanceUsecaseImpl(balanceRepo)

	// handler
	paymentHandler := rest.NewPaymentHandlerImpl(paymentUsecase)
	spendingTypeHandler := rest.NewSpendingTypeHandlerImpl(spendingTypeUsecase)
	spendingHistoryHandler := rest.NewSpendingHistoryHandlerImpl(spendingHistoryUsecase)
	balanceHandler := rest.NewBalanceHandlerImpl(balanceUsecase)

	// route
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.MethodNotAllowed(helper.MethodNotAllowed)

	r.Route("/finance", func(r chi.Router) {
		r.Get("/payment", paymentHandler.GetAll)
		r.Post("/payment", paymentHandler.Create)
		r.Put("/payment/{id}", paymentHandler.Update)

		r.Get("/spending-type/{profile-id}", spendingTypeHandler.GetAllByPeriodeAndProfileID)
		r.Get("/spending-type/{profile-id}/{id}", spendingTypeHandler.GetByIDAndProfileID)
		r.Post("/spending-type", spendingTypeHandler.Create)
		r.Put("/spending-type/{id}", spendingTypeHandler.Update)
		r.Delete("/spending-type/{profile-id}/{id}", spendingTypeHandler.Delete)

		r.Get("/spending-history/{profile-id}", spendingHistoryHandler.GetAllByProfileID)
		r.Get("/spending-history/{profile-id}/{id}", spendingHistoryHandler.GetByIDAndProfileID)
		r.Post("/spending-history", spendingHistoryHandler.Create)
		r.Put("/spending-history/{id}", spendingHistoryHandler.Update)
		r.Delete("/spending-history/{profile-id}/{id}", spendingHistoryHandler.Delete)

		r.Get("/balance/{profile-id}", balanceHandler.GetByProfileID)
	})

	log.Info().Msgf("Server is running on port %s", infra.AppAddr)
	err := http.ListenAndServe(infra.AppAddr, r)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed listen server on %s", infra.AppAddr)
	}

}
