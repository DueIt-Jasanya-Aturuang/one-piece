package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infrastructures/db"
	dbImpl "github.com/DueIt-Jasanya-Aturuang/one-piece/infrastructures/db/dbImpl"
	convertErentity "github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helpers/converter-entity"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helpers/minio"
	httpProtocol "github.com/DueIt-Jasanya-Aturuang/one-piece/internal/http-protocol"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/logs"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/utils/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/src/handlers"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/src/modules/repositories"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/src/modules/services"
	"github.com/go-playground/validator/v10"
)

func main() {
	dir, _ := os.Getwd()
	dir = fmt.Sprintf("%s/internal/logs", dir)
	logs.InitLogger(logs.Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             dir,
		Filename:              "spending.log",
		MaxSize:               200000000,
		MaxBackups:            2000,
		MaxAge:                2000,
	})

	postgresDb := db.NewPostgresConnection()
	dbImpl := dbImpl.NewDbImpl(postgresDb)
	converter := convertErentity.NewConvertImpl()
	validator := validator.New()
	validator.RegisterValidation("unique", validation.MustUnique)
	validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	paymentRepo := repositories.NewRepositoryImpl()
	minio := minio.NewMinioImpl()
	paymentService := services.NewPaymentServiceImpl(paymentRepo, dbImpl, converter, validator, minio)

	httpHandler := handlers.NewHttpHandler(paymentService)
	http := httpProtocol.NewHttpImpl(httpHandler)
	http.Listen()
}
