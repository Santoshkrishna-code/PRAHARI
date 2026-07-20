# Exponential Backoff Retry Subsystem

This package implements operational execution loops with backoffs and jitter calculations.

---

## 1. Features

- **Exponential Duration Builder**: Configures scaling multipliers and thresholds.
- **Random Jitter**: Minimizes thundering herd overlaps.
- **Context support**: Intercepts cancels signals between delay sleeps.

---

## 2. API Examples

```go
import "prahari/shared/kafka/retry"

b := retry.Backoff{Min: 100*time.Millisecond, Max: 2*time.Second, Jitter: true}
err := retry.Retry(ctx, b, 5, func() error {
    return publishMessage()
})
```
