# Platform Events Subsystem

This package defines structured message Envelopes, serializers, and version compatibility helpers.

---

## 1. Features

- **Standardized Envelope**: Injects standard headers (ID, Type, Source, Timestamp, SchemaVersion).
- **Semver checking**: Compares major versions to reject breaking messages.

---

## 2. API Examples

```go
import "prahari/shared/kafka/events"

// Wrap event data
envBytes, _ := events.WrapEvent("msg-123", "WorkerVitalsSpike", "telemetry-service", "1.0.0", payloadStruct)

// Unwrap event data
var payload VitalsPayload
env, _ := events.UnwrapEvent(envBytes, &payload)
```
