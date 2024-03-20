package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger    *zap.Logger
	loggerKey lkey
)

type lkey struct{}

func ContextLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

func FromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return l
	}
	return zap.L()
}

func New(level string) *zap.Logger {
	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel(level)),
		Development:       true,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{"stdout", "log.txt"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	logger = zap.Must(logConfig.Build())

	return logger
}

func logLevel(level string) zapcore.Level {
	switch level {
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "info":
		return zapcore.InfoLevel
	default:
		return zapcore.DebugLevel
	}
}
