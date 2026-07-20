package config

// TelemetryConfig configures logging levels, OpenTelemetry trace collectors, and Prometheus metrics ports.
type TelemetryConfig struct {
	Enabled        bool    `env:"TELEMETRY_ENABLED" envDefault:"false"`
	CollectorURL   string  `env:"TELEMETRY_COLLECTOR_URL" validate:"omitempty,url"`
	ServiceName    string  `env:"TELEMETRY_SERVICE_NAME"`
	ServiceVersion string  `env:"TELEMETRY_SERVICE_VERSION" envDefault:"1.0.0"`
	PrometheusPort int     `env:"TELEMETRY_PROMETHEUS_PORT" envDefault:"9090" validate:"required,port"`
	SampleRate     float64 `env:"TELEMETRY_SAMPLE_RATE" envDefault:"1.0" validate:"gte=0,lte=1"`
}
