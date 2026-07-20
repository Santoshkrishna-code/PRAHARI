# Secure HTTP Headers Subsystem

This package implements standard security hardening headers via HTTP middlewares.

---

## 1. Features

- **Standard Headers Injector**: Clickjacking Denial, MIME-nosniff, HSTS HTTPS enforcement, XSS filters, and strict referrers.
- **CSP Builder**: Custom Content-Security-Policies assembler.
- **Permissions Builder**: Secure browser API permissions restrictions.
- **Cross-Origin Isolations**: Pre-configures COEP, COOP, and CORP configurations.

---

## 2. API Examples

### A. Registering the Security Headers Middleware
```go
import "prahari/shared/security/headers"

// Using default production values
opts := headers.DefaultOptions()
router.Use(headers.Middleware(opts))
```

### B. Customizing CSP Rules
```go
csp := headers.NewCSPBuilder().
    Add("default-src", "'self'").
    Add("script-src", "'self'", "https://apis.google.com").
    Build()

opts := headers.DefaultOptions()
opts.CSP = csp
router.Use(headers.Middleware(opts))
```
