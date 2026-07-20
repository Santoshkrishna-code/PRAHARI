# Secure Random Generator Subsystem

This package provides cryptographically secure random bytes, high-entropy tokens, and native RFC 4122 UUIDv4 identifiers utilizing standard Go entropy sources (`crypto/rand`).

---

## 1. Features

- **Random Bytes**: Raw entropy generators.
- **Base64 & Hex Tokens**: URL-safe high-entropy tokens.
- **Alphanumeric Tokens**: Custom charset tokens for session IDs and API keys.
- **UUIDv4**: Zero-dependency, native standard compliance.

---

## 2. API Examples

### A. Generating Alphanumeric Keys
```go
import "prahari/shared/security/random"

apiKey, _ := random.GenerateAlphanumericToken(32)
```

### B. Generating UUID v4
```go
uuid, _ := random.GenerateUUIDv4()
```
