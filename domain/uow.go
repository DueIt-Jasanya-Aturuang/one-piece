package domain

import (
	"context"
	"database/sql"
)

//counterfeiter:generate -o ./mocks . UnitOfWorkRepository
type UnitOfWorkRepository interface {
	OpenConn(ctx context.Context) error
	GetConn() (*sql.Conn, error)
	CloseConn()
	StartTx(ctx context.Context, opts *sql.TxOptions, fn func() error) error
	GetTx() (*sql.Tx, error)
}
