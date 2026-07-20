# Contractor Management Service

The **PRAHARI Contractor Management Service** (`services/contractor/`) acts as the workforce compliance engine managing contractor companies, worker profiles onboarding, safety certifications, medical clearance check status, site gate entry access, and safety violations.

## Features

- **Contractor Lifecycle State Machine:** Enforces valid state steps transitions flow (`Registered` → `Document Verification` → `Safety Training` → `Medical Clearance` → `Site Induction` → `Approved` → `Active`).
- **Site access controls:** Validates worker safety inductions and access badge credentials at gate entry points.
- **Permit Eligibility check:** Evaluates crew members qualification criteria before work allocations.
- **Violations Performance tracking:** Logs safety warnings, penalizes score card metrics, and blacklists non-compliant crew members.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Hexagonal / Domain-driven design
- **Database:** PostgreSQL (GIN search vectors index)
- **Messaging:** Kafka Event-Driven integration
- **Caching:** Redis Cache
