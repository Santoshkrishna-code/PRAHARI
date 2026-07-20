# Risk Assessment Management Service

The **PRAHARI Risk Assessment Management Service** (`services/risk/`) acts as the central risk intelligence engine of the platform, tracking process safety registers, inherent/residual 5x5 matrices configs, bow-tie barrier blocks checklists, reviews, approvals, and metrics dashboards.

## Features

- **Risk Assessment Lifecycle State Machine:** Enforces valid state steps transitions flow (`Draft` → `Assessment` → `Review` → `Approval` → `Active` → `Periodic Review` → `Reassessment` → `Closed`).
- **Control Mitigations:** Evaluates barrier effectiveness ratings.
- **Dynamic Risk scoring:** Computes scores in real-time.
- **Incident Escalation:** serious incidents trigger automatic reassessment reviews.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
