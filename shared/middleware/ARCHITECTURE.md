# Architecture: HTTP Middleware Pipeline

This document details the request composition lifecycle, trace injection patterns, and safety boundaries governing the `shared/middleware` package.

---

## 1. Request Processing Pipeline

HTTP requests execute through the composition chain in first-to-last (FIFO) order, returning response headers and body outputs back up the chain:

```
    HTTP Request
         │
         ▼
  +--------------+
  |  RequestID   |  <-- Injects X-Correlation-ID
  +--------------+
         │
         ▼
  +--------------+
  |   Recovery   |  <-- Intercepts panics, formats RFC 7807 JSON
  +--------------+
         │
         ▼
  +--------------+
  |    Auth      |  <-- Parses JWT, binds Claims to context
  +--------------+
         │
         ▼
  +--------------+
  | Rate Limit  |  <-- Redis-backed bucket checks
  +--------------+
         │
         ▼
    HTTP Handler
```

---

## 2. Context Safety

- **Scope Isolation**: Request context variables are stored using distinct private `contextKey` types to prevent context key collisions.
- **Panic Protection**: Intercepts runtime panics, capturing stack traces without crashing container instances.
