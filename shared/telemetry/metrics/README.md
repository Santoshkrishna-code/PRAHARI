# OpenTelemetry Metrics Subsystem

This package wraps standard meter instruments to simplify count and latency recordings.

---

## 1. Features

- **Counter**: Int64 accumulator.
- **Histogram**: Float64 latency distribution recorder.

---

## 2. API Examples

```go
import "prahari/shared/telemetry/metrics"

m := metrics.NewMetrics("auth-service")
counter, _ := m.NewCounter("api_calls_count", "Total api queries")
counter.Add(ctx, 1)
```
