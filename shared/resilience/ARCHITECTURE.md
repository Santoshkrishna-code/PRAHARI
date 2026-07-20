# Architecture: Resilience Patterns

This document details the configuration lifecycles, circuit states, retry jitters, and bulkhead semaphore bounds checks governing the `shared/resilience` package.

---

## 1. Key Design Tenets

- **Tripping Safeguards**: Prevents cascading faults. When failures cross thresholds (e.g. 50%), the circuit opens, blocking database/API loads.
- **Resource Containment**: Bulkhead semaphores restrict concurrent executions to prevent memory/thread pool exhaustion.
- **Graceful degradation**: Exposes fallback executors, ensuring backup operations continue during network/database outages.

---

## 2. CB State transitions

```
          +-------------------------+
          |         CLOSED          | <=================+
          +-------------------------+                   |
            |                     ^                     |
            | (failure ratio      | (success count      |
            |  crossed)           |  reached)           |
            v                     |                     |
          +-------------------------+                   |
          |          OPEN           |                   |
          +-------------------------+                   |
            |                                           |
            | (cooldown time                            |
            |  elapsed)                                 |
            v                                           |
          +-------------------------+                   |
          |        HALF-OPEN        | ------------------+
          +-------------------------+ (on failure: back
                                       to OPEN)
```
