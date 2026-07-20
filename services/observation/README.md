# Safety Observation Management Service

The **PRAHARI Safety Observation Management Service** (`services/observation/`) acts as the Behavior-Based Safety (BBS) engine of the platform, tracking safe/unsafe workplace behavior reports, coaching feedback dialogues sessions, positive rewards recognitions, and effectiveness evaluations.

## Features

- **Observation Lifecycle State Machine:** Enforces valid state steps transitions flow (`Observed` → `Recorded` → `Reviewed` → `Coaching` → `Follow-up` → `Verified` → `Closed`).
- **Coaching interventions:** Evaluates safety behaviors improvements metrics scores.
- **Positive recognitions:** Rewards safe acts.
- **Unsafe behaviors escalation:** Converts critical unsafe observations to proactive hazard records.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
