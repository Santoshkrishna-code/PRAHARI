# Kafka Consumer Subsystem

This package wraps segmentio/kafka-go Reader properties to run commit-on-success partition loops and consumer groups.

---

## 1. Features

- **Read-Commit Loop**: Fetches partition messages and commits offsets on success.
- **Consumer Groups**: Scales partitions processing concurrently.
- **Manual Offset Commit**: Allows granular offset manipulation.

---

## 2. API Examples

```go
import "prahari/shared/kafka/consumer"

c, _ := consumer.NewConsumer(kafkaConfig, "worker-telemetry")
defer c.Close()

// Start consumer loop
err := c.StartConsumer(ctx, func(ctx context.Context, key, val []byte) error {
    // Process vitals
    return nil // Commits message offset automatically
})
```
