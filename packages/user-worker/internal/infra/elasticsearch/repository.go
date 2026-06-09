package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"user-worker-module/internal/application/port"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	esx "github.com/iamKienb/go-core/elasticsearch"
)

type esRepository struct {
	service esx.ESXService
	client  *elasticsearch.TypedClient
}

func NewESRepository(service esx.ESXService, client *elasticsearch.TypedClient) port.ESRepository {
	return &esRepository{
		service: service,
		client:  client,
	}
}

func (r *esRepository) SyncData(ctx context.Context, index string, id string, data any) error {
	return r.service.Sync(ctx, index, id, data)
}

func (r *esRepository) SyncNestedData(ctx context.Context, param port.NestedParams) error {
	painlessScript := `
		if (ctx._source[params.field] == null) {
			ctx._source[params.field] = new ArrayList();
		}
		ctx._source[params.field].removeIf(item -> item.id == params.fieldID);
		ctx._source[params.field].add(params.data);
	`

	upsertData := map[string]any{
		param.NestedField: []any{param.Data},
	}

	_, err := r.client.Update(param.Index, param.DocID).
		Script(&types.Script{
			Source: &painlessScript,
			Params: map[string]json.RawMessage{
				"field":   json.RawMessage(fmt.Sprintf(`"%s"`, param.NestedField)),
				"fieldID": json.RawMessage(fmt.Sprintf(`"%s"`, param.NestedFieldID)),
				"data":    r.mustMarshal(param.Data),
			},
		}).
		Upsert(upsertData).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("sync nested data to es failed: %w", err)
	}

	return nil
}

func (r *esRepository) mustMarshal(v any) json.RawMessage {
	bytes, _ := json.Marshal(v)
	return bytes
}
