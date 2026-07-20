# Circuit Breaker Subsystem

This package wraps standard circuit breaker patterns around request execution routes using `sony/gobreaker`.

---

## 1. Features

- **Tripping Ratio**: Automatically trips after failure thresholds are crossed.
- **Cooldown transition**: Allows half-open testing before closing the gate.

---

## 2. API Examples

```go
import (
    "prahari/shared/resilience/circuitbreaker"
    "time"
)

cb := circuitbreaker.NewCircuitBreaker("db-breaker", 0.5, 10*time.Second)

res, err := cb.Execute(func() (interface{}, error) {
    return queryDatabase()
})
```
