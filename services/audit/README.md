# Audit Management Service

The **PRAHARI Audit Management Service** (`services/audit/`) acts as the independent assurance engine of the platform, tracking audits, audit programs, plans, checklists, findings, non-conformities (NCR), and corrective actions plans (CAPA).

## Features

- **Audit Lifecycle State Machine:** Enforces valid state steps transitions flow (`Draft` → `Planned` → `Scheduled` → `In Progress` → `Evidence Collection` → `Review` → `Approved` → `Closed`).
- **Evidence verifications:** Checks document uploads integrity rules.
- **Findings tracking:** Logs observations, recommendations, and NCRs.
- **CAPA correctives:** Schedules and verifies corrective actions.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
