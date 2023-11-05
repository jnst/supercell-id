package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jnst/supercell-id/internal/handler"
	"github.com/jnst/supercell-id/pkg/gen/authentication/v1/v1connect"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.Handle(v1connect.NewAuthenticationServiceHandler(
		handler.NewAuthenticationServer(),
	))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
