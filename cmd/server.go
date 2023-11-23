package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/jnst/supercell-id/internal/handler"
	"github.com/jnst/supercell-id/internal/interceptor"
	"github.com/jnst/supercell-id/internal/logger"
	"github.com/jnst/supercell-id/pkg/gen/authentication/v1/v1connect"
)

func main() {
	l := logger.New()

	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal(err)
	}

	path, ahandler := v1connect.NewAuthenticationServiceHandler(
		handler.NewAuthenticationServer(),
		connect.WithInterceptors(
			interceptor.NewLoggerInterceptor(l),
			validateInterceptor,
		),
	)

	mux := http.NewServeMux()
	mux.Handle(path, ahandler)

	addr := ":8080"
	l.Info(fmt.Sprintf("listening on %s", addr))

	if err := http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		l.Error("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
