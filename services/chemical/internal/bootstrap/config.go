package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port                   int    `env:"PORT" envDefault:"8116"`
	Environment            string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL            string `env:"DATABASE_URL"`
	RedisAddr              string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers           string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	OccHealthGrpcAddr      string `env:"OCC_HEALTH_GRPC_ADDR" envDefault:"localhost:8094"`
	EnvironmentalGrpcAddr  string `env:"ENVIRONMENTAL_GRPC_ADDR" envDefault:"localhost:8095"`
	EmergencyGrpcAddr      string `env:"EMERGENCY_GRPC_ADDR" envDefault:"localhost:8112"`
	PermitGrpcAddr         string `env:"PERMIT_GRPC_ADDR" envDefault:"localhost:8082"`
	RiskGrpcAddr           string `env:"RISK_GRPC_ADDR" envDefault:"localhost:8090"`
	WorkflowGrpcAddr       string `env:"WORKFLOW_GRPC_ADDR" envDefault:"localhost:9091"`
	S3Bucket               string `env:"S3_BUCKET" envDefault:"prahari-chemical-sds"`
	S3Region               string `env:"S3_REGION" envDefault:"ap-south-1"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
