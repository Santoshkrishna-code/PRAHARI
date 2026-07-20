package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

// AppConfig maps variable configs.
type AppConfig struct {
	Port               int    `env:"PORT" envDefault:"8097"`
	Environment        string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL        string `env:"DATABASE_URL"`
	RedisAddr          string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers       string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IdentityGrpcAddr   string `env:"IDENTITY_GRPC_ADDR" envDefault:"localhost:9090"`
	WorkflowGrpcAddr   string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	ContractorGrpcAddr string `env:"CONTRACTOR_GRPC_ADDR" envDefault:"localhost:8089"`
	ComplianceGrpcAddr string `env:"COMPLIANCE_GRPC_ADDR" envDefault:"localhost:8095"`
	AuditGrpcAddr      string `env:"AUDIT_GRPC_ADDR" envDefault:"localhost:8096"`
	IncidentGrpcAddr   string `env:"INCIDENT_GRPC_ADDR" envDefault:"localhost:8084"`
	InspectionGrpcAddr string `env:"INSPECTION_GRPC_ADDR" envDefault:"localhost:8086"`
	RiskGrpcAddr       string `env:"RISK_GRPC_ADDR" envDefault:"localhost:8094"`
	S3Bucket           string `env:"S3_BUCKET" envDefault:"prahari-training-evidence"`
	S3Region           string `env:"S3_REGION" envDefault:"ap-south-1"`
}

// LoadConfig reads settings.
func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return &cfg, nil
}

