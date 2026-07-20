# Compliance Management Service

The **PRAHARI Compliance Management Service** (`services/compliance/`) acts as the statutory governance engine of the platform, tracking compliance registers, regulations catalogs, standards clauses checklist requirements, evidence logs, reviews, comments, and metrics dashboards.

## Features

- **Compliance Lifecycle State Machine:** Enforces valid state steps transitions flow (`Draft` → `Defined` → `Assigned` → `Evidence Collection` → `Review` → `Compliant` → `Monitoring` → `Renewal`).
- **Evidence verification:** Checks document uploads integrity rules.
- **Expiry alerts:** Tracks expiring certifications.
- **Incident integrations:** Triggering closed incident compliance requirements audit checks.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
