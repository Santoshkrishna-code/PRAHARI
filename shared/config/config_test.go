package config_test

import (
	"context"
	"testing"

	"prahari/shared/config"
)

func TestConfig_Health(t *testing.T) {
	cfg := &config.Config{
		Environment: "development",
		ServiceName: "test-service",
	}
	cfg.HTTP.Port = 8080
	cfg.GRPC.Port = 9090

	health := cfg.Health(context.Background())

	if health["environment"] != "development" {
		t.Errorf("expected environment 'development', got '%s'", health["environment"])
	}

	if health["service_name"] != "test-service" {
		t.Errorf("expected service 'test-service', got '%s'", health["service_name"])
	}

	if health["status"] != "READY" {
		t.Errorf("expected status 'READY', got '%s'", health["status"])
	}
}
