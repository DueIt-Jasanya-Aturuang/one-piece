package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/infrastructures/db"
	dbImpl "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/infrastructures/db/dbImpl"
	convertErentity "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/helpers/converter-entity"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/helpers/minio"
	httpProtocol "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/http-protocol"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/logs"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/handlers"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/repositories"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/services"
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
	validation := validator.New()
	validation.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	paymentRepo := repositories.NewRepositoryImpl()
	minio := minio.NewMinioImpl()
	paymentService := services.NewPaymentServiceImpl(paymentRepo, dbImpl, converter, validation, minio)

	httpHandler := handlers.NewHttpHandler(paymentService)
	http := httpProtocol.NewHttpImpl(httpHandler)
	http.Listen()
}
