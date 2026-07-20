# OpenTelemetry Tracing HTTP Middleware

This package extracts context span metadata from incoming HTTP requests, creating child spans.

---

## 1. Features

- **OTel Trace extraction**: Reads traceparent mapping headers.

---

## 2. API Examples

```go
import "prahari/shared/middleware/tracing"

router.Use(tracing.Middleware)
```
