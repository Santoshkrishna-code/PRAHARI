# Redis Pub/Sub Subsystem

This package wraps Redis Pub/Sub SDK functions to enable low-latency message streaming across microservice events.

---

## 1. Features

- **Decoupled Publisher**: Encodes payload struct to JSON.
- **Blocking Subscriber Loop**: Polling loop for active subscriptions.

---

## 2. API Examples

```go
import "prahari/shared/redis/pubsub"

pub := pubsub.NewPublisher(client)
_ = pub.Publish(ctx, "worker-alerts", alertStruct)
```
