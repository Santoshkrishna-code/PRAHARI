# Threat Model and Mitigations

This document outlines key security threats mapped across the platform and how the `shared/security` subsystem provides mitigations.

---

## 1. Mapped Threat Vectors & Mitigations

### A. Cryptographic Key Leaks (OWASP A02:2021)
- **Threat**: Attackers accessing DB passwords or credentials.
- **Mitigation**: Using **AWS KMS Envelope Encryption**. Encrypted secrets are stored in code repositories, but can only be unsealed at runtime by containers possessing appropriate IAM roles to call KMS.

### B. CSRF Attacks (OWASP A01:2021)
- **Threat**: Malicious websites firing state-changing HTTP commands on behalf of an authenticated user.
- **Mitigation**: Enforcing the **Double-Submit Cookie** pattern. State-mutating commands (POST, PUT, DELETE) must contain matching CSRF tokens in both a secure cookie and the `X-CSRF-Token` header.

### C. Brute-Force Password Cracking
- **Threat**: Offline dictionary attacks on leaked database user hashes.
- **Mitigation**: Using **Argon2id** password hashing. Because Argon2id requires memory-hard computations (64MB memory per hash), hardware cracking tools (GPUs and ASICs) are rendered economically unviable.

### D. Session Spoofing and Replay
- **Threat**: Hijacking or editing client JWT fields.
- **Mitigation**: **JWKS RSA Token Validation**. RSA signatures are verified using public keys dynamically cached from the auth provider's JWKS endpoint. Audience (`aud`) and expiration (`exp`) claims are strictly enforced.
