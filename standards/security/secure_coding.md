# PRAHARI Secure Coding Standards
## Enterprise Security Guidelines

Security is not a final step; it is a foundational component of every line of code written for the PRAHARI platform. All engineers must strictly follow these rules.

---

## 1. Secrets Management

1. **Zero Hardcoded Secrets**: Under no circumstances should database passwords, API keys, private keys, or credentials be committed to this repository.
2. **Local Environment**: Local secrets belong only in a git-ignored `.env` file.
3. **AWS Deployments**: Secrets must reside in **AWS Secrets Manager** and be injected into EKS pods at runtime using the **AWS Secrets Store CSI Driver** as environment variables or mounted files.

---

## 2. Input Validation & Sanitization

1. **Sanitize Ingress**: All inputs from public clients (REST/gRPC APIs, WebSocket messages) must be validated before triggering business processes.
   - Go: Use struct validation tags via `go-playground/validator` or `ozzo-validation`.
   - Python: Enforce schemas via **Pydantic v2** models.
2. **Prevent Injection**: Use parametrized queries (e.g., Prepared Statements) for SQL transactions. Never concatenate raw input strings into database queries.
3. **Escape Dynamic Output**: All inputs rendered in UI components must be HTML-escaped by default to prevent Cross-Site Scripting (XSS).

---

## 3. JWT Verification & Access Control (RBAC/ABAC)

1. **Token Signature Check**: Services must cryptographically verify Cognito JWT signatures using the User Pool JSON Web Key Sets (JWKS).
2. **Expiration Audits**: Handlers must check the `exp` claim and immediately reject expired tokens.
3. **Authorization Check (RBAC & ABAC)**:
   - **RBAC**: Verify user roles (`groups` claim from Cognito) like `SafetySupervisor`, `Worker`, `PlantManager`.
   - **ABAC**: Validate attribute context (e.g., verifying if the supervisor's `plant_id` matches the permit's `plant_id` before allowing modifications).

---

## 4. Cryptographic Standards

1. **In-Transit Encryption**: All communication must use TLS 1.3 (or minimum TLS 1.2).
2. **At-Rest Encryption**:
   - Files uploaded to S3 must use Server-Side Encryption (SSE-KMS) with Customer Managed Keys.
   - Sensitive database columns (e.g., worker personal data, medical vitals history) must be encrypted before database write using AWS KMS SDK (Envelope Encryption).
