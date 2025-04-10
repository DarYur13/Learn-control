package config

import (
	"log"

	"github.com/DarYur13/learn-control/internal/config/modules"
)

const (
	debug  = "debug"
	info   = "info"
	warn   = "warn"
	error  = "error"
	dPanic = "dpanic"
	panic  = "panic"
	fatal  = "fatal"
)

type config struct {
	Pg            *modules.Pg
	Log           *modules.Log
	Api           *modules.Api
	DocsGenerator *modules.DocsGenerator
	Notifier      *modules.Notifier
}

var globalConfig config

func LoadAll() {
	logger, err := modules.LoadLog()
	if err != nil {
		log.Fatalf("failed to load logger config")
	}

	switch logger.Level {
	case debug, info, warn, error, dPanic, panic, fatal:
	default:
		log.Fatalf("unknown log level in config file: %s", logger.Level)
	}

	db, err := modules.LoadPg()
	if err != nil {
		log.Fatalf("failed to load pg config")
	}

	api, err := modules.LoadApi()
	if err != nil {
		log.Fatalf("failed to load api config")
	}

	docsGenerator, err := modules.LoadDocsGenerator()
	if err != nil {
		log.Fatalf("failed to load docs generator config")
	}

	notifier, err := modules.LoadNotifier()
	if err != nil {
		log.Fatalf("failed to load notifier config")
	}

	globalConfig = config{
		Log:           logger,
		Pg:            db,
		Api:           api,
		DocsGenerator: docsGenerator,
		Notifier:      notifier,
	}
}

func ApiGrpcPort() string {
	return globalConfig.Api.GRPCPort
}

func ApiHttpPort() string {
	return globalConfig.Api.HttpPort
}

func ApiHost() string {
	return globalConfig.Api.Host
}

func LogLevel() string {
	return globalConfig.Log.Level
}

func LogFilePath() string {
	return globalConfig.Log.FilePath
}

func PgPort() string {
	return globalConfig.Pg.Port
}

func PgHost() string {
	return globalConfig.Pg.Host
}

func PgUser() string {
	return globalConfig.Pg.User
}

func PgPassword() string {
	return globalConfig.Pg.Password
}

func PgDatabase() string {
	return globalConfig.Pg.Database
}

func DocsGeneratorTamplatePath() string {
	return globalConfig.DocsGenerator.TamplatePath
}

func NotifierEmailFrom() string {
	return globalConfig.Notifier.EmailFrom
}
func NotifierEmailPassword() string {
	return globalConfig.Notifier.EmailPassword
}
func NotifierSMTPHost() string {
	return globalConfig.Notifier.SMTPHost
}
func NotifierSMTPPort() string {
	return globalConfig.Notifier.SMTPPort
}
func NotifierEmailUseTLS() string {
	return globalConfig.Notifier.EmailUseTLS
}
