# Maintenance Management Service

The **PRAHARI Maintenance Management Service** (`services/maintenance/`) acts as the execution engine owning breakdown work orders, corrective/preventive plans setups, technicians assigned, spare parts lists usages, and downtime tracking.

## Features

- **Work Order State Machine:** Enforces valid state steps transitions flow (`Draft` → `Scheduled` → `In Progress` → `Verified` → `Closed`).
- **Recurrent Planning scheduler:** Set up preventive plans cycles frequencies.
- **Safety checks preconditions:** Enforces valid permit verifications checklist approvals before executing maintenance check cards.
- **Cost validation checks policy:** Flags and starts cost approval workflows when estimated values exceed $50k limits.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search vectors)
- **Messaging:** Kafka Event-Driven integration clusters
- **Caching:** Redis Cache
