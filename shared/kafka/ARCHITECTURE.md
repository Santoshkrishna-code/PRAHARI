# Architecture: Kafka Event Streaming SDK

This document details the design patterns, event envelope structures, and partition transaction boundaries governing the `shared/kafka` package.

---

## 1. Key Design Tenets

- **Standardized Messaging**: All events are wrapped in a generic `events.Envelope` containing semver tracking version tags to support schema evolution.
- **Fail-Safe Retries**: Incorporates exponential retry backoffs with random jitter, routing permanently failed entries to a target DLQ topic.
- **Trace Propagation**: Context propagation is enforced at the message level by extracting `traceparent` headers from Kafka properties.

---

## 2. Event Envelope Layout

```
+-------------------------------------------------------+
|                       Envelope                        |
|                                                       |
|   ID: "msg-123"                                       |
|   Type: "UserRegistered"                              |
|   Source: "auth-service"                              |
|   Timestamp: 2026-07-19T20:15:00Z                     |
|   Version: "1.0.0"                                    |
|                                                       |
|   +-----------------------------------------------+   |
|   |                    Payload                    |   |
|   |                                               |   |
|   |   { "user_id": "123", "email": "user@..." }   |   |
|   +-----------------------------------------------+   |
+-------------------------------------------------------+
```
