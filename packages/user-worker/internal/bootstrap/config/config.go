package config

import (
	"time"

	configx "github.com/iamKienb/shopify-go-platform/config"
	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct {
	Brokers      []string      `env:"_KAFKA_BROKERS"`
	ClientID     string        `env:"_KAFKA_CLIENT_ID"`
	DialTimeout  time.Duration `env:"_KAFKA_DIAL_TIMEOUT" envDefault:"5s"`
	ReadTimeout  time.Duration `env:"_KAFKA_READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout time.Duration `env:"_KAFKA_WRITE_TIMEOUT" envDefault:"10s"`
}

type ProducerConfig struct {
	Topic          string             `env:"USER_WORKER_SERVICE_PRODUCER_TOPIC"`
	Balancer       string             `env:"_PRODUCER_BALANCER"`
	Compression    string             `env:"_PRODUCER_COMPRESSION"`
	BatchTimeout   time.Duration      `env:"_PRODUCER_BATCH_TIMEOUT"`
	BatchBytes     int64              `env:"_PRODUCER_BATCH_BYTES"`
	RequiredAcks   kafka.RequiredAcks `env:"_PRODUCER_REQUIRED_ACKS"`
	AllowAutoTopic bool               `env:"_PRODUCER_ALLOW_AUTO_TOPIC"`
	WriteTimeout   time.Duration      `env:"_PRODUCER_WRITE_TIMEOUT"`
	ReadTimeout    time.Duration      `env:"_PRODUCER_READ_TIMEOUT"`
	MaxAttempts    int                `env:"_PRODUCER_MAX_ATTEMPTS"`
}

type ConsumerConfig struct {
	GroupID      string `env:"_CONSUMER_GROUP_ID"`
	Topics       []string
	DLQTopic     string        `env:"_CONSUMER_DLQ_TOPIC"`
	MinBytes     int           `env:"_CONSUMER_MIN_BYTES"`
	MaxBytes     int           `env:"_CONSUMER_MAX_BYTES"`
	MaxWait      time.Duration `env:"_CONSUMER_MAX_WAIT"`
	MaxAttempts  int           `env:"_CONSUMER_MAX_ATTEMPTS"`
	RetryBackoff time.Duration `env:"_CONSUMER_RETRY_BACKOFF"`
}

type UserWorkerConfig struct {
	ES       configx.ElasticSearchConfig `envPrefix:"USER_WORKER_SERVICE"`
	Kafka    KafkaConfig                 `envPrefix:"USER_WORKER_SERVICE"`
	Consumer ConsumerConfig              `envPrefix:"USER_WORKER_SERVICE"`
}
