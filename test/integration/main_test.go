package integration

import (
	"database/sql"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/_repository"
)

var DB *sql.DB
var Uow = _repository.NewUnitOfWorkRepositoryImpl(DB)
var PaymentRepo = _repository.NewPaymentRepositoryImpl(Uow)

func TestMain(m *testing.M) {
	dockerpool := SetupDocker()
	var resources []*dockertest.Resource

	pgResource, dbPg, _ := Postgres(dockerpool)
	resources = append(resources, pgResource)
	DB = dbPg
	Uow = _repository.NewUnitOfWorkRepositoryImpl(DB)
	PaymentRepo = _repository.NewPaymentRepositoryImpl(Uow)
	if DB == nil {
		panic("db nil")
	}

	Migrate(DB)

	code := m.Run()

	for _, resource := range resources {
		if err := dockerpool.Purge(resource); err != nil {
			log.Fatal().Msgf("failed purge resource")
		}
	}
	os.Exit(code)
}

func TestInit(t *testing.T) {
	t.Run("PAYMENT_REPO", func(t *testing.T) {
		t.Run("CreatePayment", CreatePayment)
		t.Run("GetPaymentByID", GetPaymentById)
		t.Run("GetPaymentById_ERROR", GetPaymentByIdERROR)
		t.Run("UpdatePayment", UpdatePayment)
		t.Run("GetPaymentByName", GetPaymentByName)
		t.Run("GetPaymentByName_ERROR", GetPaymentByNameERROR)
	})
}
