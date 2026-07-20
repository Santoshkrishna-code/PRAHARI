# Walkthrough: Workflow Engine Service

This document summarizes the complete implementation of the central **PRAHARI Workflow Engine Service** under `services/workflow/`.

---

## 1. Directory Structure

- **`cmd/`**: Entry point invoking bootstrap app runners.
- **`migrations/`**: SQL migration scripts defining tables for definitions, instances, and tasks.
- **`internal/domain/`**: DDD Aggregate entities for Definitions, Execution Instances, Explicit States transitions check, Tasks types, Approval engines strategies, SLA timers, and Saga compensations.
- **`internal/application/`**: DSL interpreters, tasks approvals coordinators, and Saga rollback orchestrators.
- **`internal/infrastructure/`**: Postgres SQL adapters, Redis lock pools, and Kafka publishers.
- **`internal/interfaces/`**: REST handlers, gRPC workflow control servers, and Kafka events listener loops.

---

## 2. Dynamic Interpreter DSL Workflow Schema Example

Workflows are interpreted at runtime based on versioned YAML templates declarations (no hardcoded state chains):

```yaml
id: incident-approval-flow
name: "Incident Safety Verification Process"
version: 1

steps:
  - id: check-severity
    type: condition
    next_step: parallel-approvals
  - id: parallel-approvals
    type: approval
    next_step: complete-incident
  - id: complete-incident
    type: webhook
    next_step: ""
```
