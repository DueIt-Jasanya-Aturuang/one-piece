package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresConn() *sql.DB {
	fDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		PgHost, PgPort, PgUser, PgPass, PgDB, PgSsl, PgSchema,
	)

	db, err := sql.Open("postgres", fDB)
	if err != nil {
		log.Fatalf("failed open connection to %s | err %v", fDB, err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), pgPingTimeOut)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("failed ping connection to %s | err %v", fDB, err)
	}

	db.SetMaxIdleConns(setMaxIdleConnDB)
	db.SetMaxOpenConns(setMaxOpenConnDB)
	db.SetConnMaxIdleTime(SetConnMaxIdleTimeDB)
	db.SetConnMaxLifetime(setConnMaxLifetimeDB)

	log.Printf("successfully open connection to %s", PgDB)
	return db
}

const (
	setMaxIdleConnDB     = 5
	setMaxOpenConnDB     = 100
	SetConnMaxIdleTimeDB = 5 * time.Minute
	setConnMaxLifetimeDB = 60 * time.Minute
	pgPingTimeOut        = 2 * time.Second
)
