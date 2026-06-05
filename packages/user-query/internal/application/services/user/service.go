package user

import (
	"context"

	"user-query-module/internal/application/queries/get_user_detail"
	"user-query-module/internal/application/queries/get_user_profile"
	"user-query-module/internal/application/queries/list_user_addresses"
	"user-query-module/internal/application/queries/search_users"

	esx "github.com/iamKienb/go-core/elasticsearch"
)

type QueryService interface {
	GetUserDetail(ctx context.Context, query get_user_detail.Query) (*get_user_detail.Result, error)
	GetUserProfile(ctx context.Context, query get_user_profile.Query) (*get_user_profile.Result, error)
	ListUserAddresses(ctx context.Context, query list_user_addresses.Query) (*list_user_addresses.Result, error)
	SearchUsers(ctx context.Context, query search_users.Query) (*search_users.Result, error)
}

type queryService struct {
	search *SearchService
}

func NewQueryService(esService esx.ESXService) QueryService {
	return &queryService{search: NewSearchService(esService)}
}
