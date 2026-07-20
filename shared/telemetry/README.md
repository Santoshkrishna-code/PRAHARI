# PRAHARI Telemetry & Observability Platform SDK

This package provides OpenTelemetry TracerProviders, MeterProviders, Prometheus scraper registries, W3C propagators, sampling selectors, and trace segment lifecycles.

---

## 1. Directory Structure

- **`tracer/`**: Start/end spans.
- **`metrics/`**: Counters and histograms.
- **`exporters/`**: OTLP, Prometheus, and Console print exporters.

---

## 2. Usage Blueprint

Exposes connection checks (`Ping`) and manages providers lifecycle.
```go
import (
    "prahari/shared/telemetry"
    "prahari/shared/telemetry/exporters"
)

// Setup resources
res, _ := telemetry.NewResource(ctx, config)

// Setup exporters
exp := exporters.NewStdoutExporter()

// Setup Provider
tp, _ := telemetry.InitTracerProvider(res, telemetry.GetSampler(1.0), exp)
```
