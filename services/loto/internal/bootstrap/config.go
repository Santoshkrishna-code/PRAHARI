package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port                 int    `env:"PORT" envDefault:"8113"`
	Environment          string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL          string `env:"DATABASE_URL"`
	RedisAddr            string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers         string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	IdentityGrpcAddr     string `env:"IDENTITY_GRPC_ADDR" envDefault:"localhost:9090"`
	WorkflowGrpcAddr     string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	PermitGrpcAddr       string `env:"PERMIT_GRPC_ADDR" envDefault:"localhost:8082"`
	MaintenanceGrpcAddr  string `env:"MAINTENANCE_GRPC_ADDR" envDefault:"localhost:8085"`
	AssetGrpcAddr        string `env:"ASSET_GRPC_ADDR" envDefault:"localhost:8084"`
	ShiftGrpcAddr        string `env:"SHIFT_GRPC_ADDR" envDefault:"localhost:8109"`
	ContractorGrpcAddr   string `env:"CONTRACTOR_GRPC_ADDR" envDefault:"localhost:8086"`
	BarrierGrpcAddr      string `env:"BARRIER_GRPC_ADDR" envDefault:"localhost:8105"`
	DocumentGrpcAddr     string `env:"DOCUMENT_GRPC_ADDR" envDefault:"localhost:8108"`
	S3Bucket             string `env:"S3_BUCKET" envDefault:"prahari-loto-evidence"`
	S3Region             string `env:"S3_REGION" envDefault:"ap-south-1"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
