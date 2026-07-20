# AWS SQS Messaging Adapter

This package wraps SQS message publishers, long-polling consumer handlers, and dead-letter queue (DLQ) forwarders.

---

## 1. Features

- **Queue Publisher**: Sends JSON or raw string messages with delay offsets.
- **Polling Consumer**: Long-polling execution loop deletion hooks.
- **DLQ Redirection**: Moves unprocessable payloads to audit dead-letter queues.

---

## 2. API Examples

### A. Polling and processing SQS messages
```go
import "prahari/shared/aws/sqs"

client := sqs.NewClient(sqsSDKClient)

// Launches background thread
go func() {
    _ = client.StartConsumer(ctx, "https://sqs...", func(ctx context.Context, body string) error {
        // Execute message handler rules
        return nil // Deletes message automatically
    })
}()
```
