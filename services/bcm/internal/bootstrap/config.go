package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port               int    `env:"PORT" envDefault:"8107"`
	Environment        string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL        string `env:"DATABASE_URL"`
	RedisAddr          string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers       string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IdentityGrpcAddr   string `env:"IDENTITY_GRPC_ADDR" envDefault:"localhost:9090"`
	WorkflowGrpcAddr   string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	EmergencyGrpcAddr  string `env:"EMERGENCY_GRPC_ADDR" envDefault:"localhost:8106"`
	IncidentGrpcAddr   string `env:"INCIDENT_GRPC_ADDR" envDefault:"localhost:8081"`
	AssetGrpcAddr      string `env:"ASSET_GRPC_ADDR" envDefault:"localhost:8088"`
	RiskGrpcAddr       string `env:"RISK_GRPC_ADDR" envDefault:"localhost:8095"`
	ComplianceGrpcAddr string `env:"COMPLIANCE_GRPC_ADDR" envDefault:"localhost:8096"`
	S3Bucket           string `env:"S3_BUCKET" envDefault:"prahari-bcm-evidence"`
	S3Region           string `env:"S3_REGION" envDefault:"ap-south-1"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
