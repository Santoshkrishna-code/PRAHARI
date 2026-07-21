# PRAHARI AI Platform: Coding Standards

This document establishes the architecture rules and quality standards for Python codebase development.

---

## 1. Static Typing & Type Safety
- **Type Annotations**: All function parameters, return values, and variables must be fully typed.
- **Strict Mode**: MyPy must run in strict mode (`strict = true`). Implicit `Any` declarations are prohibited.
- **Type Aliases**: Use descriptive type aliases for complex generic definitions.

---

## 2. Layer Boundaries (Clean Architecture)
- **Dependency Inversion**: Outer layers (HTTP Handlers, CLI Commands) can depend on inner layers (Core, Domain), but inner layers must never depend on outer layers.
- **Entity Purity**: Inner domain logic must operate only on plain Pydantic models or domain objects. SQLAlchemy ORM models must remain isolated within the Infrastructure layer.
- **DI Container**: Do not construct instances inline. Resolve all database, cache, and client dependencies via the `ProviderRegistry` container.

---

## 3. Telemetry & Observability
- **Log Levels**: Use context-aware structured JSON logging. Debug logs are strictly for developer trace context.
- **Trace Context**: Always propagate trace contexts across asynchronous functions. Ensure trace correlation IDs are attached to log payloads.
- **Alert Targets**: Never log sensitive data (passwords, tokens, personal identifiers) to the stdout stream.
