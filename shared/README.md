# PRAHARI Shared Platform Library

This directory contains the reusable shared modules imported by all microservices within the PRAHARI ecosystem.

---

## Shared Modules Matrix

* **`errors/`**: Unified problem details formatting (RFC 7807) and error boundary translator.
* **`logger/`**: Structured JSON Zap logger with correlation contexts.
* **`config/`**: Decoupled config structural validation.
* **`validation/`**: Struct properties checks (emails, passwords, UUIDs).
* **`aws/`**: Reusable clients for S3, SNS, SQS, SES, Secrets Manager, and KMS.
* **`redis/`**: Caching pooling, health-checks, and Redlocks.
* **`kafka/`** & **`events/`**: MSK integration, CloudEvents, retries, and dead-letter queues.
* **`jwt/`** & **`cognito/`**: JWT parsing handlers and Cognito auth wrappers.
* **`middleware/`**: HTTP filters (recovery, CORS, rate limits, headers).
* **`telemetry/`**: OTel, Prometheus, and readiness probes.
* **`resilience/`**: Circuit Breakers and Retry backoffs.
* **`security/`**: Hashing and symmetric encryption.
* **`utils/`**: Shared helper primitives.

---

## How to Import in Services

All backend services resolve modules through the root `go.work` file. Simply import the shared package inside your service code:

```go
import (
	"prahari/shared/errors"
	"prahari/shared/logger"
)
```
