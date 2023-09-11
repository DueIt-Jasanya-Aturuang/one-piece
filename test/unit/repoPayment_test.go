package unit

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/_repository"
)

func TestRepoCreatePayment(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)

	query := regexp.QuoteMeta(`INSERT INTO m_payment_methods (id, name, description, image, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`)

	payment := domain.Payment{
		ID:          "paymentID",
		Name:        "bca",
		Description: sql.NullString{},
		Image:       "bca.png",
		CreatedAt:   0,
		CreatedBy:   "userID",
		UpdatedAt:   0,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	t.Run("SUCCESS_commit", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			payment.ID,
			payment.Name,
			payment.Description,
			payment.Image,
			payment.CreatedAt,
			payment.CreatedBy,
			payment.UpdatedAt,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		err = paymentRepo.StartTx(context.TODO(), &sql.TxOptions{}, func() error {
			err = paymentRepo.CreatePayment(context.TODO(), &payment)
			assert.NoError(t, err)
			return err
		})
		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_rollback", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			payment.ID,
			payment.Name,
			payment.Description,
			payment.Image,
			payment.CreatedAt,
			payment.CreatedBy,
			payment.UpdatedAt,
		).WillReturnError(errors.New("go rollback"))
		mocksql.ExpectRollback()

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		err = paymentRepo.StartTx(context.TODO(), &sql.TxOptions{}, func() error {
			err = paymentRepo.CreatePayment(context.TODO(), &payment)
			assert.Error(t, err)
			return err
		})
		assert.Error(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestRepoUpdatePayment(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)

	query := regexp.QuoteMeta(`UPDATE m_payment_methods SET name = $1, description = $2, image = $3, updated_at = $4, updated_by = $5 
            	WHERE id = $6`)

	payment := domain.Payment{
		ID:          "paymentID",
		Name:        "bca",
		Description: sql.NullString{},
		Image:       "bca.png",
		CreatedAt:   0,
		CreatedBy:   "userID",
		UpdatedAt:   0,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	t.Run("SUCCESS_commit", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			payment.Name,
			payment.Description,
			payment.Image,
			payment.UpdatedAt,
			payment.UpdatedBy,
			payment.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		err = paymentRepo.StartTx(context.TODO(), &sql.TxOptions{}, func() error {
			err = paymentRepo.UpdatePayment(context.TODO(), &payment)
			assert.NoError(t, err)
			return err
		})
		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_rollback", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			payment.Name,
			payment.Description,
			payment.Image,
			payment.UpdatedAt,
			payment.UpdatedBy,
			payment.ID,
		).WillReturnError(errors.New("go rollback"))
		mocksql.ExpectRollback()

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		err = paymentRepo.StartTx(context.TODO(), &sql.TxOptions{}, func() error {
			err = paymentRepo.UpdatePayment(context.TODO(), &payment)
			assert.Error(t, err)
			return err
		})
		assert.Error(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestRepoGetPaymentById(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)

	query := regexp.QuoteMeta(`SELECT id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_payment_methods WHERE id = $1`)

	payment := domain.Payment{
		ID:          "paymentID",
		Name:        "bca",
		Description: sql.NullString{},
		Image:       "bca.png",
		CreatedAt:   0,
		CreatedBy:   "userID",
		UpdatedAt:   0,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "description", "image", "created_at", "created_by",
		"updated_at", "updated_by", "deleted_at", "deleted_by"})

	t.Run("SUCCESS", func(t *testing.T) {
		rows = rows.AddRow(payment.ID, payment.Name, payment.Description, payment.Image, payment.CreatedAt, payment.CreatedBy, payment.UpdatedAt, payment.UpdatedBy, payment.DeletedAt, payment.DeletedBy)
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			payment.ID,
		).WillReturnRows(rows)

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		paymentRes, err := paymentRepo.GetPaymentByID(context.TODO(), payment.ID)
		assert.NoError(t, err)
		assert.NotNil(t, paymentRes)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_not-found", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			"nil",
		).WillReturnError(sql.ErrNoRows)

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		paymentResp, err := paymentRepo.GetPaymentByID(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Nil(t, paymentResp)
		assert.Equal(t, sql.ErrNoRows, err)
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestRepoGetPaymentByName(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)

	query := regexp.QuoteMeta(`SELECT id, name, description, image, created_at, created_by, 
       				updated_at, updated_by, deleted_at, deleted_by 
			 FROM m_payment_methods WHERE name = $1`)

	payment := domain.Payment{
		ID:          "paymentID",
		Name:        "bca",
		Description: sql.NullString{},
		Image:       "bca.png",
		CreatedAt:   0,
		CreatedBy:   "userID",
		UpdatedAt:   0,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "description", "image", "created_at", "created_by",
		"updated_at", "updated_by", "deleted_at", "deleted_by"})

	t.Run("SUCCESS", func(t *testing.T) {
		rows = rows.AddRow(payment.ID, payment.Name, payment.Description, payment.Image, payment.CreatedAt, payment.CreatedBy, payment.UpdatedAt, payment.UpdatedBy, payment.DeletedAt, payment.DeletedBy)
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			payment.Name,
		).WillReturnRows(rows)

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		paymentRes, err := paymentRepo.GetPaymentByName(context.TODO(), payment.Name)
		assert.NoError(t, err)
		assert.NotNil(t, paymentRes)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_not-found", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			"nil",
		).WillReturnError(sql.ErrNoRows)

		err = paymentRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer paymentRepo.CloseConn()

		paymentResp, err := paymentRepo.GetPaymentByName(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Nil(t, paymentResp)
		assert.Equal(t, sql.ErrNoRows, err)
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
