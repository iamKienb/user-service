package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"user-shared-module/events"
	"user-worker-module/internal/bootstrap/config"
	"user-worker-module/internal/bootstrap/module"
	kafkax "user-worker-module/internal/infra/kafka"

	configx "github.com/iamKienb/shopify-go-platform/config"
)

type App struct {
	logger *slog.Logger
	infra  *module.InfraModule
}

func NewApp(logger *slog.Logger) *App {
	return &App{
		logger: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	cfg, err := configx.Loader[config.UserWorkerConfig]()

	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	infra, err := module.NewInfraModule(ctx, cfg)
	if err != nil {
		return err
	}

	a.infra = infra
	cfg.Consumer.Topics = events.Topics
	application := module.NewApplicationModule(infra)

	consumer, err := kafkax.NewConsumer(infra.Kafka, cfg.Consumer, a.logger, application.EventProcessor)
	if err != nil {
		return fmt.Errorf("create consumer: %w", err)
	}
	consumer.Start(ctx)

	a.logger.Info("starting user worker")

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("shutting down")

	if a.infra != nil {
		if a.infra.Kafka != nil {
			a.logger.Info("closing kafka client...")
			_ = a.infra.Kafka.Close()
		}

		if a.infra.ESService != nil {
			a.logger.Info("closing elasticsearch client...")
			_ = a.infra.ESService.Close(ctx)
		}
	}

	a.logger.Info("app stopped cleanly")

	return nil
}
