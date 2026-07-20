package telemetry

// Config holds target OTLP endpoints, Prometheus metrics server port, and trace sampling rates.
type Config struct {
	ServiceName    string  `json:"service_name"`
	Environment    string  `json:"environment"`
	Version        string  `json:"version"`
	OTLPEndpoint   string  `json:"otlp_endpoint"`
	PrometheusPort int     `json:"prometheus_port"`
	SampleRatio    float64 `json:"sample_ratio"` // e.g. 0.1 for 10% sampling, 1.0 for 100%
}
