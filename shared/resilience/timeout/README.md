# Timeout Execution Subsystem

This package wraps context timeouts around arbitrary execution loops to enforce request deadlines.

---

## 1. Features

- **Deadline enforcement**: Exposes `Execute(ctx, duration, fn)`.

---

## 2. API Examples

```go
import "prahari/shared/resilience/timeout"

err := timeout.Execute(ctx, 500*time.Millisecond, func(ctx context.Context) error {
    return queryBillingAPI(ctx)
})
```
