# AWS CloudWatch Custom Metrics Adapter

This package wraps AWS CloudWatch SDK client actions for custom metric records publications.

---

## 1. Features

- **Custom Metric exporter**: Sends values with standard units parameters.
- **Alarm Quotas Checker**: Supports health-checks connectivity diagnostics.

---

## 2. API Examples

```go
import "prahari/shared/aws/cloudwatch"

client := cloudwatch.NewClient(cwSDKClient)
_ = client.PutMetricData(ctx, "Prahari/VitalsAnalytics", "ActiveAlertsCount", 2.0, "Count")
```
