package logger

import (
	"context"

	"go.uber.org/zap"
)

type loggerContextKey struct{}

var Log *zap.Logger

func New() error {

	l, err := zap.NewProduction()
	if err != nil {
		return err
	}

	Log = l
	return nil
}

func WithContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey{}, logger)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {

	if ctx == nil {
		return Log
	}

	if l, ok := ctx.Value(loggerContextKey{}).(*zap.Logger); ok {
		return l
	}
	return Log
}
