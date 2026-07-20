# Secrets and Key Management Subsystem

This package implements KMS client adapters for cryptographic key unsealing and maps webhook structures for automatic secret rotation.

---

## 1. Features

- **KMSClient**: Integrates with AWS KMS to encrypt/decrypt envelope keys.
- **RotationEvent**: Standard JSON models mapping Secrets Manager lambda webhooks.

---

## 2. API Examples

### A. Initializing the KMS Client
```go
import (
    "prahari/shared/security/secrets"
    awsKMS "github.com/aws/aws-sdk-go-v2/service/kms"
)

// Construct from AWS Configuration Client
sdkClient := awsKMS.NewFromConfig(awsConfig)
kmsAdapter := secrets.NewKMSClient(sdkClient)
```
