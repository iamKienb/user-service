package user

import (
	"context"

	"user-query-module/internal/application/queries/get_user_detail"
	"user-query-module/internal/application/queries/get_user_profile"
	"user-query-module/internal/application/queries/list_user_addresses"
	"user-query-module/internal/application/queries/search_users"
	"user-query-module/internal/application/service/models"

	"connectrpc.com/connect"
	api "github.com/iamKienb/api-contract/gen/user"
	"github.com/iamKienb/api-contract/gen/user/userconnect"
)

type queryServer struct {
	getUserDetailExecutor     get_user_detail.Executor
	getUserProfileExecutor    get_user_profile.Executor
	listUserAddressesExecutor list_user_addresses.Executor
	searchUsersExecutor       search_users.Executor
}

func NewQueryServer(
	getUserDetailExecutor get_user_detail.Executor,
	getUserProfileExecutor get_user_profile.Executor,
	listUserAddressesExecutor list_user_addresses.Executor,
	searchUsersExecutor search_users.Executor,
) *queryServer {
	return &queryServer{
		getUserDetailExecutor:     getUserDetailExecutor,
		getUserProfileExecutor:    getUserProfileExecutor,
		listUserAddressesExecutor: listUserAddressesExecutor,
		searchUsersExecutor:       searchUsersExecutor,
	}
}

func (s *queryServer) GetUserDetail(ctx context.Context, req *connect.Request[api.GetUserDetailRequest]) (*connect.Response[api.GetUserDetailResponse], error) {
	result, err := s.getUserDetailExecutor.Execute(ctx, get_user_detail.Query{UserID: req.Msg.GetUserId()})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api.GetUserDetailResponse{User: ToUserView(result.User)}), nil
}

func (s *queryServer) GetUserProfile(ctx context.Context, req *connect.Request[api.GetUserProfileRequest]) (*connect.Response[api.GetUserProfileResponse], error) {
	result, err := s.getUserProfileExecutor.Execute(ctx, get_user_profile.Query{UserID: req.Msg.GetUserId()})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api.GetUserProfileResponse{Profile: ToUserProfileView(result.Profile)}), nil
}

func (s *queryServer) ListUserAddresses(ctx context.Context, req *connect.Request[api.ListUserAddressesRequest]) (*connect.Response[api.ListUserAddressesResponse], error) {
	result, err := s.listUserAddressesExecutor.Execute(ctx, list_user_addresses.Query{UserID: req.Msg.GetUserId()})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api.ListUserAddressesResponse{Addresses: ToUserAddressViews(result.Addresses)}), nil
}

func (s *queryServer) SearchUsers(ctx context.Context, req *connect.Request[api.SearchUsersRequest]) (*connect.Response[api.SearchUsersResponse], error) {
	result, err := s.searchUsersExecutor.Execute(ctx, search_users.Query{
		Keyword: req.Msg.GetKeyword(),
		Status:  req.Msg.GetStatus(),
		Page: models.Page{
			Size:  int(req.Msg.GetPageSize()),
			Token: req.Msg.GetPageToken(),
		},
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api.SearchUsersResponse{
		Users:         ToUserViews(result.Items),
		Total:         result.Total,
		NextPageToken: result.NextPageToken,
	}), nil
}

var _ userconnect.UserQueryServiceHandler = (*queryServer)(nil)
