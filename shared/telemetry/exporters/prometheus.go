package exporters

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	prometheusExporter "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

// NewPrometheusReader instantiates the Prometheus exporter reader.
func NewPrometheusReader() (metric.Reader, error) {
	reader, err := prometheusExporter.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Prometheus metric reader: %w", err)
	}
	return reader, nil
}

// ServePrometheus starts a background HTTP server on the specified port to expose scraping endpoints.
func ServePrometheus(port int) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		// Suppress logs to keep output clean, handles listen-and-serve in the background
		_ = server.ListenAndServe()
	}()

	return server
}
