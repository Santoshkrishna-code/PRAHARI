# PRAHARI Shared Middleware Platform SDK

This package provides a composition chain, request context helpers, JWT authenticator, RBAC check validations, latency loggers, panic recoveries, duration timeouts, CORS policies, rate limiters, and idempotency cache filters.

---

## 1. Directory Structure

- **`requestid/`**: Correlation ID header injector.
- **`logging/`**: Latencies and statuses request logger.
- **`recovery/`**: Panic recovery handling, formatting RFC 7807 responses.
- **`timeout/`**: Service Unavailable request duration limits.
- **`cors/`**: CORS header configurations.
- **`auth/`**: Token claims extraction and RBAC.
- **`ratelimit/`**: Redis token bucket HTTP rate limits.
- **`idempotency/`**: Redis locking duplicate request playbacks.
- **`tracing/`**: OpenTelemetry child span tracers.
- **`metrics/`**: Throughput and latency telemetry.

---

## 2. Usage Blueprint

```go
import (
	"prahari/shared/middleware"
	"prahari/shared/middleware/requestid"
	"prahari/shared/middleware/recovery"
)

// Setup composition chain
pipeline := middleware.New(
	requestid.Middleware,
	recovery.Middleware,
)

// Attach router
http.ListenAndServe(":8080", pipeline.Then(myRouter))
```
