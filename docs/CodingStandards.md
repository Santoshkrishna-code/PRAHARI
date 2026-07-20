# Coding Standards & Guidelines

This document details coding standards, safety rules, and implementation patterns required when developing or extending components in the PRAHARI ecosystem.

---

## 1. Core Go Guidelines

- **Context-Awareness**: Every network call, query execution, or client request must accept a `context.Context` as its first parameter.
- **Error Handling**: Use structured app errors wrapping native exceptions. Never discard errors.
- **Panic Protection**: Never spawn raw goroutines without recovery handlers. Always intercept panics using middleware recoveries or local recover blocks.

---

## 2. Resource Management

- **Connection Reuse**: Keep S3, SQS, Redis, and Kafka clients long-lived. Do not reconstruct wrappers inside execution paths.
- **Release Resources**: Always use deferred `Close()` blocks on streams, sockets, and files.
