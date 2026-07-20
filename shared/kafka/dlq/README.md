# Dead Letter Queue Subsystem

This package implements routing interfaces to dispatch poisoned or unprocessable event payloads into designated DLQ topics.

---

## 1. Features

- **Failure Handler**: Catches processing errors, writes metadata to telemetry, and routes payloads to audit logs.

---

## 2. API Examples

```go
import "prahari/shared/kafka/dlq"

dlqProd := dlq.NewDLQProducer(kafkaProducer)
dlqHandler := dlq.NewHandler(dlqProd, "worker-alerts-dlq")

err := processMessage(msg)
if err != nil {
    // Forward to DLQ
    _ = dlqHandler.HandleFailure(ctx, msg.Key, msg.Value, err)
}
```
