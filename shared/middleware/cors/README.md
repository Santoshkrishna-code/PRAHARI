# CORS Protection Middleware

This package provides Cross-Origin Resource Sharing settings, appending safety headers and intercepting browser preflight OPTIONS requests.

---

## 1. Features

- **Standard CORS rules**: Enforces allowed origins, headers, and credentials rules.
- **Preflight Interceptor**: Automatically returns `204 No Content` on OPTIONS paths.

---

## 2. API Examples

```go
import "prahari/shared/middleware/cors"

opts := cors.DefaultOptions()
router.Use(cors.Middleware(opts))
```
