package repositories

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/entities"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, tx *sql.Tx, entity entities.Payment) (*entities.Payment, error)
	UpdatePayment(ctx context.Context, tx *sql.Tx, entity entities.Payment) (*entities.Payment, error)
	GetPaymentById(ctx context.Context, db *sql.DB, id string) (*entities.Payment, error)
}

type PaymentRepositoryImpl struct{}

func NewPaymentRepositoryImpl() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (repo *PaymentRepositoryImpl) CreatePayment(ctx context.Context, tx *sql.Tx, entity entities.Payment) (*entities.Payment, error) {
	_, err := tx.Exec(`set search_path='dueit'`)
	if err != nil {
		return nil, err
	}

	SQL := "INSERT INTO m_payment_methods (id, name, description, image, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = tx.ExecContext(ctx, SQL, entity.Id, entity.Name, entity.Description, entity.Image, entity.CreatedAt, entity.CreatedBy, entity.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (repo *PaymentRepositoryImpl) UpdatePayment(ctx context.Context, tx *sql.Tx, entity entities.Payment) (*entities.Payment, error) {
	_, err := tx.Exec(`set search_path='dueit'`)
	if err != nil {
		return nil, err
	}

	SQL := "UPDATE m_payment_methods SET name = $1, description = $2, image = $3, updated_at = $4, updated_by = $5 WHERE id = $6"
	_, err = tx.ExecContext(ctx, SQL, entity.Name, entity.Description, entity.Image, entity.UpdatedBy, entity.UpdatedAt, entity.Id)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *PaymentRepositoryImpl) GetPaymentById(ctx context.Context, db *sql.DB, id string) (*entities.Payment, error) {
	_, err := db.Exec(`set search_path='dueit'`)
	if err != nil {
		return nil, err
	}

	SQL := "SELECT id, name, description, image, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM m_payment_methods WHERE id = $1 LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, nil
	}

	var payment entities.Payment

	row.Scan(
		&payment.Id,
		&payment.Name,
		&payment.Description,
		&payment.CreatedAt,
		&payment.CreatedBy,
		&payment.UpdatedAt,
		&payment.UpdatedBy,
		&payment.DeletedAt,
		&payment.DeletedBy,
	)
	return &payment, nil
}
