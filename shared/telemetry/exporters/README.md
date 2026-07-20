# Telemetry Exporters Subsystem

This package registers OTLP gRPC trace exporters, Prometheus HTTP scraper endpoints, and development console exporters.

---

## 1. Features

- **OTLP Exporter**: Connects to Jaeger or AWS OTel collectors.
- **Prometheus Reader**: Exposes scraper paths on `/metrics`.
- **Console Exporter**: Outputs to stdout for local testing.

---

## 2. API Examples

```go
import "prahari/shared/telemetry/exporters"

// Expose metrics port
srv := exporters.ServePrometheus(9090)
defer srv.Close()
```
