# HTTP Telemetry Metrics Middleware

This package records processing latency sums and total throughput stats.

---

## 1. Features

- **Total Requests Counter**: atomic counts increments.
- **Latency tracker**: aggregates latency times.

---

## 2. API Examples

```go
import "prahari/shared/middleware/metrics"

tracker := metrics.NewTracker()
router.Use(tracker.Middleware)
```
