# Bulkhead Execution Subsystem

This package restricts concurrent request executions to prevent thread pools exhaustion.

---

## 1. Features

- **Semaphore limits**: Exposes `NewBulkhead(maxConcurrent)`.
- **Immediate rejection**: Returns `ErrBulkheadFull` when queue overflows.

---

## 2. API Examples

```go
import "prahari/shared/resilience/bulkhead"

b := bulkhead.NewBulkhead(10) // Limit: 10 concurrent requests

err := b.Execute(ctx, func() error {
    return heavySQLOperation()
})
```
