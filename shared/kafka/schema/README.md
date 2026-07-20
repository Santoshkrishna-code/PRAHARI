# Schema Validation Subsystem

This package provides structure constraint validations for event payloads.

---

## 1. Features

- **Field checking**: Enforces structural attributes validation tags (emails, ports, ranges).

---

## 2. API Examples

```go
import "prahari/shared/kafka/schema"

validator := schema.NewSchemaValidator()
err := validator.Validate(somePayload)
if err != nil {
    // Bad payload: reject or drop
}
```
