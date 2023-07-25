package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/config"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewPostgresConnection() *sql.DB {
	dbHost := config.Get().DB.Postgres.Host
	dbPort := config.Get().DB.Postgres.Port
	dbUser := config.Get().DB.Postgres.User
	dbPass := config.Get().DB.Postgres.Pass
	dbName := config.Get().DB.Postgres.Name

	fDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", fDB)
	if err != nil {
		log.Err(err).Dict("errors", zerolog.Dict().Str("file", "postgres.go").Str("line", "23")).Msg("error load to connect db")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	log.Info().Msgf("postgres started on dbname : %s", dbName)
	return db
}
