# JWT Parsing and Signature Validation

This package handles JSON Web Tokens (JWT) parsing, signature checks, claim decodes, and AWS Cognito JWKS key resolution.

---

## 1. Features

- **Consolidated Claims**: Maps standard and custom platform claims (`Email`, `Role`, `TenantID`).
- **Dynamic JWKS cache**: Fetches and caches keys, avoiding Cognito network lookups.
- **Algorithm Confusion Guards**: Restricts validation signing algorithms.
- **Token Extractors**: Helper extractors parsing tokens from headers, query paths, and cookies.

---

## 2. API Examples

### A. Validating Bearer Tokens via JWKS
```go
import "prahari/shared/security/jwt"

// 1. Setup resolver and validator
resolver := jwt.NewKeyResolver(jwksURL, 10*time.Minute)
validator := jwt.NewValidator(resolver, "issuer-url", "app-client-id")

// 2. Validate token
claims, err := validator.Validate(ctx, tokenStr)
if err != nil {
    // Rejected
}
```
