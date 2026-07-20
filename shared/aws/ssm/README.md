# AWS SSM Parameter Store Adapter

This package wraps AWS SSM Parameter Store SDK client actions for key-value configuration values retrieval.

---

## 1. Features

- **Parameter string retriever**: Resolves configurations dynamically.
- **Secure parameters decryption**: Handles KMS decryptions under the hood.

---

## 2. API Examples

```go
import "prahari/shared/aws/ssm"

client := ssm.NewClient(ssmSDKClient)
port, _ := client.GetParameter(ctx, "/prahari/auth-service/http-port")
```
