# Kafka Producer Subsystem

This package wraps segmentio/kafka-go Writer properties to support synchronous, asynchronous, and batched message publications.

---

## 1. Features

- **Sync Publisher**: Blocks execution awaiting broker ACKs.
- **Async Publisher**: Writes messages asynchronously, returning immediately.
- **Batching**: Sends multiple records in a single transactional request.

---

## 2. API Examples

```go
import "prahari/shared/kafka/producer"

p, _ := producer.NewProducer(kafkaConfig)
defer p.Close()

// Sync publish
_ = p.PublishSync(ctx, "worker-telemetry", []byte("w-1"), []byte("vitals-data"))
```
