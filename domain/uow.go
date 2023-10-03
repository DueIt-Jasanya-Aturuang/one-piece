package domain

import (
	"context"
	"database/sql"
)

type AuditInfo struct {
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy sql.NullString
	DeletedAt sql.NullInt64
	DeletedBy sql.NullString
}

type RequestGetAllPaginate struct {
	ProfileID string
	ID        string
	Operation string
	Order     string
}

//counterfeiter:generate -o ./mocks . UnitOfWorkRepository
type UnitOfWorkRepository interface {
	OpenConn(ctx context.Context) error
	GetConn() (*sql.Conn, error)
	CloseConn()
	StartTx(ctx context.Context, opts *sql.TxOptions, fn func() error) error
	GetTx() (*sql.Tx, error)
}
