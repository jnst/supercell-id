package handler

import (
	"context"

	"connectrpc.com/connect"

	"github.com/jnst/supercell-id/internal/logger"
	v1 "github.com/jnst/supercell-id/pkg/gen/authentication/v1"
)

type AuthenticationServer struct {
}

func NewAuthenticationServer() *AuthenticationServer {
	return &AuthenticationServer{}
}

func (s *AuthenticationServer) Register(
	ctx context.Context,
	req *connect.Request[v1.RegisterRequest],
) (*connect.Response[v1.RegisterResponse], error) {

	// TODO: implement here
	logger.FromContext(ctx).Info("register")

	return connect.NewResponse[v1.RegisterResponse](&v1.RegisterResponse{
		Status: v1.RegisterResponse_STATUS_SUCCESS,
	}), nil
}

func (s *AuthenticationServer) Signin(
	ctx context.Context,
	req *connect.Request[v1.SigninRequest],
) (*connect.Response[v1.SigninResponse], error) {

	// TODO: implement here

	return connect.NewResponse[v1.SigninResponse](&v1.SigninResponse{
		Status: v1.SigninResponse_STATUS_SUCCESS,
	}), nil
}
