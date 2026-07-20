# Training & Competency Management Service

The **PRAHARI Training & Competency Management Service** (`services/training/`) acts as the workforce capability and safety engine of the platform, tracking courses, programs, curricula, competency matrices, skills, enrollments, attendance, assessments, examinations, results, and renewals.

## Features

- **Training Lifecycle State Machine:** Enforces valid state steps transitions flow (`Draft` → `Planned` → `Scheduled` → `Enrollment` → `In Progress` → `Assessment` → `Certified` → `Active` → `Renewal`).
- **Enrollment verifications:** Checks trainee eligibility.
- **Competency Dashboard:** Visualizes competency matrix coverage metrics.
- **Automatic Refresher Triggers:** Incidents and audit findings trigger mandatory training.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
