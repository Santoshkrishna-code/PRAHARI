package bootstrap

import (
	"context"
	"fmt"

	prahariConfig "prahari/shared/config"
)

// AppConfig maps database URL keys and Redis addresses.
type AppConfig struct {
	Port             int    `env:"PORT" envDefault:"8083"`
	Environment      string `env:"ENVIRONMENT" envDefault:"development"`
	RedisAddr        string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	DatabaseURL      string `env:"DATABASE_URL"`
	IdentityGrpcAddr string `env:"IDENTITY_GRPC_ADDR"`
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
