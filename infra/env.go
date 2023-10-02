package infra

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func EnvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load file .env | err : %v", err)
	}

	// app
	AppAddr = os.Getenv("APPLICATION_ADDR")
	AppApiKeyFinance = os.Getenv("APPLICATION_KEY_FINANCE")

	// pg db
	PgHost = os.Getenv("POSTGRESQL_HOST")
	PgPort = os.Getenv("POSTGRESQL_PORT")
	PgUser = os.Getenv("POSTGRESQL_USER")
	PgPass = os.Getenv("POSTGRESQL_PASS")
	PgDB = os.Getenv("POSTGRESQL_NAME")
	PgSchema = os.Getenv("POSTGRESQL_SCHEMA")
	PgSsl = os.Getenv("POSTGRESQL_SSL")

	// minio
	minioSslBool, err := strconv.ParseBool(os.Getenv("MINIO_SSL"))
	if err != nil {
		panic(err)
	}
	MinIoID = os.Getenv("MINIO_ID")
	MinIoSecretKey = os.Getenv("MINIO_SECRETKEY")
	MinIoEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinIoBucket = os.Getenv("MINIO_BUCKET")
	MinIoSSL = minioSslBool
}

var (
	AppAddr          string
	AppApiKeyFinance string
)
var (
	PgHost   string
	PgPort   string
	PgUser   string
	PgPass   string
	PgDB     string
	PgSchema string
	PgSsl    string
)

var (
	MinIoID        string
	MinIoSecretKey string
	MinIoEndpoint  string
	MinIoBucket    string
	MinIoSSL       bool
)
