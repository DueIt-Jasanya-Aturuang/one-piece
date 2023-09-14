package unit

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_repository"
)

func TestRepoSpendingTypeCreate(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingTypeRepo := _repository.NewSpendingTypeRepositoryImpl(uow)

	query := regexp.QuoteMeta(`INSERT INTO m_spending_type (id, profile_id, title, maximum_limit, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	spendingType := &domain.SpendingType{
		ID:           "test",
		ProfileID:    "test",
		Title:        "test",
		MaximumLimit: 123,
		CreatedAt:    0,
		CreatedBy:    "test",
		UpdatedAt:    0,
	}
	t.Run("SUCCESS", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			"test", "test", "test", 123, 0, "test", 0,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		err = spendingTypeRepo.StartTx(context.TODO(), &sql.TxOptions{}, func() error {
			err = spendingTypeRepo.Create(context.TODO(), spendingType)
			assert.NoError(t, err)
			return nil
		})
		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestRepoSpendingTypeUpdate(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingTypeRepo := _repository.NewSpendingTypeRepositoryImpl(uow)

	query := regexp.QuoteMeta(`UPDATE m_spending_type SET title = $1, maximum_limit = $2, updated_at = $3, updated_by = $4 
                    WHERE id = $5 AND profile_id = $6 AND deleted_at IS NULL`)
	spendingType := &domain.SpendingType{
		ID:           "test",
		ProfileID:    "test",
		Title:        "test",
		MaximumLimit: 123,
		UpdatedAt:    0,
		UpdatedBy:    sql.NullString{String: "test", Valid: true},
	}
	t.Run("SUCCESS", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			"test", 123, 0, "test", "test", "test",
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		err = spendingTypeRepo.StartTx(context.TODO(), &sql.TxOptions{}, func() error {
			err = spendingTypeRepo.Update(context.TODO(), spendingType)
			assert.NoError(t, err)
			return nil
		})
		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

}

func TestRepoSpendingTypeDelete(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)

	mocksql.ExpectPrepare("")
	defer db.Close()
}

func TestRepoSpendingTypeGetByID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)

	mocksql.ExpectPrepare("")
	defer db.Close()
}

func TestRepoSpendingTypeGetByIDAndProfileID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)

	mocksql.ExpectPrepare("")
	defer db.Close()
}

func TestRepoSpendingTypeGetAllByProfileID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)

	mocksql.ExpectPrepare("")
	defer db.Close()
}
