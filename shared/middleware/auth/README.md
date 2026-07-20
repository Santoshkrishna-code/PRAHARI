# Authentication and Authorization Middleware

This package provides token authorization checkpoints and role-based permissions validations filters.

---

## 1. Features

- **JWT Authenticator**: Decodes and binds token claims.
- **RBAC Validator**: Enforces permissions rules across routes.

---

## 2. API Examples

```go
import "prahari/shared/middleware/auth"

// 1. Token validation filter
authenticator := auth.NewAuthenticator(validator)
router.Use(authenticator.Middleware)

// 2. Enforce Role checks on specific endpoints
router.Handle("/admin/settings", auth.AuthorizeRole("Admin")(adminHandler))
```
