# Hazard Management Service

The **PRAHARI Hazard Management Service** (`services/hazard/`) acts as the proactive safety engine of the platform, tracking unsafe conditions identification, assessments, mitigation controls implementation, and verification checking.

## Features

- **Hazard Lifecycle State Machine:** Enforces valid state steps transitions flow (`Identified` → `Reported` → `Risk Assessment` → `Mitigation Planning` → `Approval` → `Implementation` → `Verification` → `Closed`).
- **5x5 Risk matrix scoring:** Automatically maps Consequence and Likelihood metrics combinations to risk levels.
- **Risk Mitigation Controls Hierarchy:** Manages Elimination, Substitution, Engineering, Administration, and PPE barriers assignments.
- **Risk policies enforcer:** Auto-escalates critical risk levels into the Incident Service and verifies residual risks are below safety bounds prior to closures.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
