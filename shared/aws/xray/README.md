# AWS X-Ray Tracing Adapter

This package wraps AWS X-Ray SDK actions for distributed tracing subsegments lifecycle.

---

## 1. Features

- **Segment Recorder**: Records span checkpoints.
- **Context-aware**: Integrates with standard net/http contexts.

---

## 2. API Examples

```go
import "prahari/shared/aws/xray"

client := xray.NewClient(true)
ctx, segment := client.StartSegment(ctx, "DynamoDB:SaveVitals")
defer segment.Close()
```
