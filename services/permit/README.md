# Permit-to-Work Service (PTW)

The **PRAHARI Permit-to-Work Service** (`services/permit/`) manages the lifecycle of hazardous work authorizations within industrial plants, enforcing strict safety rules, verification checkpoints, and approval workflows.

## Features

- **Lifecycle State Machine:** Enforces transitions through `DRAFT` → `SUBMITTED` → `RISK_ASSESSMENT` → `APPROVAL` → `ISSUED` → `ACTIVE` → `SUSPENDED` → `COMPLETED` → `CLOSED`.
- **Multi-Level Signs:** Validates sequences of sign-offs by roles (Supervisor, Safety Officer, Area Authority, Plant Manager).
- **LOTO Isolations:** Records energy isolations with Lock-Out/Tag-Out verification tracking.
- **Gas Testing:** Evaluates atmospheric test pass/fail results prior to confined area work.
- **Observability:** Custom dashboard counters, tracer spans, and audit changes logging.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Clean/Hexagonal Domain-Driven Design (DDD)
- **Database:** PostgreSQL (with generated full-text search tsvector)
- **Caching:** Redis Cache
- **Messaging:** Event-Driven integration using Kafka

## API Endpoints

Refer to `internal/interfaces/http/router.go` for all HTTP endpoints.
Refer to `internal/interfaces/grpc/server.go` for all gRPC endpoints.
