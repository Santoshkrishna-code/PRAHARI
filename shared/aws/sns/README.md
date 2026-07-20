# AWS SNS Notification Adapter

This package wraps AWS SNS SDK client actions for topic publications and direct SMS dispatches.

---

## 1. Features

- **Topic Publisher**: Dispatches notifications.
- **SMS dispatch**: Sends direct mobile notifications.

---

## 2. API Examples

### A. Publishing to topics
```go
import "prahari/shared/aws/sns"

client := sns.NewClient(snsSDKClient)
msgID, _ := client.PublishMessage(ctx, "arn:aws:sns:...", "Vitals stream Alert")
```

### B. Sending SMS alerts
```go
client.PublishSMS(ctx, "+15555555555", "Emergency Alert: fire detected")
```
