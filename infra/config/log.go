package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogInit() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logFile, err := os.Create("app.log")
	if err != nil {
		log.Fatal().Msgf("failed create file log | err %v", err)
	}

	multi := zerolog.MultiLevelWriter(logFile, os.Stdout)
	log.Logger = zerolog.New(multi).With().Caller().Timestamp().Logger()

	log.Info().Msgf("successfully init logger")
}
