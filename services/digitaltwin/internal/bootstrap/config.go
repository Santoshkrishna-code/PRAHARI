package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port         int    `env:"PORT" envDefault:"8122"`
	Environment  string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL  string `env:"DATABASE_URL"`
	RedisAddr    string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	GraphAddr    string `env:"GRAPH_STORE_ADDR" envDefault:"localhost:8182"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
