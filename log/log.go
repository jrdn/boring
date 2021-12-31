package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var rootLogger *zap.Logger

type Config struct {
	Level       string
	Development bool
	Encoding    string
	TZ          *time.Location
}

func SetupLogging(cfg Config) error {

	encoderConfig := zap.NewProductionEncoderConfig()

	if cfg.TZ == nil {
		cfg.TZ = time.UTC
	}
	rfc3339EncodingInLocation := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		zapcore.RFC3339NanoTimeEncoder(t.In(cfg.TZ), enc)
	}
	encoderConfig.TimeKey = "ts"
	encoderConfig.EncodeTime = rfc3339EncodingInLocation

	if cfg.Encoding == "console" && cfg.Development {
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	config := zap.Config{
		Development:      cfg.Development,
		Encoding:         cfg.Encoding,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig,
	}

	switch cfg.Level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	l, err := config.Build()
	if err != nil {
		return err
	}

	rootLogger = l
	return nil
}

func GetRootLogger() *zap.Logger {
	return rootLogger
}

func GetLogger(name string) *zap.Logger {
	rootLogger.WithOptions()
	return rootLogger.Named(name)
}

func GetSugaredLogger(name string) *zap.SugaredLogger {
	return rootLogger.Named(name).Sugar()
}
