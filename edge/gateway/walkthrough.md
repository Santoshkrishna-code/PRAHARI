# Walkthrough: API Gateway Service

This document summarizes the complete implementation of the central **PRAHARI API Gateway Service** under `edge/gateway/`.

---

## 1. Directory Structure

- **`cmd/`**: Entry point server main.go.
- **`configs/`**: Route patterns overrides (`development.yaml`).
- **`internal/domain/`**: Route, Upstream, Limits, and Discovery models.
- **`internal/application/`**: Route matching table and HTTP reverse proxy logic.
- **`internal/infrastructure/`**: IAM gRPC validator clients and Kubernetes DNS resolvers.
- **`internal/interfaces/`**: Proxy handlers and administrative admin REST controls.

---

## 2. Gateway Request Pipeline

Incoming HTTP traffic follows this execution filter:

```
  Incoming Request
         │
         ▼
  [ Middleware Chaining ]
  - Request ID Injection
  - CORS Policies checks
  - Logging metrics latencies
  - Context Tracing propagator
         │
         ▼
  [ Routing Matching Table ]
  - Scans prefix paths (e.g. /api/v1/incidents/*)
  - Identifies target upstreams
         │
         ▼
  [ Single Host Reverse Proxy ]
  - Rewrites request host headers
  - Handles prefix stripping (e.g. /api/v1/incidents/create -> /create)
  - Proxies payloads downstream
```
