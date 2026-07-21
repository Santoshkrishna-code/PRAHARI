# PRAHARI Platform: Low Level Design (LLD) Guidelines

## 1. Tactical DDD and Layered Architecture
All microservices in PRAHARI are written in Go and structured following Clean Architecture and Domain-Driven Design (DDD) principles. The system decouples business logic from external frameworks, databases, and network transports.

```
┌─────────────────────────────────────────────────────────────┐
│                       Microservice Layers                   │
├─────────────────────────────────────────────────────────────┤
│ 1. Interfaces (http/handler, grpc/server, events/listener)  │
├───────────────────┬─────────────────────────────────────────┤
│ 2. Application    │ Orchestrates use cases, manages txs     │
├───────────────────┼─────────────────────────────────────────┤
│ 3. Domain         │ Contains entities, aggregates, rules    │
├───────────────────┼─────────────────────────────────────────┤
│ 4. Infrastructure │ DB adapters, Kafka clients, external API│
└───────────────────┴─────────────────────────────────────────┘
```

---

## 2. Directory Structure Conventions
Each microservice must conform to the following folder topology:

```
<service-root>/
├── cmd/
│   └── server/
│       └── main.go           # Application entrypoint & dependency injection
├── migrations/               # SQL migration scripts
├── Dockerfile                # Multi-stage production container definition
├── Makefile                  # Build, test, and run tasks
├── internal/
│   ├── bootstrap/            # Server setup, database connections, middleware
│   ├── application/          # Use cases, commands, queries, and handlers
│   ├── domain/               # Aggregates, Entities, Value Objects, Domain Policies
│   ├── infrastructure/       # Database implementations, Redis cache, Kafka adapters
│   └── interfaces/           # Transport layers (HTTP controllers, gRPC, Event listeners)
```

---

## 3. CQRS Implementation Pattern
To maintain scalability, microservices segregate write operations (Commands) from read operations (Queries):
1. **Commands**: Triggered by HTTP POST/PUT/DELETE or Kafka messages. They invoke domain aggregates, execute business rules, commit state via repositories, and dispatch events.
2. **Queries**: Fetch data optimized for presentation. Queries can bypass domain rules and execute direct database queries or read from caches.

---

## 4. Shared SDK Utilization
Microservices must utilize packages from the `prahari/shared` library:
- **Resilience**: Import `prahari/shared/resilience` to implement circuit breakers, retries, and fallbacks.
- **Security**: Use `prahari/shared/security` for KMS envelope encryption, hashing, and token parsing.
- **Telemetry**: Initialize traces and metrics using providers from `prahari/shared/telemetry`.

---

## 5. Microservice Low-Level Designs
Refer to the individual service LLD specifications for domain models, schemas, and API definitions:
- [Chemical Management Service LLD](file:///Users/Santosh/PRAHARI/docs/04_LLD/chemical_service.md)
- [Process Hazard Analysis (PHA) Service LLD](file:///Users/Santosh/PRAHARI/docs/04_LLD/pha_service.md)
- [Analytics & Incident Intelligence Service LLD](file:///Users/Santosh/PRAHARI/docs/04_LLD/analytics_service.md)
