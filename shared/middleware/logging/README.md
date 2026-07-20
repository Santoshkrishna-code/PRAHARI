# HTTP Request Logging Middleware

This package wraps HTTP handlers to record latencies, status codes, paths, and correlation identifiers to structured system logs.

---

## 1. Features

- **Status Code Interceptor**: Wraps HTTP writers.
- **Latency tracker**: Captures durations.

---

## 2. API Examples

```go
import "prahari/shared/middleware/logging"

router.Use(logging.Middleware)
```
