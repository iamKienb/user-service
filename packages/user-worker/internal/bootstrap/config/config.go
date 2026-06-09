package config

import (
	configx "github.com/iamKienb/go-core/config"
)

type UserWorkerConfig struct {
	ES       configx.ElasticSearchConfig `envPrefix:"USER_WORKER_SERVICE"`
	Redis    configx.RedisConfig         `envPrefix:"USER_WORKER_SERVICE"`
	Kafka    configx.KafkaConfig         `envPrefix:"USER_WORKER_SERVICE"`
	Consumer configx.ConsumerConfig      `envPrefix:"USER_WORKER_SERVICE"`
}
