package dbimpl

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
)

type DbImpl struct {
	DB *sql.DB
}

func NewDbImpl(db *sql.DB) *DbImpl {
	return &DbImpl{
		DB: db,
	}
}

func (dbIMPL *DbImpl) StartTX(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := dbIMPL.DB.BeginTx(ctx, opts)
	return tx, err
}

func (dbIMPL *DbImpl) RunWithTransaction(ctx context.Context, opts *sql.TxOptions, fn func(*sql.Tx) error) error {
	tx, err := dbIMPL.StartTX(ctx, opts)
	if err != nil {
		log.Err(err).Msg("cannot start tx")
		return err
	}

	err = fn(tx)
	if err != nil {
		log.Err(err).Msg("roolback")
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
