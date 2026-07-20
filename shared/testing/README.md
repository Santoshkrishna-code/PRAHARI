# PRAHARI Testing SDK Platform

This package provides hook mocks for AWS (S3, SQS), Redis, and Kafka sync/batch publishers, data fixtures for claims and payloads, and request builders wrapping `httptest`.

---

## 1. Directory Structure

- **`mocks/`**: Hook mocks overriding AWS, Redis, and Kafka.
- **`fixtures/`**: Test credentials and payloads templates.
- **`builders/`**: Request builders wrapping `httptest`.

---

## 2. Usage Blueprint

```go
import (
	"prahari/shared/testing/mocks"
	"prahari/shared/testing/builders"
)

s3Mock := &mocks.MockS3Client{}
req := builders.BuildRequest("GET", "/health", nil)
```
