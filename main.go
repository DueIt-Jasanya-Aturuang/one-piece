package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/config"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_usecase"
)

func main() {
	config.LogInit()
	config.EnvInit()

	pgConn := config.NewPostgresConn()
	minioConn := config.NewMinioConn()

	// repository
	uow := _repository.NewUnitOfWorkRepositoryImpl(pgConn)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioConn)

	// usecase
	paymentUsecase := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	// handler
	paymentHandler := rest.NewPaymentHandlerImpl(paymentUsecase)

	// route
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.MethodNotAllowed(helper.MethodNotAllowed)

	r.Get("/finance/payment", paymentHandler.GetAllPayment)
	r.Post("/finance/payment", paymentHandler.CreatePayment)
	r.Put("/finance/payment/{id}", paymentHandler.UpdatePayment)

	log.Info().Msgf("Server is running on port %s", config.AppAddr)
	err := http.ListenAndServe(config.AppAddr, r)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed listen server on %s", config.AppAddr)
	}

}
