package module

import (
	"fmt"
	"user-query-module/internal/bootstrap/config"

	esx "github.com/iamKienb/go-core/elasticsearch"
)

type InfraModule struct {
	ESService esx.ESXService
}

func NewInfraModule(cfg *config.UserQueryConfig) (*InfraModule, error) {
	esService, err := esx.New(cfg.ES)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch: %w", err)
	}

	return &InfraModule{
		ESService: esService,
	}, nil
}
