package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LOG_ENV string
var LOG_LEVEL string
var LOG_FORMAT string

// TODO: use core not config
var core zapcore.Core

/*
func Init() {
	logger := zap.Must(zap.NewProduction())

	defer logger.Sync()

	logger.Info("Hello from Zap logger!")
}*/

func Init() *zap.Logger {
	LOG_ENV = "prod"
	LOG_LEVEL = "info"
	LOG_FORMAT = "json"

	/*config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout", "logfile"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if LOG_ENV == "dev" {
		config.Development = true
		config.Encoding = "console"
	}

	if LOG_LEVEL == "debug" {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	if LOG_LEVEL == "warn" {
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	}

	if LOG_LEVEL == "error" {
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	if LOG_FORMAT == "console" {
		config.Encoding = "console"
	}

	logger := zap.Must(config.Build())

	//logger := zap.Must(zap.NewProduction())
	/*if LOG_ENV == "dev" {
		logger = zap.Must(zap.NewDevelopment())
	}*/

	//TODO: check if logger instance exists first
	logger := createLogger()

	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	return logger
}

func L() *zap.Logger {
	return zap.L()
}

func Named(name string) *zap.Logger {
	return zap.L().Named(name)
}

func New() {

}

func createLogger() *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	if LOG_LEVEL == "debug" {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	if LOG_LEVEL == "warn" {
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	}

	if LOG_LEVEL == "error" {
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if LOG_ENV == "dev" {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	/*consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	jsonEncoder := zapcore.NewJSONEncoder(productionCfg)*/
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if LOG_FORMAT == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	/*core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(jsonEncoder, stdout, level),
	)*/

	core := zapcore.NewCore(encoder, stdout, level)

	return zap.New(core)
}
