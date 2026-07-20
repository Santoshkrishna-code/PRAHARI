# Near Miss Management Service

The **PRAHARI Near Miss Management Service** (`services/nearmiss/`) acts as the safety learning-feedback engine of the platform, tracking near miss reporting, classifications, root cause investigations, CAPA action plans, verifications, and safety lessons logs.

## Features

- **Near Miss Lifecycle State Machine:** Enforces valid state steps transitions flow (`Reported` → `Classified` → `Investigation` → `Root Cause Analysis` → `Corrective Actions` → `Verification` → `Closed` / `Escalated`).
- **Anonymity checks:** Allows anonymous near miss reports without exposing user ID references.
- **Investigation workflows:** Supports methodology tagging (e.g. 5 Whys, Fishbone diagram).
- **Incident Escalation:** serious potential level triggers auto incident log creation.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
