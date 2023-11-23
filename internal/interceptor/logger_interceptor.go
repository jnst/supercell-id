package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"

	"github.com/jnst/supercell-id/internal/logger"
)

func NewLoggerInterceptor(l *slog.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			l.Info("incoming request", "req", req)
			ctx = logger.WithContext(ctx, l)
			return next(ctx, req)
		}
	}
}
