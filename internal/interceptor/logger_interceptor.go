package interceptor

import (
	"context"

	"connectrpc.com/connect"

	"github.com/jnst/supercell-id/pkg/logger"
)

func NewLoggerInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			l := logger.New()
			l.Info("request: %v", req)
			ctx = logger.WithContext(ctx, l)
			return next(ctx, req)
		}
	}
}
