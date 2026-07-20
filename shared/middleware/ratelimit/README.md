# Redis Rate Limiting HTTP Middleware

This package integrates Redis Token-Bucket rate limiters into HTTP pipelines, checking and rejecting client quota overflows.

---

## 1. Features

- **Custom client keys**: Groups requests by IP, authorization token, or route.
- **Problem Details Response**: Formats 429 Too Many Requests payload.

---

## 2. API Examples

```go
import "prahari/shared/middleware/ratelimit"

rl := ratelimit.NewRateLimiter(redisClient, 50, 5.0)

// Limit based on client IP
limiterMiddleware := rl.Middleware(func(r *http.Request) string {
    return r.RemoteAddr
})

router.Use(limiterMiddleware)
```
