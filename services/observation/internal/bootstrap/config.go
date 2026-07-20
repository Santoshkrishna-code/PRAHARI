package bootstrap

import (
	"context"
	"fmt"

	prahariConfig "prahari/shared/config"
)

// AppConfig maps variable configs.
type AppConfig struct {
	Port             int    `env:"PORT" envDefault:"8093"`
	Environment      string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL      string `env:"DATABASE_URL"`
	RedisAddr        string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers     string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IdentityGrpcAddr string `env:"IDENTITY_GRPC_ADDR" envDefault:"localhost:9090"`
	WorkflowGrpcAddr string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	HazardGrpcAddr   string `env:"HAZARD_GRPC_ADDR" envDefault:"localhost:8091"`
	NearMissGrpcAddr string `env:"NEARMISS_GRPC_ADDR" envDefault:"localhost:8092"`
	IncidentGrpcAddr string `env:"INCIDENT_GRPC_ADDR" envDefault:"localhost:8084"`
	MaintenanceGrpcAddr string `env:"MAINTENANCE_GRPC_ADDR" envDefault:"localhost:8088"`
	ContractorGrpcAddr string `env:"CONTRACTOR_GRPC_ADDR" envDefault:"localhost:8089"`
	InspectionGrpcAddr string `env:"INSPECTION_GRPC_ADDR" envDefault:"localhost:8086"`
	AssetGrpcAddr    string `env:"ASSET_GRPC_ADDR" envDefault:"localhost:8087"`
	S3Bucket         string `env:"S3_BUCKET" envDefault:"prahari-observation-evidence"`
	S3Region         string `env:"S3_REGION" envDefault:"ap-south-1"`
}

// LoadConfig reads settings.
func LoadConfig(ctx context.Context) (*AppConfig, error) {
	loader, err := prahariConfig.NewLoader(prahariConfig.Options{})
	if err != nil {
		return nil, fmt.Errorf("failed to create config loader: %w", err)
	}

	var cfg AppConfig
	if err := loader.Load(ctx, &cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return &cfg, nil
}
