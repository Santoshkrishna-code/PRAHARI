# Context Structured Logger

This package wraps Uber Zap to provide optimized JSON logging in production and colorized logs in development, featuring trace propagation.

---

## 1. Features

- **Zap Logger Integration**: High-performance structured logging.
- **Context Extraction**: Auto-extracts Correlation IDs, OpenTelemetry Trace IDs, and Span IDs.
- **Strongly Typed Fields**: Re-exports types for String, Int, Duration, and Error logs.
- **Request Logger**: Standardized HTTP request tracking.

---

## 2. API Reference & Code Examples

### A. Initializing the logger
```go
import "prahari/shared/logger"

// Standard development console logger
log := logger.NewDefault("auth-service")

// Custom options loader
opts := logger.Options{
    Env:         "production",
    Level:       "info",
    Format:      logger.FormatJSON,
    OutputPaths: []string{"stdout"},
    ServiceName: "permit-service",
}
log, err := logger.New(opts)
```

### B. Standard logging with context
```go
// Bind correlation context
ctx := logger.WithCorrelationID(context.Background(), "corr-12345")

// Child logger logs automatically mapping correlation_id field
log.WithContext(ctx).Info("Processing worker vitals stream")
```

### C. Logging HTTP requests
```go
// Log request
logger.LogRequest(log, request, status, size, latency)
```
