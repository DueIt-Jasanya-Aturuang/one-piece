package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load file .env | err : %v", err)
	}

	PgHost = os.Getenv("POSTGRESQL_HOST")
	PgPort = os.Getenv("POSTGRESQL_PORT")
	PgUser = os.Getenv("POSTGRESQL_USER")
	PgPass = os.Getenv("POSTGRESQL_PASS")
	PgDB = os.Getenv("POSTGRESQL_NAME")
	PgSchema = os.Getenv("POSTGRESQL_SCHEMA")
	PgSsl = os.Getenv("POSTGRESQL_SSL")
}

var (
	PgHost   string
	PgPort   string
	PgUser   string
	PgPass   string
	PgDB     string
	PgSchema string
	PgSsl    string
)
