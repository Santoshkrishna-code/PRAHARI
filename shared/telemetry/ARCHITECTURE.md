# Architecture: OpenTelemetry Observability

This document details the configuration lifecycles, trace context propagation, and metric collection rules governing the `shared/telemetry` package.

---

## 1. Trace context propagation

Distributed systems propagate tracing contexts across boundaries via W3C formats:

```
  [ Client request ]
         │
         ▼ (traceparent header)
  +--------------+
  | HTTP Service |  <-- Extracts trace context from headers
  +--------------+
         │
         ▼ (W3C text map propagator propagation)
  +--------------+
  | Kafka Broker |  <-- Links event processing to active span
  +--------------+
         │
         ▼
  +--------------+
  | Database log |  <-- Trace ID bound to database queries
  +--------------+
```

---

## 2. Resources & Exporters

- **Batching**: Span details are batched and exported asynchronously in the background.
- **Resource tags**: Emits service names, environments, versions, and process details with every entry.
