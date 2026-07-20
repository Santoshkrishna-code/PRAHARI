# Testing Data Fixtures Subsystem

This package registers predefined JWT claims objects and incident JSON payloads, avoiding test arrays duplication.

---

## 1. Features

- **Predefined profiles**: Exposes `NewAdminClaims` and `NewWorkerClaims`.
- **Predefined alerts**: Exposes JSON alerts payloads.

---

## 2. API Examples

```go
import "prahari/shared/testing/fixtures"

claims := fixtures.NewAdminClaims()
```
