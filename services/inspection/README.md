# Inspection Management Service

The **PRAHARI Inspection Management Service** (`services/inspection/`) handles safety checklists audits, walkthrough routines, finding severities, and corrective actions (CAPA).

## Features

- **Lifecycle State Machine:** Enforces transitions through `DRAFT` → `SCHEDULED` → `ASSIGNED` → `IN_PROGRESS` → `COMPLETED` → `UNDER_REVIEW` → `APPROVED` → `CLOSED`.
- **Checklist Templates:** Dynamic reuse of blueprints instantiating custom check questionnaires.
- **NCR & CAPA Action lists:** Records Corrective Actions linked to failed walkthrough findings.
- **Auto Incident Trigger:** Raises Incident Service tickets from failed critical item findings.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Clean Domain-Driven Design (DDD)
- **Database:** PostgreSQL (GIN full-text search tsvector)
- **Caching:** Redis Cache
- **Messaging:** Kafka Event-Driven publish/subscribe integrations
