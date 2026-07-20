# Retry Pattern Subsystem

This package implements exponential backoff retries with randomized jitter to manage transient connection failures.

---

## 1. Features

- **Jitter delays**: Decouples retry peaks.
- **Context bindings**: Intercepts request cancellation loops.

---

## 2. API Examples

```go
import "prahari/shared/resilience/retry"

b := retry.Backoff{Min: 100*time.Millisecond, Max: 2*time.Second, Jitter: true}
err := retry.Retry(ctx, b, 3, func() error {
    return runDatabaseQuery()
})
```
