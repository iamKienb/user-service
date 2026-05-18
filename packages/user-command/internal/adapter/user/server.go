package user

import (
	"context"
	"user-command-module/internal/application/command/add_user_address"
	"user-command-module/internal/application/command/login_user"
	"user-command-module/internal/application/command/register_user"

	"connectrpc.com/connect"
	"github.com/iamKienb/api-contract/gen/user"
	"github.com/iamKienb/api-contract/gen/user/userconnect"
)

type userServer struct {
	registerExecutor register_user.Executor
	loginExecutor    login_user.Executor
	addAddress       add_user_address.Executor
}

func NewUserServer(
	registerExecutor register_user.Executor,
	loginExecutor login_user.Executor,
	addAddressExecutor add_user_address.Executor,
) *userServer {
	return &userServer{
		registerExecutor: registerExecutor,
		loginExecutor:    loginExecutor,
		addAddress:       addAddressExecutor,
	}
}

func (s *userServer) RegisterUser(ctx context.Context, req *connect.Request[user.RegisterUserRequest]) (*connect.Response[user.RegisterUserResponse], error) {
	result, err := s.registerExecutor.Execute(ctx, ToRegisterCommand(req.Msg))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToRegisterResponse(result)), nil
}

func (s *userServer) LoginUser(ctx context.Context, req *connect.Request[user.LoginUserRequest]) (*connect.Response[user.LoginUserResponse], error) {
	result, err := s.loginExecutor.Execute(ctx, ToLoginCommand(req.Msg))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToLoginResponse(result)), nil
}

func (s *userServer) AddUserAddress(ctx context.Context, req *connect.Request[user.AddUserAddressRequest]) (*connect.Response[user.AddUserAddressResponse], error) {
	cmd, err := ToAddAddressCommand(req.Msg)
	if err != nil {
		return nil, err
	}

	result, err := s.addAddress.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToAddAddressResponse(result)), nil
}

var _ userconnect.UserCommandServiceHandler = (*userServer)(nil)
