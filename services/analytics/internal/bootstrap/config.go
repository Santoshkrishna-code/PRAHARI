package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port         int    `env:"PORT" envDefault:"8119"`
	Environment  string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL  string `env:"DATABASE_URL"`
	RedisAddr    string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	AdminAddr    string `env:"ADMINISTRATION_GRPC_ADDR" envDefault:"localhost:9001"`
	NotifAddr    string `env:"NOTIFICATION_GRPC_ADDR" envDefault:"localhost:9002"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
