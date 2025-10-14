package logging

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var core zapcore.Core

func Init() *zap.Logger {
	//set env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	LOG_ENV := os.Getenv("LOG_ENV")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")
	LOG_FORMAT := os.Getenv("LOG_FORMAT")

	//set up core
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

	if LOG_ENV == "dev" {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if LOG_FORMAT == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core = zapcore.NewCore(encoder, stdout, level)

	//build global logger
	logger, err := New()
	if err != nil {
		panic(err)
	}

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

func New() (*zap.Logger, error) {
	if core == nil {
		return zap.New(core), errors.New("Core not initialized")
	}
	return zap.New(core), nil
}
