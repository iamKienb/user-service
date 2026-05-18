package kafkax

import (
	"context"
	"net"
	"user-worker-module/internal/bootstrap/config"

	"github.com/segmentio/kafka-go"
)

type KafkaX struct {
	cfg       config.KafkaConfig
	dialer    *kafka.Dialer
	transport *kafka.Transport
}

func New(cfg config.KafkaConfig) (KafkaXService, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	dialer := &kafka.Dialer{
		ClientID:  cfg.ClientID,
		Timeout:   cfg.DialTimeout,
		DualStack: true,
	}

	transport := &kafka.Transport{
		ClientID: cfg.ClientID,
		Dial: (&net.Dialer{
			Timeout: cfg.DialTimeout,
		}).DialContext,
	}

	client := &KafkaX{
		cfg:       cfg,
		dialer:    dialer,
		transport: transport,
	}

	if err := client.Ping(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
