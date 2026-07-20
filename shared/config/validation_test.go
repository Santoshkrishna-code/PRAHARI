package config_test

import (
	"testing"

	"prahari/shared/config"
)

func TestValidation_ValidConfig(t *testing.T) {
	cfg := config.Config{
		Environment: "production",
		ServiceName: "test-service",
		AWS: config.AWSConfig{
			Region:  "us-west-2",
			RoleARN: "arn:aws:iam::123456789012:role/PrahariTestRole",
		},
		Database: config.DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "secure-password",
			Name:     "prahari_db",
			SSLMode:  "require",
		},
		Redis: config.RedisConfig{
			Address: "localhost:6379",
		},
		Kafka: config.KafkaConfig{
			Brokers:  []string{"localhost:9092"},
			ClientID: "test-client",
		},
		HTTP: config.HTTPConfig{
			Port:         8081,
			ReadTimeout:  5,
			WriteTimeout: 10,
			IdleTimeout:  120,
		},
		GRPC: config.GRPCConfig{
			Port:               9091,
			MaxConnectionIdle:  60,
			MaxConnectionAge:   300,
		},
		Telemetry: config.TelemetryConfig{
			PrometheusPort: 9092,
			SampleRate:     0.5,
		},
		Security: config.SecurityConfig{
			AllowedOrigins: []string{"https://prahari.internal"},
			RateLimitRPS:   100.0,
		},
	}

	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected valid config, got error: %v", err)
	}
}

func TestValidation_InvalidPort(t *testing.T) {
	cfg := config.Config{
		Environment: "production",
		ServiceName: "test-service",
		HTTP: config.HTTPConfig{
			Port: 70000, // Invalid port: > 65535
		},
	}

	if err := cfg.Validate(); err == nil {
		t.Error("expected validation failure for out-of-bound HTTP port (>65535)")
	}
}

func TestValidation_InvalidAWSRegion(t *testing.T) {
	cfg := config.Config{
		Environment: "production",
		ServiceName: "test-service",
		AWS: config.AWSConfig{
			Region: "invalid-region-format", // Not matching region regex
		},
	}

	if err := cfg.Validate(); err == nil {
		t.Error("expected validation failure for malformed AWS Region")
	}
}

func TestValidation_InvalidAWSARN(t *testing.T) {
	cfg := config.Config{
		Environment: "production",
		ServiceName: "test-service",
		AWS: config.AWSConfig{
			Region:  "us-east-1",
			RoleARN: "invalid-arn-string",
		},
	}

	if err := cfg.Validate(); err == nil {
		t.Error("expected validation failure for malformed AWS ARN string")
	}
}
