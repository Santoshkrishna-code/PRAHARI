# PRAHARI Shared Resilience Patterns SDK

This package provides circuit breakers, exponential retries, execution timeouts, bulkhead concurrency limiters, fallback drivers, and health status coordinators.

---

## 1. Directory Structure

- **`circuitbreaker/`**: Sony/gobreaker failure ratio tripping.
- **`retry/`**: Exponential backoffs with jitter.
- **`timeout/`**: Context execution deadlines.
- **`bulkhead/`**: Semaphore concurrency limits.
- **`fallback/`**: Primary command errors redirect.
- **`health/`**: Centralized downstream health aggregator.

---

## 2. Usage Blueprint

```go
import (
	"prahari/shared/resilience/circuitbreaker"
	"prahari/shared/resilience/fallback"
)

cb := circuitbreaker.NewCircuitBreaker("http-breaker", 0.5, 5*time.Second)

res, _ := cb.Execute(func() (interface{}, error) {
    return fallback.Execute(
        func() (interface{}, error) { return callAPIService() },
        func(err error) (interface{}, error) { return "fallback-value", nil },
    )
})
```
