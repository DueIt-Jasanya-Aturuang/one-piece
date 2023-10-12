package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/balance_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/incomeHistory_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/incomeType_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/minio_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/payment_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/spendingHistory_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/spendingType_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository/uow_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/balance_usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/incomeHistory_usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/incomeType_usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/payment_usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/spendingHistory_usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/spendingType_usecase"
)

func main() {
	infra.LogInit()
	infra.EnvInit()

	pgConn := infra.NewPostgresConn()
	minioConn := infra.NewMinioConn()

	depen := dependency(pgConn, minioConn)

	httpServer, err := rapi.NewPresenter(rapi.PresenterConfig{
		Dependency: depen,
	})
	if err != nil {
		log.Fatal().Msgf("creating new presenter: %s", err.Error())
	}

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	go func() {
		<-exitSignal
		log.Info().Msgf("Interrupt signal recivied, existing...")

		shudownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Minute)
		defer shutdownCancel()

		err = httpServer.Shutdown(shudownCtx)
		if err != nil {
			log.Err(err).Msg("shutting down HTTP server")
		}
	}()
	log.Info().Msgf("Server is running on port %s", infra.AppAddr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed listen server on %s", infra.AppAddr)
	}
}

func dependency(db *sql.DB, minioClient *minio.Client) *rapi.Dependency {
	uow := uow_repository.NewUnitOfWorkRepositoryImpl(db)
	paymentRepo := payment_repository.NewPaymentRepositoryImpl(uow)
	minioRepo := minio_repository.NewMinioImpl(minioClient)
	spendingTypeRepo := spendingType_repository.NewSpendingTypeRepositoryImpl(uow)
	spendingHistoryRepo := spendingHistory_repository.NewSpendingHistoryRepositoryImpl(uow)
	balanceRepo := balance_repository.NewBalanceRepositoryImpl(uow)
	incomeTypeRepo := incomeType_repository.NewIncomeTypeRepositoryImpl(uow)
	incomeHistoryRepo := incomeHistory_repository.NewIncomeHistoryRepositoryImpl(uow)

	// usecase_old
	balanceUsecase := balance_usecase.NewBalanceUsecaseImpl(balanceRepo)
	paymentUsecase := payment_usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)
	spendingTypeUsecase := spendingType_usecase.NewSpendingTypeUsecaseImpl(spendingTypeRepo, spendingHistoryRepo)
	spendingHistoryUsecase := spendingHistory_usecase.NewSpendingHistoryUsecaseImpl(spendingHistoryRepo, spendingTypeRepo, balanceUsecase, paymentRepo)
	incomeTypeUsecase := incomeType_usecase.NewIncomeTypeUsecaseImpl(incomeTypeRepo)
	incomeHistoryUsecase := incomeHistory_usecase.NewIncomeHistoryUsecaseImpl(incomeTypeRepo, incomeHistoryRepo, paymentRepo, balanceUsecase)

	return &rapi.Dependency{
		BalanceUsecase:         balanceUsecase,
		IncomeHistoryUsecase:   incomeHistoryUsecase,
		IncomeTypeUsecase:      incomeTypeUsecase,
		PaymentUsecase:         paymentUsecase,
		SpendingHistoryUsecase: spendingHistoryUsecase,
		SpendingTypeUsecase:    spendingTypeUsecase,
	}
}
