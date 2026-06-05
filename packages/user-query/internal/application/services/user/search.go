package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"user-query-module/internal/application/port"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	esx "github.com/iamKienb/go-core/elasticsearch"
)

type SortDirection string

const (
	SortAsc  SortDirection = "asc"
	SortDesc SortDirection = "desc"
)

type Filter map[string]any

type Sort struct {
	Field     string
	Direction SortDirection
}

type SearchSpec struct {
	Index   string
	Page    port.Page
	Filters []Filter
	Must    []Filter
	Sorts   []Sort
}

type PageResult[T any] struct {
	Items         []T
	Total         int64
	NextPageToken string
}

type SearchService struct {
	esService esx.ESXService
}

func NewSearchService(esService esx.ESXService) *SearchService {
	return &SearchService{esService: esService}
}

func SearchDocuments[T any](ctx context.Context, service *SearchService, spec SearchSpec) (*PageResult[T], error) {
	page := normalizePage(spec.Page)
	raw, err := json.Marshal(buildSearchBody(page, spec.Filters, spec.Must, spec.Sorts))
	if err != nil {
		return nil, fmt.Errorf("marshal search query: %w", err)
	}

	response, err := service.esService.GetClient().Search().
		Index(spec.Index).
		Raw(strings.NewReader(string(raw))).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("search %s: %w", spec.Index, err)
	}

	items, err := decodeHits[T](response.Hits.Hits)
	if err != nil {
		return nil, err
	}

	total := totalHits(response.Hits)
	return &PageResult[T]{
		Items:         items,
		Total:         total,
		NextPageToken: nextPageToken(page, len(items), total),
	}, nil
}

func Term(field string, value any) Filter {
	return Filter{"term": map[string]any{field: value}}
}

func MultiMatch(query string, fields []string) Filter {
	return Filter{"multi_match": map[string]any{"query": query, "fields": fields}}
}

func buildSearchBody(page port.Page, filters []Filter, must []Filter, sorts []Sort) map[string]any {
	boolQuery := map[string]any{}
	if len(filters) > 0 {
		boolQuery["filter"] = filters
	}
	if len(must) > 0 {
		boolQuery["must"] = must
	}
	if len(boolQuery) == 0 {
		boolQuery["must"] = []Filter{{"match_all": map[string]any{}}}
	}

	body := map[string]any{
		"from":             pageOffset(page.Token),
		"size":             page.Size,
		"track_total_hits": true,
		"query": map[string]any{
			"bool": boolQuery,
		},
	}

	return body
}

func decodeHits[T any](hits []types.Hit) ([]T, error) {
	items := make([]T, 0, len(hits))
	for _, hit := range hits {
		var item T
		if err := json.Unmarshal(hit.Source_, &item); err != nil {
			return nil, fmt.Errorf("decode search hit: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
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

func nextPageToken(page port.Page, resultCount int, total int64) string {
	nextOffset := pageOffset(page.Token) + resultCount
	if int64(nextOffset) >= total || resultCount == 0 {
		return ""
	}
	return strconv.Itoa(nextOffset)
}

func totalHits(hits types.HitsMetadata) int64 {
	if hits.Total == nil {
		return 0
	}
	return hits.Total.Value
}
