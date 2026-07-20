# Distributed Rate Limiting Subsystem

This package implements distributed rate limiters supporting Fixed Window, Sliding Window, and Token Bucket algorithms.

---

## 1. Features

- **Fixed Window**: Tracks request epoch boundaries.
- **Sliding Window Log**: Prevents micro-burst limit bypasses using sorted sets.
- **Token Bucket**: Implements smooth request throttle parameters via Lua.

---

## 2. API Examples

```go
import "prahari/shared/redis/ratelimit"

limiter := ratelimit.NewTokenBucketLimiter(client, 100, 10.0) // Capacity 100, fills 10/s
allowed, _ := limiter.Allow(ctx, "ip:192.168.1.1")
if !allowed {
    // Return 429 Too Many Requests
}
```
