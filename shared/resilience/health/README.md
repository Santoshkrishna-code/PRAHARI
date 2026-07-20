# Health Aggregator Subsystem

This package aggregates multiple downstream `Checker` implementations to expose a unified connection diagnostics report.

---

## 1. Features

- **Consolidated Health checks**: Maps downstream names to active connection errors.

---

## 2. API Examples

```go
import (
    "prahari/shared/resilience/health"
)

// Aggregate systems checkers
reg := health.NewRegistry(dbChecker, redisChecker, kafkaChecker)

// Check status on health check endpoint trigger
results := reg.CheckHealth(ctx)
for service, err := range results {
    if err != nil {
        // service is offline
    }
}
```
