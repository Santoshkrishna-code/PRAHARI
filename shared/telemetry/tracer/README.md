# OpenTelemetry Tracer Subsystem

This package wraps standard trace span recorders to simplify telemetry integrations.

---

## 1. Features

- **Decoupled Tracer**: Starts span contexts.

---

## 2. API Examples

```go
import "prahari/shared/telemetry/tracer"

tr := tracer.NewTracer("auth-service")
ctx, span := tr.Start(ctx, "DB:GetUser")
defer span.End()
```
