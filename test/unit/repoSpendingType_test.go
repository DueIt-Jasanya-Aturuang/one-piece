package unit

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

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

	query := regexp.QuoteMeta(`INSERT INTO m_spending_type (id, profile_id, title, maximum_limit, icon, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	spendingType := &domain.SpendingType{
		ID:           "test",
		ProfileID:    "test",
		Title:        "test",
		MaximumLimit: 123,
		Icon:         "test",
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "test",
			UpdatedAt: 0,
		},
	}
	t.Run("SUCCESS", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			"test", "test", "test", 123, "test", 0, "test", 0,
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

	query := regexp.QuoteMeta(`UPDATE m_spending_type SET title = $1, maximum_limit = $2, icon = $3, updated_at = $4, updated_by = $5 
                    WHERE id = $6 AND profile_id = $7 AND deleted_at IS NULL`)
	spendingType := &domain.SpendingType{
		ID:           "test",
		ProfileID:    "test",
		Title:        "test",
		MaximumLimit: 123,
		Icon:         "test",
		AuditInfo: domain.AuditInfo{
			UpdatedAt: 0,
			UpdatedBy: sql.NullString{String: "test", Valid: true},
		},
	}
	t.Run("SUCCESS", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			"test", 123, "test", 0, "test", "test", "test",
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
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingTypeRepo := _repository.NewSpendingTypeRepositoryImpl(uow)

	query := regexp.QuoteMeta(`UPDATE m_spending_type SET deleted_by = $1, deleted_at = $2 WHERE id = $3 AND profile_id = $4`)

	t.Run("SUCCESS", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			"test", time.Now().Unix(), "test", "test",
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		err = spendingTypeRepo.StartTx(context.TODO(), &sql.TxOptions{}, func() error {
			err = spendingTypeRepo.Delete(context.TODO(), "test", "test")
			assert.NoError(t, err)
			return nil
		})
		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestRepoSpendingTypeGetByID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingTypeRepo := _repository.NewSpendingTypeRepositoryImpl(uow)

	query := regexp.QuoteMeta(`SELECT id, profile_id, title, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_spending_type WHERE id = $1 AND deleted_at IS NULL`)
	rows := sqlmock.NewRows([]string{"id", "profile_id", "title", "maximum_limit", "icon", "created_at", "created_by",
		"updated_at", "updated_by", "deleted_at", "deleted_by"})

	t.Run("SUCCESS", func(t *testing.T) {
		rows.AddRow("test", "test", "test", 123, "test", 0, "test", 0, nil, nil, nil)
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs("test").WillReturnRows(rows)

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		spendingType, err := spendingTypeRepo.GetByID(context.TODO(), "test")
		assert.NoError(t, err)
		assert.NotNil(t, spendingType)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs("test").WillReturnError(sql.ErrNoRows)

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		spendingType, err := spendingTypeRepo.GetByID(context.TODO(), "test")
		assert.Error(t, err)
		assert.Nil(t, spendingType)
		assert.Equal(t, sql.ErrNoRows, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestRepoSpendingTypeGetByIDAndProfileID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingTypeRepo := _repository.NewSpendingTypeRepositoryImpl(uow)

	query := regexp.QuoteMeta(`SELECT id, profile_id, title, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_spending_type WHERE id = $1 AND profile_id = $2 AND deleted_at IS NULL`)
	rows := sqlmock.NewRows([]string{"id", "profile_id", "title", "maximum_limit", "icon", "created_at", "created_by",
		"updated_at", "updated_by", "deleted_at", "deleted_by"})

	t.Run("SUCCESS", func(t *testing.T) {
		rows.AddRow("test", "test", "test", 123, "test", 0, "test", 0, nil, nil, nil)
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs("test", "test").WillReturnRows(rows)

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		spendingType, err := spendingTypeRepo.GetByIDAndProfileID(context.TODO(), "test", "test")
		assert.NoError(t, err)
		assert.NotNil(t, spendingType)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs("test", "test").WillReturnError(sql.ErrNoRows)

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		spendingType, err := spendingTypeRepo.GetByIDAndProfileID(context.TODO(), "test", "test")
		assert.Error(t, err)
		assert.Nil(t, spendingType)
		assert.Equal(t, sql.ErrNoRows, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestRepoSpendingTypeGetAllByProfileID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingTypeRepo := _repository.NewSpendingTypeRepositoryImpl(uow)

	query := regexp.QuoteMeta(`SELECT mst.id, mst.profile_id, mst.title, mst.maximum_limit, mst.icon, mst.created_at, mst.created_by, 
       				mst.updated_at, mst.updated_by, mst.deleted_at, mst.deleted_by, 
       				COALESCE(SUM(CASE WHEN tsh.time_spending_history BETWEEN $2 AND $3 AND tsh.deleted_at IS NULL THEN tsh.spending_amount ELSE 0 END), 0)
				FROM m_spending_type mst
				LEFT JOIN t_spending_history tsh ON mst.id = tsh.spending_type_id
				WHERE mst.profile_id = $1 AND tsh.deleted_at IS NULL AND mst.deleted_at IS NULL
				GROUP BY mst.id`)
	rows := sqlmock.NewRows([]string{"id", "profile_id", "title", "maximum_limit", "icon", "created_at", "created_by",
		"updated_at", "updated_by", "deleted_at", "deleted_by", "used"})

	t.Run("SUCCESS", func(t *testing.T) {
		startPeriod := time.Date(time.Now().Year(), time.Now().Month(), 14, 0, 0, 0, 0, time.UTC)
		endPeriod := startPeriod.AddDate(0, 1, 0)
		rows.AddRow("test", "test", "test", 123, "test", 0, "test", 0, nil, nil, nil, 0)
		rows.AddRow("test1", "test", "test1", 123, "test", 0, "test1", 0, nil, nil, nil, 0)
		rows.AddRow("test2", "test", "test2", 123, "test", 0, "test2", 0, nil, nil, nil, 0)
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs("test", startPeriod, endPeriod).WillReturnRows(rows)

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		spendingType, err := spendingTypeRepo.GetAllByProfileID(context.TODO(), "test", startPeriod, endPeriod)
		assert.NoError(t, err)
		assert.NotNil(t, spendingType)
		assert.Equal(t, 3, len(*spendingType))
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_nil", func(t *testing.T) {
		startPeriod := time.Date(time.Now().Year(), time.Now().Month(), 14, 0, 0, 0, 0, time.UTC)
		endPeriod := startPeriod.AddDate(0, 1, 0)
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs("test", startPeriod, endPeriod).WillReturnRows(rows)

		err := spendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		spendingType, err := spendingTypeRepo.GetAllByProfileID(context.TODO(), "test", startPeriod, endPeriod)
		assert.NoError(t, err)
		assert.NotNil(t, spendingType)
		assert.Equal(t, 0, len(*spendingType))

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
