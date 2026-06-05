package user

import (
	"context"

	"user-query-module/internal/application/port"
	"user-query-module/internal/application/queries/get_user_detail"
	"user-query-module/internal/application/queries/get_user_profile"
	"user-query-module/internal/application/queries/list_user_addresses"
	"user-query-module/internal/application/queries/search_users"
	"user-shared-module/alias"

	"github.com/iamKienb/go-core/app_error"
)

const errMsgUserNotFound = "user was not found"

func (s *queryService) GetUserDetail(ctx context.Context, query get_user_detail.Query) (*get_user_detail.Result, error) {
	user, err := s.findUser(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	return &get_user_detail.Result{User: user}, nil
}

func (s *queryService) GetUserProfile(ctx context.Context, query get_user_profile.Query) (*get_user_profile.Result, error) {
	user, err := s.findUser(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	return &get_user_profile.Result{Profile: user.Profile}, nil
}

func (s *queryService) ListUserAddresses(ctx context.Context, query list_user_addresses.Query) (*list_user_addresses.Result, error) {
	user, err := s.findUser(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	return &list_user_addresses.Result{Addresses: user.Addresses}, nil
}

func (s *queryService) SearchUsers(ctx context.Context, query search_users.Query) (*search_users.Result, error) {
	filters := make([]Filter, 0, 2)
	if query.Status != "" {
		filters = append(filters, Term("status.keyword", query.Status))
	}

	must := make([]Filter, 0, 1)
	if query.Keyword != "" {
		must = append(must, MultiMatch(query.Keyword, []string{"email", "profile.full_name"}))
	}

	result, err := SearchDocuments[port.User](ctx, s.search, SearchSpec{
		Index:   alias.UserAlias,
		Page:    normalizePage(query.Page),
		Filters: filters,
		Must:    must,
		Sorts: []Sort{
			{Field: "updated_at", Direction: SortDesc},
			{Field: "id.keyword", Direction: SortAsc},
		},
	})
	if err != nil {
		return nil, err
	}

	return &search_users.Result{
		Items:         result.Items,
		Total:         result.Total,
		NextPageToken: result.NextPageToken,
	}, nil
}

func (s *queryService) findUser(ctx context.Context, userID string) (*port.User, error) {
	result, err := SearchDocuments[port.User](ctx, s.search, SearchSpec{
		Index: alias.UserAlias,
		Page:  port.Page{Size: 1},
		Filters: []Filter{
			Term("id.keyword", userID),
		},
	})
	if err != nil {
		return nil, err
	}
	if len(result.Items) == 0 {
		return nil, app_error.NotFound(errMsgUserNotFound, nil)
	}

	return &result.Items[0], nil
}

func normalizePage(page port.Page) port.Page {
	if page.Size <= 0 || page.Size > 100 {
		page.Size = 20
	}
	return page
}
