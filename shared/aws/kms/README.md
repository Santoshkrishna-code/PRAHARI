# AWS KMS Adapter

This package wraps AWS KMS SDK client actions for custom data keys generation and master key integrations.

---

## 1. Features

- **Data Key Generator**: Generates 256-bit symmetric data keys.
- **Master Key Integration**: Integrates with IAM controls.

---

## 2. API Examples

```go
import "prahari/shared/aws/kms"

client := kms.NewClient(kmsSDKClient)
plaintextKey, encryptedKey, _ := client.GenerateDataKey(ctx, "alias/prahari-master-key")
```
