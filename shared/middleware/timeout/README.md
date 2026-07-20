# HTTP Request Timeout Middleware

This package enforces execution duration boundaries across HTTP requests, releasing execution contexts and returning 503 errors.

---

## 1. Features

- **Timeout Boundaries**: Uses `http.TimeoutHandler`.
- **JSON Error formatting**: Standardized returns.

---

## 2. API Examples

```go
import "prahari/shared/middleware/timeout"

// Enforce 10s maximum timeout
router.Use(timeout.Middleware(10 * time.Second))
```
