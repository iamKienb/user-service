package shop

import (
	"context"
	"user-command-module/internal/application/command/add_shop_address"
	"user-command-module/internal/application/command/assign_member"
	"user-command-module/internal/application/command/create_shop"
	"user-command-module/internal/application/command/verify_permission"

	"connectrpc.com/connect"
	"github.com/iamKienb/api-contract/gen/shop"
	"github.com/iamKienb/api-contract/gen/shop/shopconnect"
	authx "github.com/iamKienb/go-core/middleware/auth"
)

type shopServer struct {
	createShopExecutor       create_shop.Executor
	assignMemberExecutor     assign_member.Executor
	addAddressExecutor       add_shop_address.Executor
	verifyPermissionExecutor verify_permission.Executor
}

func NewShopServer(
	createShopExecutor create_shop.Executor,
	assignMemberExecutor assign_member.Executor,
	addAddressExecutor add_shop_address.Executor,
	verifyPermissionExecutor verify_permission.Executor,
) *shopServer {
	return &shopServer{
		createShopExecutor:       createShopExecutor,
		assignMemberExecutor:     assignMemberExecutor,
		addAddressExecutor:       addAddressExecutor,
		verifyPermissionExecutor: verifyPermissionExecutor,
	}
}

func (s *shopServer) CreateShop(ctx context.Context, req *connect.Request[shop.CreateShopRequest]) (*connect.Response[shop.CreateShopResponse], error) {
	currentUser := authx.GetUserInfoFromCtx(ctx)

	cmd, err := ToCreateShopCommand(currentUser.UserID, currentUser.FullName, req.Msg)
	if err != nil {
		return nil, err
	}

	result, err := s.createShopExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToCreateShopResponse(result)), nil
}

func (s *shopServer) AddShopAddress(ctx context.Context, req *connect.Request[shop.AddShopAddressRequest]) (*connect.Response[shop.AddShopAddressResponse], error) {
	currentUser := authx.GetUserInfoFromCtx(ctx)
	cmd, err := ToAddAddressCommand(currentUser.UserID, req.Msg)
	if err != nil {
		return nil, err
	}

	result, err := s.addAddressExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToAddAddressResponse(result)), nil
}

func (s *shopServer) AssignMemberRoles(ctx context.Context, req *connect.Request[shop.AssignMemberRolesRequest]) (*connect.Response[shop.AssignMemberRolesResponse], error) {
	currentUser := authx.GetUserInfoFromCtx(ctx)
	cmd, err := ToAssignMemberCommand(currentUser.UserID, currentUser.FullName, req.Msg)
	if err != nil {
		return nil, err
	}

	result, err := s.assignMemberExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToAssignMemberResponse(result)), nil
}

func (s *shopServer) VerifyPermission(ctx context.Context, req *connect.Request[shop.VerifyPermissionRequest]) (*connect.Response[shop.VerifyPermissionResponse], error) {
	currentUser := authx.GetUserInfoFromCtx(ctx)
	cmd, err := ToVerifyPermissionCommand(currentUser.UserID, req.Msg)
	if err != nil {
		return nil, err
	}

	result, err := s.verifyPermissionExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToVerifyPermissionResponse(result)), nil
}

var _ shopconnect.ShopCommandServiceHandler = (*shopServer)(nil)
