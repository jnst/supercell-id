package logger_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jnst/supercell-id/pkg/logger"
)

func TestNew(t *testing.T) {
	t.Parallel()
	require.NotNil(t, logger.New())
}

func TestWithContext(t *testing.T) {
	type args struct {
		ctx    context.Context
		logger *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				logger: logger.New(),
			},
			want: func() context.Context {
				return context.WithValue(
					context.Background(),
					logger.ExportContextKey,
					logger.New(),
				)
			}(),
		},
		{
			name: "nil context",
			args: args{
				ctx:    nil,
				logger: logger.New(),
			},
			want: context.Background(),
		},
		{
			name: "nil logger",
			args: args{
				ctx:    context.Background(),
				logger: nil,
			},
			want: context.Background(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := logger.WithContext(tt.args.ctx, tt.args.logger)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromContext(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *slog.Logger
	}{
		{
			name: "ok",
			args: args{
				ctx: logger.WithContext(context.Background(), logger.New()),
			},
			want: logger.New(),
		},
		{
			name: "not ok",
			args: args{
				ctx: context.Background(),
			},
			want: logger.New(),
		},
		{
			name: "nil",
			args: args{
				ctx: nil,
			},
			want: logger.New(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := logger.FromContext(tt.args.ctx)
			require.Equal(t, tt.want, got)
		})
	}
}
