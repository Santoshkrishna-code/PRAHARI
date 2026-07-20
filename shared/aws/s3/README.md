# AWS S3 Storage Adapter

This package wraps AWS S3 SDK actions for file uploads, downloads, and secure pre-signing URL generations.

---

## 1. Features

- **Upload & Multipart Support**: Stream uploads from readers.
- **Pre-signed URLs**: Secure temporary access URL generation.
- **Mockable interfaces**: Unit test support.

---

## 2. API Examples

### A. Uploading files
```go
import "prahari/shared/aws/s3"

client := s3.NewClient(s3SDKClient)
err := client.Upload(ctx, "bucket-name", "path/to/key.png", fileReader)
```

### B. Pre-signing file downloads
```go
url, _ := client.PresignGetObject(ctx, "bucket-name", "path/to/key.png", 15*time.Minute)
```
