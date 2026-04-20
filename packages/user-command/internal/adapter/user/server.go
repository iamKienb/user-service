package user

import (
	"context"
	"shopify-user-command-module/contract/protogen/user"
	"shopify-user-command-module/contract/protogen/user/userconnect"
	"shopify-user-command-module/internal/application/command/register_user"

	"connectrpc.com/connect"
)

type UserServer struct {
	registerExecutor register_user.Executor
}

func NewUserServer(registerExecutor register_user.Executor) *UserServer {
	return &UserServer{
		registerExecutor: registerExecutor,
	}
}

func (s *UserServer) Register(ctx context.Context, req *connect.Request[user.RegisterRequest]) (*connect.Response[user.RegisterResponse], error) {
	cmd := ToRegisterCommand(req.Msg)

	result, err := s.registerExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToRegisterResponse(result)), nil
}

var _ userconnect.UserCommandServiceHandler = (*UserServer)(nil)
