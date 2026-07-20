# PRAHARI Kafka & Event Streaming Platform SDK

This package provides a connection loader, message publishers, consumer groups, message envelopes, schema validators, retry backoffs, DLQ handles, and middleware interceptors.

---

## 1. Directory Structure

- **`producer/`**: Sync, async, and batch event publishers.
- **`consumer/`**: Polling loops, consumer groups, and manual offset commits.
- **`events/`**: Standard message Envelopes, serializers, and version compatibility helpers.
- **`schema/`**: Validation checks for payloads structure.
- **`retry/`**: Exponential backoff with random jitter.
- **`dlq/`**: Poisoned message producers and handlers.
- **`middleware/`**: Logging, OTel tracing, and atomic metrics trackers.

---

## 2. Usage Blueprint

Exposes connection checks (`Ping`) and wraps message handling.
```go
import (
    "prahari/shared/kafka"
    "prahari/shared/kafka/producer"
)

client := kafka.NewManager(kafka.Config{Brokers: []string{"localhost:9092"}})
prod, _ := producer.NewProducer(kafkaConfig)
```
