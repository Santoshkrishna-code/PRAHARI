# Architecture Blueprint: PRAHARI Platform SDK

This document describes the design patterns, package hierarchy, and dependency directions that govern the PRAHARI Shared Platform SDK.

---

## 1. Global Package Hierarchy

The SDK is organized as a set of decoupled, layered modules. The dependency direction flows **downwards** toward the standard library:

```
                    Business Microservices
                              │
                              ▼
┌───────────────────────────────────────────────────────────────┐
│                  Shared Platform SDK                          │
├───────────────────────────────────────────────────────────────┤
│ Middleware  (FIFO request chains, Auth, Rate limiters)        │
│ Telemetry   (OpenTelemetry TracerProviders & Exporters)       │
│ Resilience  (Circuit breakers, exponential retries)           │
├───────────────────────────────────────────────────────────────┤
│ AWS SDK     (S3 object stores, SQS queue brokers)             │
│ Redis       (Distributed caching locks & buckets)              │
│ Kafka       (Producers, Consumer Groups, DLQs)                │
├───────────────────────────────────────────────────────────────┤
│ Config      (SSM & Secrets Manager layers overrides)          │
│ Security    (Symmetric KMS encryption, JWT signatures)        │
│ Errors      (RFC 7807 problem details error mapping)          │
│ Logger      (Context-aware zap logger)                        │
└───────────────────────────────────────────────────────────────┘
                              │
                              ▼
           AWS SDK • Redis • Segmentio Kafka • OTel
```

---

## 2. Structural Layering Rules

- **Clean Boundaries**: Domain layers must not refer directly to database sockets or messaging clients. Infrastructure clients must always be injected as interface abstractions.
- **Acyclic Flows**: Circular package references are strictly prohibited. Packages must be layered horizontally.
- **W3C Propagation**: Context propagation must be carried across HTTP, gRPC, and Kafka brokers using standard W3C propagators.
