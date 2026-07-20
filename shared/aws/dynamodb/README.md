# AWS DynamoDB Database Adapter

This package wraps AWS DynamoDB SDK client actions for document CRUD and transactional operations.

---

## 1. Features

- **Document CRUD**: Marshals and unmarshals structures automatically.
- **Transactions**: Supports ACID transactions across multiple items.

---

## 2. API Examples

### A. Inserting document items
```go
import "prahari/shared/aws/dynamodb"

client := dynamodb.NewClient(ddbSDKClient)

type WorkerVitals struct {
    WorkerID  string  `dynamodbav:"worker_id"`
    HeartRate float64 `dynamodbav:"heart_rate"`
}

vitals := WorkerVitals{WorkerID: "worker-123", HeartRate: 78.5}
err := client.PutItem(ctx, "WorkerVitalsTable", vitals)
```
