package config

import (
	"context"
	"fmt"
)

// Config represents the unified, consolidated configuration for a PRAHARI microservice.
type Config struct {
	Environment string            `env:"ENVIRONMENT" envDefault:"development" validate:"required,oneof=development test staging production"`
	ServiceName string            `env:"SERVICE_NAME" envDefault:"prahari-service" validate:"required"`
	AWS         AWSConfig         `json:"aws"`
	Database    DatabaseConfig    `json:"database"`
	Redis       RedisConfig       `json:"redis"`
	Kafka       KafkaConfig       `json:"kafka"`
	HTTP        HTTPConfig        `json:"http"`
	GRPC        GRPCConfig        `json:"grpc"`
	Security    SecurityConfig    `json:"security"`
	Telemetry   TelemetryConfig   `json:"telemetry"`
	FeatureFlags FeatureFlagConfig `json:"feature_flags"`
}

// Validate executes structured validations across all configuration fields.
func (c *Config) Validate() error {
	return ValidateStruct(c)
}

// Health reports configuration health for diagnostics logs.
func (c *Config) Health(ctx context.Context) map[string]string {
	return map[string]string{
		"environment":  c.Environment,
		"service_name": c.ServiceName,
		"status":       "READY",
		"http_port":    fmt.Sprintf("%d", c.HTTP.Port),
		"grpc_port":    fmt.Sprintf("%d", c.GRPC.Port),
	}
}
