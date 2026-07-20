package bootstrap

import (
	"context"
	"fmt"

	prahariConfig "prahari/shared/config"
	"prahari/edge/gateway/internal/domain/route"
)

// AppConfig maps ports and target upstreams mappings.
type AppConfig struct {
	Port             int           `env:"PORT" envDefault:"8080"`
	Environment      string        `env:"ENVIRONMENT" envDefault:"development"`
	RedisAddr        string        `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	IdentityGrpcAddr string        `env:"IDENTITY_GRPC_ADDR"`
	Routes           []route.Route `json:"routes"`
}

// LoadConfig reads settings from YAML overrides and environment keys.
func LoadConfig(ctx context.Context) (*AppConfig, error) {
	loader, err := prahariConfig.NewLoader(prahariConfig.Options{})
	if err != nil {
		return nil, fmt.Errorf("failed to create config loader: %w", err)
	}

	var cfg AppConfig
	if err := loader.Load(ctx, &cfg); err != nil {
		return nil, fmt.Errorf("failed to load overrides: %w", err)
	}

	return &cfg, nil
}
