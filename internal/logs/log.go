package logs

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	ConsoleLoggingEnabled bool
	EncodeLogsAsJson      bool
	FileLoggingEnabled    bool
	Directory             string
	Filename              string
	MaxSize               int
	MaxBackups            int
	MaxAge                int
}

func InitLogger(cfg Config) {
	var writers []io.Writer

	if cfg.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if cfg.FileLoggingEnabled {
		writers = append(writers, newRollingFile(cfg))
	}

	mw := io.MultiWriter(writers...)

	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("message : %s ", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("| %s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()
	log.Logger = logger

	logger.Output(output)
	logger.Info().
		Bool("fileLogging", cfg.FileLoggingEnabled).
		Bool("jsonLogOutput", cfg.EncodeLogsAsJson).
		Str("logDirectory", cfg.Directory).
		Str("fileName", cfg.Filename).
		Int("maxSizeMB", cfg.MaxSize).
		Int("maxBackups", cfg.MaxBackups).
		Int("maxAgeInDays", cfg.MaxAge).
		Msg("logging configured")
}

func newRollingFile(cfg Config) io.Writer {
	if err := os.MkdirAll(cfg.Directory, 0o744); err != nil {
		log.Error().Err(err).Str("path", cfg.Directory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(cfg.Directory, cfg.Filename),
		MaxBackups: cfg.MaxBackups, // files
		MaxSize:    cfg.MaxSize,    // megabytes
		MaxAge:     cfg.MaxAge,     // days
	}
}
