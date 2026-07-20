# Platform Mock Client Adapters

This package registers custom struct-based mocks allowing developers to hook and stub transactions.

---

## 1. Features

- **Decoupled stubs**: Overrides S3, SQS, Redis, and Kafka transactions using hook functions.

---

## 2. API Examples

```go
import "prahari/shared/testing/mocks"

s3Mock := &mocks.MockS3Client{
    UploadFunc: func(ctx context.Context, bucket, key string, data []byte) error {
        return errors.New("s3 write disk failure")
    },
}
```
