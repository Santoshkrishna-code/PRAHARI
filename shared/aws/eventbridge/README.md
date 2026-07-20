# AWS EventBridge Integration Adapter

This package wraps AWS EventBridge SDK client actions for publishing custom events to inter-service event buses.

---

## 1. Features

- **Custom Bus Event Publisher**: Puts events to target event buses.
- **Failures validation**: Checks for failed entry counts returned by EventBridge.

---

## 2. API Examples

```go
import "prahari/shared/aws/eventbridge"

client := eventbridge.NewClient(ebSDKClient)
eventID, _ := client.PublishEvent(ctx, "prahari-events-bus", "auth-service", "UserRegistered", `{"user_id":"123"}`)
```
