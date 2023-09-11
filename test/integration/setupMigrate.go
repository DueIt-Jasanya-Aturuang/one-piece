package integration

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog/log"
)

func Migrate(db *sql.DB) {
	log.Info().Msg("start migrate")
	_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS dueit;")
	if err != nil {
		log.Err(err).Msgf("Failed to create schema dueit: %s", err)
		os.Exit(1)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS m_payment_methods (
    id 			VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    name        VARCHAR(128),
    description TEXT,
    image       VARCHAR(255) default 'default-icon.png',
    created_at  DECIMAL     NOT NULL,
    created_by  VARCHAR(64),
    updated_at  DECIMAL     NOT NULL,
    updated_by  VARCHAR(64),
    deleted_at  DECIMAL,
    deleted_by  VARCHAR(64)
    )`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_payment_methods: %s", err)
		os.Exit(1)
	}
}
