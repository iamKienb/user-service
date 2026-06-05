package service

import (
	"context"
	"strconv"

	"user-query-module/internal/application/queries/get_user_detail"
	"user-query-module/internal/application/queries/get_user_profile"
	"user-query-module/internal/application/queries/list_user_addresses"
	"user-query-module/internal/application/queries/search_users"
	"user-query-module/internal/application/service/models"
	"user-shared-module/alias"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/iamKienb/go-core/app_error"
	esx "github.com/iamKienb/go-core/elasticsearch"
)

type QueryService interface {
	GetUserDetail(ctx context.Context, query get_user_detail.Query) (*get_user_detail.Result, error)
	GetUserProfile(ctx context.Context, query get_user_profile.Query) (*get_user_profile.Result, error)
	ListUserAddresses(ctx context.Context, query list_user_addresses.Query) (*list_user_addresses.Result, error)
	SearchUsers(ctx context.Context, query search_users.Query) (*search_users.Result, error)
}

type queryService struct {
	esClient *elasticsearch.TypedClient
	index    string
}

const (
	defaultPageSize   = 20
	maxPageSize       = 100
	sortAsc           = "asc"
	sortDesc          = "desc"
	errMsgUserMissing = "user was not found"
)

func NewQueryService(esService esx.ESXService) QueryService {
	return &queryService{
		esClient: esService.GetClient(),
		index:    alias.UserAlias,
	}
}

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
	page := normalizePage(query.Page)
	builder := NewQueryBuilder().
		WithPagination(pageOffset(page.Token), page.Size).
		MustMultiMatch(query.Keyword, []string{"email", "profile.full_name"}).
		WithSort("updated_at", sortDesc).
		WithSort("id.keyword", sortAsc)

	if query.Status != "" {
		builder.FilterTerm("status.keyword", query.Status)
	}

	result, err := SearchDocuments[models.User](ctx, s.esClient, s.index, builder.Build())
	if err != nil {
		return nil, err
	}

	items := usersFromHits(result.Hits)
	return &search_users.Result{
		Items:         items,
		Total:         result.Total,
		NextPageToken: nextPageToken(page, len(items), result.Total),
	}, nil
}

func (s *queryService) findUser(ctx context.Context, userID string) (*models.User, error) {
	searchQuery := NewQueryBuilder().
		WithPagination(0, 1).
		FilterTerm("id.keyword", userID).
		Build()

	result, err := SearchDocuments[models.User](ctx, s.esClient, s.index, searchQuery)
	if err != nil {
		return nil, err
	}
	if len(result.Hits) == 0 {
		return nil, app_error.NotFound(errMsgUserMissing, nil)
	}
	return &result.Hits[0].Source, nil
}

func usersFromHits(hits []SearchHit[models.User]) []models.User {
	items := make([]models.User, 0, len(hits))
	for _, hit := range hits {
		items = append(items, hit.Source)
	}
	return items
}

func normalizePage(page models.Page) models.Page {
	if page.Size <= 0 || page.Size > maxPageSize {
		page.Size = defaultPageSize
	}
	return page
}

func pageOffset(token string) int {
	if token == "" {
		return 0
	}
	offset, err := strconv.Atoi(token)
	if err != nil || offset < 0 {
		return 0
	}
	return offset
}

func nextPageToken(page models.Page, resultCount int, total int64) string {
	nextOffset := pageOffset(page.Token) + resultCount
	if int64(nextOffset) >= total || resultCount == 0 {
		return ""
	}
	return strconv.Itoa(nextOffset)
}
