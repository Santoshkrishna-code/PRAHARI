# Hashing Subsystem

This package provides hashing wrappers supporting Bcrypt (for general compatibility) and Argon2id (for modern, highly-hardened setups), alongside password strength validation and cost upgrade trackers.

---

## 1. Features

- **BcryptHasher**: High-performance CPU-bound hashing with configurable cost factors.
- **Argon2Hasher**: Memory-hard password hashing (Argon2id) protecting against GPU/ASIC brute-force attacks.
- **Password Strength Checker**: Validates length, case patterns, numbers, and symbol metrics.
- **Bcrypt Cost Upgrade Tracker**: Detects if stored password hashes require rotation upgrades.

---

## 2. API Examples

### A. Argon2id password hashing
```go
import "prahari/shared/security/hashing"

hasher := hashing.NewArgon2Hasher(65536, 3, 4) // 64MB, 3 iterations, 4 threads
hash, _ := hasher.Hash(ctx, "SecurePass1!")
match, _ := hasher.Verify(ctx, "SecurePass1!", hash)
```

### B. Password strength check
```go
err := hashing.VerifyPasswordStrength("weakpass")
if err != nil {
    // Rejected
}
```
