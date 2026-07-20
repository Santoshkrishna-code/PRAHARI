# Distributed Locks Subsystem

This package implements distributed locks using Redis SETNX and safe Lua release scripts, including watchdog lease extensions and Redlock coordinators.

---

## 1. Features

- **Distributed Mutex**: SETNX locks.
- **Safe Lua Unlock**: Checks value ownership, avoiding releasing expired keys.
- **Lease Watchdog**: Automatically extends active lock leases in the background.
- **Redlock Multi-instance**: Coordinates locks across independent nodes.

---

## 2. API Examples

### A. Acquiring distributed locks
```go
import (
    "prahari/shared/redis/locks"
    "time"
)

m := locks.NewMutex(redisClient, "lock:worker:123", "node-uuid", 10*time.Second)

err := m.Lock(ctx)
if err == nil {
    defer m.Unlock(ctx)
    // Execute critical block
}
```

### B. Long running tasks with lease extensions
```go
err := m.Lock(ctx)
if err == nil {
    // Check every 3s and extend to 10s
    cancelWatchdog := locks.StartLeaseExtension(ctx, m, 3*time.Second, 10*time.Second)
    defer func() {
        cancelWatchdog()
        _ = m.Unlock(ctx)
    }()

    // Execute heavy task
}
```
