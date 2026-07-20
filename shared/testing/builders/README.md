# HTTP Testing Request Builders Subsystem

This package registers HTTP request builders and recorder runners wrapping Go's standard `httptest` package.

---

## 1. Features

- **Decoupled requests**: Builds request objects using byte arrays or structural payloads.
- **Handler runner**: Simulates route execution loops.

---

## 2. API Examples

```go
import "prahari/shared/testing/builders"

req := builders.BuildJSONRequest("POST", "/register", userPayload)
rec := builders.ExecuteRequest(router, req)
```
