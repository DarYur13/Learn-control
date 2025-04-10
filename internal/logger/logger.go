package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

var globalLogger *zap.SugaredLogger = initLogger()
var level zap.AtomicLevel

func initLogger() *zap.SugaredLogger {

	globalLogger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatalf("globalLogger initialization error: %s", err.Error())
	}

	return globalLogger.Sugar()

}

func New(pathToLogFile string) (*zap.SugaredLogger, error) {
	level = zap.NewAtomicLevel()

	outputPaths := []string{pathToLogFile}
	errorOutputPaths := []string{pathToLogFile}

	if pathToLogFile == "stdout" || pathToLogFile == "" {
		outputPaths = []string{"stdout"}
		errorOutputPaths = []string{"stderr"}
	}

	cfg := zap.Config{
		Encoding:         "console",
		Level:            level,
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      outputPaths,
		ErrorOutputPaths: errorOutputPaths,
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("globalLogger build config error: %s", err.Error())
	}

	return logger.Sugar(), nil
}

func SetLogger(logger *zap.SugaredLogger) {
	globalLogger = logger
}

func SetLogLevel(logLevel string) error {

	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return fmt.Errorf("parse logger level error: %s", err.Error())
	}
	return nil
}

func Info(args ...interface{}) {
	globalLogger.Info(args)
}

func Warn(args ...interface{}) {
	globalLogger.Warn(args)
}

func Debug(args ...interface{}) {
	globalLogger.Debug(args)
}

func DPanic(args ...interface{}) {
	globalLogger.DPanic(args)
}

func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	globalLogger.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	globalLogger.Panic(args)
}

func Fatal(args ...interface{}) {
	globalLogger.Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	globalLogger.Fatalf(template, args)
}

func Sync() {
	globalLogger.Sync()
}
