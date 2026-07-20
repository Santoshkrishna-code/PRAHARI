package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port                 int    `env:"PORT" envDefault:"8117"`
	Environment          string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL          string `env:"DATABASE_URL"`
	RedisAddr            string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers         string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IamGrpcAddr          string `env:"IAM_GRPC_ADDR" envDefault:"localhost:9090"`
	NotificationGrpcAddr string `env:"NOTIFICATION_GRPC_ADDR" envDefault:"localhost:9091"`
	WorkflowGrpcAddr     string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9092"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
