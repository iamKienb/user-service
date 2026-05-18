package elasticsearch

import (
	"context"
	"user-worker-module/internal/application/port"

	esx "github.com/iamKienb/shopify-go-platform/elasticsearch"
)

type esRepository struct {
	service esx.ESXService
}

func NewESRepository(service esx.ESXService) port.ESRepository {
	return &esRepository{
		service: service,
	}
}

func (r *esRepository) SyncData(ctx context.Context, index string, id string, data any) error {
	return r.service.Sync(ctx, index, id, data)
}
