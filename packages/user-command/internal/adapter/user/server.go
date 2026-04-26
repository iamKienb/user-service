package user

import (
	"context"
	"shopify-user-command-module/internal/application/command/login_user"
	"shopify-user-command-module/internal/application/command/register_user"

	"connectrpc.com/connect"
	userv1 "github.com/iamKienb/shopify-go-api/gen/user"
	"github.com/iamKienb/shopify-go-api/gen/user/userconnect"
)

type userServer struct {
	registerExecutor register_user.Executor
	loginExecutor    login_user.Executor
}

func NewUserServer(registerExecutor register_user.Executor, loginExecutor login_user.Executor) *userServer {
	return &userServer{
		registerExecutor: registerExecutor,
		loginExecutor:    loginExecutor,
	}
}

func (s *userServer) Register(ctx context.Context, req *connect.Request[userv1.RegisterRequest]) (*connect.Response[userv1.RegisterResponse], error) {
	result, err := s.registerExecutor.Execute(ctx, ToRegisterCommand(req.Msg))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToRegisterResponse(result)), nil
}

func (s *userServer) Login(ctx context.Context, req *connect.Request[userv1.LoginRequest]) (*connect.Response[userv1.LoginResponse], error) {
	result, err := s.loginExecutor.Execute(ctx, ToLoginCommand(req.Msg))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToLoginResponse(result)), nil
}

var _ userconnect.UserCommandServiceHandler = (*userServer)(nil)
