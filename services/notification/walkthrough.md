# Walkthrough: Notification Service

This document summarizes the complete implementation of the central **PRAHARI Notification Service** under `services/notification/`.

---

## 1. Directory Structure

- **`cmd/`**: Entry point invoking bootstrap app runners.
- **`migrations/`**: SQL migration scripts defining tables for logs, templates, and preferences.
- **`internal/domain/`**: DDD Aggregate entities for Payload details, Localized templates, Pluggable channels, Preferences rules, Delays schedules, Retry policies, and Delivery lifecycles.
- **`internal/application/`**: Compiled templates renderers, channel flow dispatchers, and exponential backoff retry managers.
- **`internal/infrastructure/`**: Postgres stores, Email/SMS client adapters, and IAM contact detail resolvers.
- **`internal/interfaces/`**: REST handlers, gRPC send server gateways, and Kafka events listener loops.

---

## 2. Pluggable Channel Flow Architecture

Outbound dispatches are abstracted away from business services logic:

```
  Kafka Trigger Event (e.g. approval.requested)
                      │
                      ▼
            [ Preference Evaluator ]
            - Checks opt-out rules on user profiles
                      │
                      ▼
             [ Template Compiler ]
             - Renders HTML/Markdown structures
                      │
                      ▼
             [ Channel Flow Dispatcher ]
             - Matches channel keys (email, sms, push)
             - Executes pluggable Channel adapters
                      │
                      ▼
             [ Delivery Lifecycles ]
             - Updates state statuses on database tables (QUEUED, SENT, READ)
```
