package user

import (
	"context"
	"user-command-module/internal/application/commands/add_user_address"
	"user-command-module/internal/application/commands/login_user"
	"user-command-module/internal/application/commands/register_user"
	get_user_address_by_id "user-command-module/internal/application/queries/get_address_by_id"

	"connectrpc.com/connect"
	"github.com/iamKienb/api-contract/gen/user"
	"github.com/iamKienb/api-contract/gen/user/userconnect"
	authx "github.com/iamKienb/go-core/middleware/auth"
)

type userServer struct {
	registerExecutor   register_user.Executor
	loginExecutor      login_user.Executor
	addAddressExecutor add_user_address.Executor
	getAddressExecutor get_user_address_by_id.Executor
}

func NewUserServer(
	registerExecutor register_user.Executor,
	loginExecutor login_user.Executor,
	addAddressExecutor add_user_address.Executor,
	getAddressExecutor get_user_address_by_id.Executor,
) *userServer {
	return &userServer{
		registerExecutor:   registerExecutor,
		loginExecutor:      loginExecutor,
		addAddressExecutor: addAddressExecutor,
		getAddressExecutor: getAddressExecutor,
	}
}

func (s *userServer) RegisterUser(ctx context.Context, req *connect.Request[user.RegisterUserRequest]) (*connect.Response[user.RegisterUserResponse], error) {
	result, err := s.registerExecutor.Execute(ctx, toRegisterCommand(req.Msg))

	if err != nil {
		return nil, toMapError(err)
	}

	return connect.NewResponse(toRegisterResponse(result)), nil
}

func (s *userServer) LoginUser(ctx context.Context, req *connect.Request[user.LoginUserRequest]) (*connect.Response[user.LoginUserResponse], error) {
	result, err := s.loginExecutor.Execute(ctx, toLoginCommand(req.Msg))
	if err != nil {
		return nil, toMapError(err)
	}

	return connect.NewResponse(toLoginResponse(result)), nil
}

func (s *userServer) AddUserAddress(ctx context.Context, req *connect.Request[user.AddUserAddressRequest]) (*connect.Response[user.AddUserAddressResponse], error) {
	currentUser := authx.GetUserInfoFromCtx(ctx)
	cmd, err := toAddAddressCommand(currentUser.UserID, req.Msg)
	if err != nil {
		return nil, err
	}

	result, err := s.addAddressExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, toMapError(err)
	}

	return connect.NewResponse(toAddAddressResponse(result)), nil
}

func (s *userServer) GetUserAddressByID(ctx context.Context, req *connect.Request[user.GetUserAddressByIDRequest]) (*connect.Response[user.GetUserAddressByIDResponse], error) {
	qry, err := toGetUserAddressQuery(req.Msg)
	if err != nil {
		return nil, toMapError(err)
	}

	result, err := s.getAddressExecutor.Execute(ctx, qry)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(toGetUserAddressResponse(result)), nil
}

var _ userconnect.UserCommandServiceHandler = (*userServer)(nil)
