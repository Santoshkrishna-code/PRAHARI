package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port                 int    `env:"PORT" envDefault:"8114"`
	Environment          string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL          string `env:"DATABASE_URL"`
	RedisAddr            string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers         string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IdentityGrpcAddr     string `env:"IDENTITY_GRPC_ADDR" envDefault:"localhost:9090"`
	WorkflowGrpcAddr     string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	IncidentGrpcAddr     string `env:"INCIDENT_GRPC_ADDR" envDefault:"localhost:8081"`
	AuditGrpcAddr        string `env:"AUDIT_GRPC_ADDR" envDefault:"localhost:8093"`
	InspectionGrpcAddr   string `env:"INSPECTION_GRPC_ADDR" envDefault:"localhost:8083"`
	HazardGrpcAddr       string `env:"HAZARD_GRPC_ADDR" envDefault:"localhost:8087"`
	S3Bucket             string `env:"S3_BUCKET" envDefault:"prahari-action-evidence"`
	S3Region             string `env:"S3_REGION" envDefault:"ap-south-1"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
