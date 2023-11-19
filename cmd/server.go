package main

import (
	"log/slog"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/jnst/supercell-id/internal/handler"
	"github.com/jnst/supercell-id/internal/interceptor"
	"github.com/jnst/supercell-id/pkg/gen/authentication/v1/v1connect"
)

func main() {
	mux := http.NewServeMux()
	path, ahandler := v1connect.NewAuthenticationServiceHandler(
		handler.NewAuthenticationServer(),
		connect.WithInterceptors(interceptor.NewLoggerInterceptor()),
	)
	mux.Handle(path, ahandler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("listening on :8080")

	if err := http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{})); err != nil {
		logger.Error("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
