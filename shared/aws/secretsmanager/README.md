# AWS Secrets Manager Adapter

This package wraps AWS Secrets Manager SDK client actions for credentials lookup and structure unmarshaling.

---

## 1. Features

- **String Secret Getter**: Pulls raw credential strings.
- **JSON Secret Unmarshaler**: Decodes configurations into target structs.

---

## 2. API Examples

```go
import "prahari/shared/aws/secretsmanager"

client := secretsmanager.NewClient(smSDKClient)

type DBCreds struct {
    Host     string `json:"host"`
    Password string `json:"password"`
}

var creds DBCreds
_ = client.GetSecretJSON(ctx, "prod/db/credentials", &creds)
```
