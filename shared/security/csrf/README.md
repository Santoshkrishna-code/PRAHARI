# CSRF Protection Subsystem

This package implements cross-site request forgery protection using the double-submit cookie pattern.

---

## 1. Features

- **Double-Submit Cookie**: Compares the secure cookie value against the incoming request header.
- **Timing attack protection**: Uses `subtle.ConstantTimeCompare` to check tokens.
- **Secured HTTP cookies**: Generates secure cookies containing Lax attributes and HttpOnly restrictions.

---

## 2. API Examples

### A. Injecting the CSRF Middleware
```go
import "prahari/shared/security/csrf"

// Add to HTTP router filters list
router.Use(csrf.Middleware)
```
