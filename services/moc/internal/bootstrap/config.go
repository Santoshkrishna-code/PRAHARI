package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port                int    `env:"PORT" envDefault:"8103"`
	Environment         string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL         string `env:"DATABASE_URL"`
	RedisAddr           string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers        string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IdentityGrpcAddr    string `env:"IDENTITY_GRPC_ADDR" envDefault:"localhost:9090"`
	WorkflowGrpcAddr    string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	RiskGrpcAddr        string `env:"RISK_GRPC_ADDR" envDefault:"localhost:8095"`
	AssetGrpcAddr       string `env:"ASSET_GRPC_ADDR" envDefault:"localhost:8088"`
	MaintenanceGrpcAddr string `env:"MAINTENANCE_GRPC_ADDR" envDefault:"localhost:8087"`
	PermitGrpcAddr      string `env:"PERMIT_GRPC_ADDR" envDefault:"localhost:8089"`
	ComplianceGrpcAddr  string `env:"COMPLIANCE_GRPC_ADDR" envDefault:"localhost:8096"`
	AuditGrpcAddr       string `env:"AUDIT_GRPC_ADDR" envDefault:"localhost:8097"`
	TrainingGrpcAddr    string `env:"TRAINING_GRPC_ADDR" envDefault:"localhost:8098"`
	S3Bucket            string `env:"S3_BUCKET" envDefault:"prahari-moc-evidence"`
	S3Region            string `env:"S3_REGION" envDefault:"ap-south-1"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
