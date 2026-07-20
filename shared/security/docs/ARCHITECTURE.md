# Architecture: Cryptographic and Security Infrastructure

This document details the design patterns, cryptographic algorithms, and validation structures implemented in the `shared/security` library.

---

## 1. Key Design Tenets

- **Algorithm Hardening**: Avoid legacy algorithm configurations (such as MD5, SHA-1, RC4, or RSA without OAEP padding).
- **Decoupled Interfaces**: Encryption algorithms are placed behind mockable interfaces, separating system business rules from concrete implementations.
- **Timing Attack Prevention**: Uses constant-time string checks (`subtle.ConstantTimeCompare`) for passwords and token validations.

---

## 2. Cryptographic Algorithm Matrix

| Security Need | Standard | Parameters | Details |
| :--- | :--- | :--- | :--- |
| Symmetric Encryption | **AES-256 GCM** | 32-byte keys, 12-byte nonces | Appends nonces to output cipher blocks |
| Asymmetric Encryption | **RSA-OAEP** | 2048/4096-bit keys, SHA-256 | Hashing padding protection |
| Password Hashing | **Argon2id** | 64MB memory, 3 iterations, 4 threads | Memory-hard defense against GPU brute forcing |
| Signature Tokens | **JWT** | RSA public keys, dynamic JWKS | Verifies kid header fields |
| CSRF Shields | **Double-Submit** | Secure HttpOnly cookies, lax settings | Header matches cookie |
