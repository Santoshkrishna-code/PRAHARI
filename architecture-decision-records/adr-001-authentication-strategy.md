# ADR-001: Platform Authentication Strategy

## Status
Approved

## Context
PRAHARI is an enterprise-grade Industrial Safety Intelligence Platform deploying to AWS. The authentication and identity management subsystem must:
1. Scale to support thousands of concurrent workers, sensors, and supervisors.
2. Comply with strict security frameworks (MFA, password complexities, brute-force protections, account lockout rules).
3. Issue cryptographically signed tokens (JWTs) compatible with API Gateways and microservices.
4. Integrate with external identity systems (SAML/OIDC) for enterprise customers.

---

## Decision
We will use **Amazon Cognito User Pools** as the primary Identity Provider (IdP) and user credentials directory for the PRAHARI platform.

The Auth Service microservice will act as a facade, exposing clean REST endpoints (`/register`, `/login`, `/refresh`, `/profile`) and translating domain schemas into Cognito API calls via the AWS SDK.

---

## Alternatives Considered

### 1. Self-hosted Keycloak
* **Pros**: Highly customizable, open-source, avoids AWS vendor lock-in.
* **Cons**: Introduces high operational overhead (requires managing HA PostgreSQL cluster, Helm deployment, JVM tuning, patching vulnerabilities, and backup orchestration).

### 2. Custom Identity Implementation (PostgreSQL + JWT)
* **Pros**: Complete control over schemas and databases.
* **Cons**: Reinventing cryptography, session handling, token rotation, and MFA flows. High security risk of software vulnerabilities.

---

## Consequences

* **Pros**:
  - Zero operational overhead for credential storage and security audits.
  - Built-in multi-factor authentication (MFA) and adaptive risk-based authentication.
  - Out-of-the-box support for SSO (SAML 2.0 / OpenID Connect).
  - AWS Key Management Service (KMS) handles password hashing and token signatures.
* **Cons**:
  - Hard dependency on AWS Cognito APIs, creating vendor lock-in.
  - Latency impact of issuing external network calls to Cognito endpoint (mitigated by token verification caching).
