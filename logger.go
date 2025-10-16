package logging

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var core zapcore.Core
var global_init bool = false

func Init() *zap.Logger {
	if core == nil {
		initCore()
	}

	//build global logger
	logger := New()

	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	global_init = true

	return logger
}

func L() *zap.Logger {
	if !global_init {
		Init()
	}
	return zap.L()
}

func Named(name string) *zap.Logger {
	return L().Named(name)
}

func New() *zap.Logger {
	if core == nil {
		initCore()
	}
	return zap.New(core)
}

func initCore() {
	//set env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	LOG_ENV := os.Getenv("LOG_ENV")
	LOG_FORMAT := os.Getenv("LOG_FORMAT")

	//set up core
	stdout := zapcore.AddSync(os.Stdout)

	//default level for prod is "error" dev is "info"
	level := zap.NewAtomicLevelAt(zap.ErrorLevel)
	if LOG_ENV == "dev" {
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	//specified log level overrides LOG_ENV default selection
	switch LOG_LEVEL := os.Getenv("LOG_LEVEL"); LOG_LEVEL {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
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
}
