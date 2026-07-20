# AWS SES Email Adapter

This package wraps AWS SES SDK client actions for dispatching HTML emails.

---

## 1. Features

- **HTML Email Sender**: Sends notifications.
- **Quota Health check**: Reports connectivity metrics.

---

## 2. API Examples

```go
import "prahari/shared/aws/ses"

client := ses.NewClient(sesSDKClient)
msgID, _ := client.SendEmail(ctx, "alerts@prahari.com", "supervisor@prahari.com", "Safety Alert", "<h1>Fire Detected</h1>")
```
