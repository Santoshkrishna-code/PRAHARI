# PRAHARI Structured Logging Standards
## Enterprise JSON Logging Schema

This document details the structured logging standard mandatory for all microservices, edge components, and AI agents within the PRAHARI platform.

---

## 1. Core Principles

1. **JSON Output**: All logs in environments above local development MUST be written in JSON format to stdout. Plain text logging is forbidden in staging and production.
2. **Non-Blocking Execution**: Logging must not impact API latency. Log calls should be asynchronous or highly optimized (using buffered streams).
3. **No Secret Leaks**: Loggers must filter out sensitive parameters (passwords, JWT tokens, AWS access keys, personally identifiable vitals).

---

## 2. Standard JSON Log Schema

Every log line must be a single flat JSON object containing these core fields:

| Field | Type | Description | Example |
| :--- | :--- | :--- | :--- |
| **`timestamp`** | string | RFC3339 format, always in UTC. | `"2026-07-19T05:17:14.002Z"` |
| **`level`** | string | Log severity (DEBUG, INFO, WARN, ERROR, FATAL). | `"INFO"` |
| **`service`** | string | Name of the originating microservice. | `"permit-service"` |
| **`environment`**| string | Host env (dev, staging, production). | `"production"` |
| **`trace_id`** | string | OpenTelemetry trace ID for distributed tracing. | `"4bf92f3577b34da6a3ce929d0e0e4736"` |
| **`span_id`** | string | OpenTelemetry span ID for tracing. | `"00f067aa0ba902b7"` |
| **`message`** | string | Human-readable log summary. | `"Permit approved successfully"` |
| **`error`** | object | Error context block (only on ERROR/FATAL). | *See error payload schema* |

### Log Example (Success)
```json
{
  "timestamp": "2026-07-19T10:48:32.190Z",
  "level": "INFO",
  "service": "permit-service",
  "environment": "production",
  "trace_id": "8fb3ac1d45bc88efb09c8191981cb02c",
  "span_id": "c19a9d0a1b028da8",
  "message": "Permit issued to worker",
  "permit_id": "ptw-99201",
  "worker_id": "usr-8819"
}
```

### Log Example (Error with Context)
```json
{
  "timestamp": "2026-07-19T10:49:01.002Z",
  "level": "ERROR",
  "service": "permit-service",
  "environment": "production",
  "trace_id": "8fb3ac1d45bc88efb09c8191981cb02c",
  "span_id": "c19a9d0a1b028da8",
  "message": "Permit creation failed",
  "error": {
    "message": "worker has expired safety induction credentials",
    "code": "EXPIRED_CREDENTIALS",
    "stacktrace": "main.go:42: repository.CreatePermit() -> permit.go:121: checkWorkerInductions()"
  }
}
```

---

## 3. Approved Libraries

* **Go Services**: Use **`uber-go/zap`** or **`rs/zerolog`** configured for JSON encoder output.
* **Python Services**: Use **`structlog`** or python standard library's JSON formatter wrappers.
* **Node/React**: Backend CLI tasks must use **`pino`**. Frontends must use console logs only under local dev, filtered out in production builds.
