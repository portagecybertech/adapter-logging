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

	/*file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})*/

	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, stdout, level),
	)

	return zap.New(core)
}
