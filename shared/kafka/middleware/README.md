# Kafka Middleware Interceptors

This package implements logging, tracing, and metric counting middlewares wrapping Kafka event processing handlers.

---

## 1. Features

- **Logging**: Captures latency and success rates.
- **OTel Tracing**: Resolves `traceparent` headers, propagating execution traces.
- **Metrics**: Exposes atomic success/failure counters.

---

## 2. API Examples

```go
import "prahari/shared/kafka/middleware"

tracker := &middleware.MetricsTracker{}
handler := middleware.LoggingMiddleware(actualHandler)
metricsHandler := tracker.MetricsMiddleware(handler)

// Run reader loop wrapping handler
_ = c.StartConsumer(ctx, metricsHandler)
```
