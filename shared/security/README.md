# PRAHARI Reusable Security Platform SDK

This package provides cryptographic helpers, password hashing adapters, JWT token validation rules, CSRF shields, and secure HTTP header middleware.

---

## 1. Directory Structure

- **`encryption/`**: AES-256 GCM, RSA OAEP, and AWS KMS-backed envelope encryption.
- **`hashing/`**: Password hashing wrappers for Bcrypt and Argon2id.
- **`jwt/`**: Parser and Validator integrating with JWKS public keys providers.
- **`csrf/`**: Double-Submit Cookie pattern middleware.
- **`headers/`**: Injection of HTTP security headers (CSP, HSTS, permissions, cross-origin).
- **`random/`**: Cryptographically secure bytes, UUIDv4s, and high-entropy API key generators.
- **`secrets/`**: AWS KMS adapter wrapping and rotation events mapping.
- **`docs/`**: Threat Models and Cryptographic Architectures.

---

## 2. Developer Integration

### Unified Security Manager Facade
Instead of managing multiple ciphers directly, instantiate the master `Manager` inside your Dependency Injection container:

```go
import (
	"prahari/shared/security"
	"prahari/shared/security/encryption"
	"prahari/shared/security/hashing"
)

// Setup adapters
aesCipher, _ := encryption.NewAESCipher(aes32BytesKey)
bcryptHasher := hashing.NewBcryptHasher(12)

// Build facade manager
secManager := security.NewManager(cfg, aesCipher, bcryptHasher)

// Execute secure actions
ciphertext, err := secManager.Encrypt(ctx, plaintext)
```
