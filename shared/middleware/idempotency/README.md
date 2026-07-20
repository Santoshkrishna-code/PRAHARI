# Idempotency Guards Middleware

This package integrates with Redis to lock incoming request executions, preventing double post submissions and playing back finished responses on retries.

---

## 1. Features

- **Distributed Locks**: Claims request processing tokens via Redis SETNX.
- **Finished Response caching**: Serializes HTTP status, body, and headers.

---

## 2. API Examples

```go
import "prahari/shared/middleware/idempotency"

manager := idempotency.NewManager(redisClient)

// Claim key
ok, _ := manager.AcquireLock(ctx, key, 5*time.Minute)
if !ok {
    // Already processing or completed
    resp, _ := manager.GetResponse(ctx, key)
    if resp != nil {
        // Playback cached response
    }
}
```
