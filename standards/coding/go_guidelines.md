# PRAHARI Go Coding Standards
## Enterprise Go Programming Guidelines

This document outlines the coding standards, naming conventions, and best practices for writing clean, production-grade Go microservices in the PRAHARI platform.

---

## 1. Code Layout & Architecture

Every Go microservice MUST follow the Clean Architecture / Ports & Adapters layout:
- **`cmd/`**: Simple main entrypoint. Avoid putting business logic here.
- **`internal/domain/`**: Pure business models, core domain entities, and interface declarations for ports (repositories, services).
- **`internal/service/`**: Core business orchestrations implementing the domain service interfaces.
- **`internal/api/`**: Adapters for ingress (REST controllers, gRPC handlers).
- **`internal/repository/`**: Adapters for egress databases (SQL, Redis, external APIs).

---

## 2. Naming Conventions

1. **Package Names**: Short, single-word, lowercase, noun, singular. Avoid generic packages like `utils`, `helper`, or `common`.
2. **Interfaces**: Interface names should end with `er` if they define a single method (e.g., `Reader`, `Writer`, `PermitValidator`).
3. **Variables**: Short name inside local functions, descriptive name for package/struct level. Use camelCase.
4. **Acronyms**: Keep acronyms consistent casing: `userID` (NOT `userId`), `apiGateway` (NOT `ApiGateway`).

---

## 3. Error Handling Guidelines

1. **Always Return Errors**: Return `error` as the last parameter of a function if it can fail.
2. **Error Wrapping**: Wrap errors with context when returning up the call stack to ease troubleshooting.
   ```go
   if err != nil {
       return nil, fmt.Errorf("failed to fetch permit %s: %w", permitID, err)
   }
   ```
3. **Check Errors Immediately**: Handle errors immediately rather than deferring them. Fail fast.
4. **Panic Isolation**: Do NOT use `panic` for regular error handling. Panics are reserved for unrecoverable startup conditions.

---

## 4. Concurrency & Telemetry

1. **Goroutine Lifecycles**: Never spawn a goroutine without knowing how and when it will terminate. Avoid memory leaks.
2. **Context Propagation**: Always propagate `context.Context` as the first argument in all network calls, database queries, and parallel workers to support cancellation and tracing propagation.
3. **Mutex Safety**: Protect shared mutable state using sync tools. Prefer channels for coordination.
4. **Channel Sizes**: Always document why buffered channels are allocated with a specific size.

---

## 5. Testing Requirements

- Unit tests must write clean assertions using `stretchr/testify/assert` or `require`.
- Mock external layers (e.g., database, external services) using interfaces and generated mocks (via Mockery).
- Files must be named `*_test.go`.
