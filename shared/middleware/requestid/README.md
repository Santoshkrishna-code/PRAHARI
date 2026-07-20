# Request Correlation ID Middleware

This package implements context-aware request correlation identifiers extraction and generation filters.

---

## 1. Features

- **Correlation ID extraction**: Reads `X-Correlation-ID` header.
- **UUID generator**: Generates RFC 4122 compliant UUIDs.

---

## 2. API Examples

```go
import "prahari/shared/middleware/requestid"

router.Use(requestid.Middleware)
```
