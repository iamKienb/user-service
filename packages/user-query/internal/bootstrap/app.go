package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"user-query-module/internal/bootstrap/config"
	"user-query-module/internal/bootstrap/module"
	"user-shared-module/indexing"

	configx "github.com/iamKienb/go-core/config"
)

type App struct {
	logger *slog.Logger
	infra  *module.InfraModule
}

func NewApp(logger *slog.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Start(ctx context.Context) error {
	cfg, err := configx.Loader[config.UserQueryConfig]()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	infra, err := module.NewInfraModule(cfg)
	if err != nil {
		return fmt.Errorf("infra: %w", err)
	}
	a.infra = infra

	if err := a.bootstrapIndices(ctx); err != nil {
		return err
	}

	a.logger.Info("starting user query", slog.Any("aliases", []string{indexing.UserAlias, indexing.ShopAlias}))

	<-ctx.Done()
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("shutting down")

	if a.infra != nil && a.infra.ESService != nil {
		a.logger.Info("closing elasticsearch client...")
		if err := a.infra.ESService.Close(ctx); err != nil {
			return fmt.Errorf("close elasticsearch: %w", err)
		}
	}

	a.logger.Info("app stopped cleanly")

	return nil
}

func (a *App) bootstrapIndices(ctx context.Context) error {
	if err := a.infra.ESService.BootstrapIndex(ctx, indexing.UserAlias, indexing.UserMapping); err != nil {
		return fmt.Errorf("bootstrap user index: %w", err)
	}

	if err := a.infra.ESService.BootstrapIndex(ctx, indexing.ShopAlias, indexing.ShopMapping); err != nil {
		return fmt.Errorf("bootstrap shop index: %w", err)
	}

	return nil
}
