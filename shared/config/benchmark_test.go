package config_test

import (
	"context"
	"os"
	"testing"

	"prahari/shared/config"
)

// BenchmarkLoadConfig benchmarks configuration loading loop parsing rates.
func BenchmarkLoadConfig(b *testing.B) {
	_ = os.Setenv("ENVIRONMENT", "production")
	_ = os.Setenv("SERVICE_NAME", "bench-service")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASSWORD", "pass")
	_ = os.Setenv("DB_NAME", "prahari")
	_ = os.Setenv("REDIS_ADDRESS", "localhost:6379")
	_ = os.Setenv("KAFKA_BROKERS", "localhost:9092")
	_ = os.Setenv("KAFKA_CLIENT_ID", "bench-client")
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = config.Load(context.Background(), nil, nil)
	}
}
