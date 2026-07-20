# Panic Recovery Middleware

This package intercepts runtime panics in the request execution flow, preventing process crash loops and formatting RFC 7807 problem details.

---

## 1. Features

- **Recovery hook**: Catches Go panic statements.
- **Problem Details Response**: Formats 500 error messages.

---

## 2. API Examples

```go
import "prahari/shared/middleware/recovery"

router.Use(recovery.Middleware)
```
