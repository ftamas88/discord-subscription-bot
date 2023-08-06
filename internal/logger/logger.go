// Package logger Configuration for the Uber/Zap logger package
package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	EnvLogLevel = "GO_LOG_LEVEL"
	EnvLogMode  = "GO_ENV"

	EnvModeDevelopment = "development"
)

func NewLogger() *zap.SugaredLogger {
	var logLevel zapcore.Level

	switch os.Getenv(EnvLogLevel) {
	case zapcore.WarnLevel.String():
		logLevel = zapcore.WarnLevel
	case zapcore.DebugLevel.String():
		logLevel = zapcore.DebugLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	if os.Getenv(EnvLogMode) != EnvModeDevelopment {
		logger := zap.New(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(
					zap.NewProductionEncoderConfig(),
				),
				os.Stdout,
				logLevel,
			),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)

		return logger.Sugar()
	}

	opt := zap.NewDevelopmentEncoderConfig()
	opt.EncodeLevel = zapcore.CapitalColorLevelEncoder
	opt.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02|15:04:05.00")
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(opt),
			os.Stdout,
			logLevel,
		),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.PanicLevel),
	)

	return logger.Sugar()
}
