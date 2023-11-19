package logger

import (
	"context"
	"log/slog"
	"os"
)

var contextKey = struct{}{}

func New() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func WithContext(ctx context.Context, logger *slog.Logger) context.Context {
	if ctx == nil || logger == nil {
		return context.Background()
	}
	return context.WithValue(ctx, contextKey, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return New()
	}
	l, ok := ctx.Value(contextKey).(*slog.Logger)
	if !ok {
		return New()
	}
	return l
}
