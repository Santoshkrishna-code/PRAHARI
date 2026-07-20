package config

type ObservabilityConfig struct {
	EnableTracing bool   `env:"OBSERVABILITY_ENABLE_TRACING" envDefault:"false"`
	CollectorURL  string `env:"OBSERVABILITY_COLLECTOR_URL"` // OTLP collector endpoint
}
