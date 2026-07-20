# Fallback Execution Subsystem

This package wraps primary function execution flows, executing backup recovery commands on primary errors.

---

## 1. Features

- **Decoupled Fallback**: Exposes `Execute(primary, backup)`.

---

## 2. API Examples

```go
import "prahari/shared/resilience/fallback"

res, _ := fallback.Execute(
    func() (interface{}, error) { return queryMainDatabase() },
    func(err error) (interface{}, error) { return queryBackupCache() },
)
```
