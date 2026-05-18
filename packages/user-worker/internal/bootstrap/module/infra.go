package module

import (
	"context"
	"fmt"
	"user-worker-module/internal/application/port"
	"user-worker-module/internal/bootstrap/config"
	"user-worker-module/internal/infra/elasticsearch"
	kafkax "user-worker-module/internal/infra/kafka"

	esx "github.com/iamKienb/shopify-go-platform/elasticsearch"
)

type InfraModule struct {
	ESService esx.ESXService
	Kafka     kafkax.KafkaXService
	ESRepo    port.ESRepository
}

func NewInfraModule(ctx context.Context, cfg *config.UserWorkerConfig) (*InfraModule, error) {
	esService, err := esx.New(cfg.ES)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch: %w", err)
	}

	kafka, err := kafkax.New(cfg.Kafka)
	if err != nil {
		return nil, fmt.Errorf("kafka: %w", err)
	}

	return &InfraModule{
		ESService: esService,
		Kafka:     kafka,
		ESRepo:    elasticsearch.NewESRepository(esService),
	}, nil
}
