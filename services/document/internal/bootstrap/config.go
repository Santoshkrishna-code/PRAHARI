package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port             int    `env:"PORT" envDefault:"8108"`
	Environment      string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL      string `env:"DATABASE_URL"`
	RedisAddr        string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers     string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IdentityGrpcAddr string `env:"IDENTITY_GRPC_ADDR" envDefault:"localhost:9090"`
	WorkflowGrpcAddr string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	MOCGrpcAddr      string `env:"MOC_GRPC_ADDR" envDefault:"localhost:8103"`
	PHAGrpcAddr      string `env:"PHA_GRPC_ADDR" envDefault:"localhost:8104"`
	AuditGrpcAddr    string `env:"AUDIT_GRPC_ADDR" envDefault:"localhost:8097"`
	S3Bucket         string `env:"S3_BUCKET" envDefault:"prahari-document-repository"`
	S3Region         string `env:"S3_REGION" envDefault:"ap-south-1"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
