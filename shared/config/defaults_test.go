package config_test

import (
	"testing"

	"prahari/shared/config"
)

func TestSetDefaults(t *testing.T) {
	var cfg config.Config

	config.SetDefaults(&cfg)

	if cfg.Environment != "development" {
		t.Errorf("expected default Environment 'development', got '%s'", cfg.Environment)
	}

	if cfg.AWS.Region != "us-east-1" {
		t.Errorf("expected default AWS Region 'us-east-1', got '%s'", cfg.AWS.Region)
	}

	if cfg.HTTP.Port != 8080 {
		t.Errorf("expected default HTTP port 8080, got %d", cfg.HTTP.Port)
	}

	if cfg.GRPC.Port != 9090 {
		t.Errorf("expected default gRPC port 9090, got %d", cfg.GRPC.Port)
	}

	if cfg.Database.Port != 5432 {
		t.Errorf("expected default database port 5432, got %d", cfg.Database.Port)
	}
}
