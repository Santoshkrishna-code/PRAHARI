package config_test

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"prahari/shared/config"
)

// MockSSMClient mock wrapper for SSM ParametersStore API.
type MockSSMClient struct {
	Value string
	Err   error
}

func (m *MockSSMClient) GetParameter(ctx context.Context, params *ssm.GetParameterInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return &ssm.GetParameterOutput{
		Parameter: &types.Parameter{
			Value: &m.Value,
		},
	}, nil
}

// MockSecretsManagerClient mock wrapper for SecretsManager API.
type MockSecretsManagerClient struct {
	Value string
	Err   error
}

func (m *MockSecretsManagerClient) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return &secretsmanager.GetSecretValueOutput{
		SecretString: &m.Value,
	}, nil
}

func TestLoad_Precedence(t *testing.T) {
	// Setup standard environment overrides
	_ = os.Setenv("ENVIRONMENT", "production")
	_ = os.Setenv("SERVICE_NAME", "override-service")
	_ = os.Setenv("DB_HOST", "prod-database-host")
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASSWORD", "secret")
	_ = os.Setenv("DB_NAME", "prahari_prod")
	_ = os.Setenv("REDIS_ADDRESS", "prod-redis:6379")
	_ = os.Setenv("KAFKA_BROKERS", "prod-kafka:9092")
	_ = os.Setenv("KAFKA_CLIENT_ID", "prahari-test")
	_ = os.Setenv("SECURITY_ALLOWED_ORIGINS", "https://prahari.com")
	
	defer func() {
		_ = os.Unsetenv("ENVIRONMENT")
		_ = os.Unsetenv("SERVICE_NAME")
		_ = os.Unsetenv("DB_HOST")
		_ = os.Unsetenv("DB_USER")
		_ = os.Unsetenv("DB_PASSWORD")
		_ = os.Unsetenv("DB_NAME")
		_ = os.Unsetenv("REDIS_ADDRESS")
		_ = os.Unsetenv("KAFKA_BROKERS")
		_ = os.Unsetenv("KAFKA_CLIENT_ID")
		_ = os.Unsetenv("SECURITY_ALLOWED_ORIGINS")
	}()

	cfg, err := config.Load(context.Background(), nil, nil)
	if err != nil {
		t.Fatalf("expected successful config load, got error: %v", err)
	}

	if cfg.Environment != "production" {
		t.Errorf("expected Environment override 'production', got '%s'", cfg.Environment)
	}

	if cfg.ServiceName != "override-service" {
		t.Errorf("expected ServiceName override 'override-service', got '%s'", cfg.ServiceName)
	}
}

func TestLoad_WithMockSSMAndSecrets(t *testing.T) {
	_ = os.Setenv("ENVIRONMENT", "production")
	_ = os.Setenv("SERVICE_NAME", "override-service")
	_ = os.Setenv("DB_HOST", "defaults-host") // Initial env host setting
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASSWORD", "env-password")
	_ = os.Setenv("DB_NAME", "prahari_prod")
	_ = os.Setenv("REDIS_ADDRESS", "prod-redis:6379")
	_ = os.Setenv("KAFKA_BROKERS", "prod-kafka:9092")
	_ = os.Setenv("KAFKA_CLIENT_ID", "prahari-test")
	_ = os.Setenv("SECURITY_ALLOWED_ORIGINS", "https://prahari.com")

	// Set parameters search keys triggering AWS API lookups
	_ = os.Setenv("SSM_DB_HOST_PARAM", "/prahari/prod/db/host")
	_ = os.Setenv("SECRETS_DB_PASSWORD_ID", "/prahari/prod/db/password")

	defer func() {
		_ = os.Unsetenv("ENVIRONMENT")
		_ = os.Unsetenv("SERVICE_NAME")
		_ = os.Unsetenv("DB_HOST")
		_ = os.Unsetenv("DB_USER")
		_ = os.Unsetenv("DB_PASSWORD")
		_ = os.Unsetenv("DB_NAME")
		_ = os.Unsetenv("REDIS_ADDRESS")
		_ = os.Unsetenv("KAFKA_BROKERS")
		_ = os.Unsetenv("KAFKA_CLIENT_ID")
		_ = os.Unsetenv("SECURITY_ALLOWED_ORIGINS")
		_ = os.Unsetenv("SSM_DB_HOST_PARAM")
		_ = os.Unsetenv("SECRETS_DB_PASSWORD_ID")
	}()

	mockSSM := &MockSSMClient{Value: "ssm-overridden-db-host"}
	mockSecrets := &MockSecretsManagerClient{Value: "secrets-manager-overridden-password"}

	cfg, err := config.Load(context.Background(), mockSSM, mockSecrets)
	if err != nil {
		t.Fatalf("expected successful config load, got error: %v", err)
	}

	// Verify that SSM parameter took precedence over initial environment host value
	if cfg.Database.Host != "ssm-overridden-db-host" {
		t.Errorf("expected Host override from SSM 'ssm-overridden-db-host', got '%s'", cfg.Database.Host)
	}

	// Verify that Secrets Manager took precedence over env password value
	if cfg.Database.Password != "secrets-manager-overridden-password" {
		t.Errorf("expected Password override from SecretsManager 'secrets-manager-overridden-password', got '%s'", cfg.Database.Password)
	}
}
