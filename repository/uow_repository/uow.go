package uow_repository

import (
	"database/sql"
)

type UnitOfWorkRepositoryImpl struct {
	tx *sql.Tx
	db *sql.DB
}

func NewUnitOfWorkRepositoryImpl(db *sql.DB) *UnitOfWorkRepositoryImpl {
	return &UnitOfWorkRepositoryImpl{
		db: db,
	}
}
