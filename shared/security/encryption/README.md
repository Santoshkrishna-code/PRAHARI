# Encryption Subsystem

This package provides cryptographic helpers for AES-256 GCM symmetric encryption, RSA-OAEP asymmetric encryption, and AWS KMS-backed Envelope Encryption.

---

## 1. Features

- **AES-256 GCM**: Secure symmetric encryption utilizing 12-byte random initialization vectors.
- **RSA-OAEP**: PKCS1/PKIX PEM serialization, RSA public/private key generation, and SHA-256 padding encryptions.
- **Envelope Encryption**: Data key generation wrapped by AWS KMS client wrappers, preventing large payload encryption performance bottlenecks.

---

## 2. API Examples

### A. AES-256 GCM encryption
```go
import "prahari/shared/security/encryption"

cipher, _ := encryption.NewAESCipher(key32Bytes)
ciphertext, _ := cipher.Encrypt(ctx, plaintext)
```

### B. AWS KMS Envelope encryption
```go
// kmsClient implements KMSAPI
envelope := encryption.NewEnvelopeCipher(kmsClient, "alias/master-key")
sealedPayload, _ := envelope.Encrypt(ctx, confidentialVitals)
```
