# PRAHARI AWS SDK Wrapper Platform SDK

This package wraps the AWS SDK v2 connection builders, providing mockable, context-aware service wrappers for S3, SNS, SQS, SES, KMS, Secrets Manager, Parameter Store, EventBridge, DynamoDB, CloudWatch, and X-Ray.

---

## 1. Directory Structure

- **`s3/`**: Uploads, downloads, presigned URLs.
- **`sns/`**: SMS and topic notifications publish.
- **`sqs/`**: Message producers, long polling consumers, and DLQ forwarders.
- **`ses/`**: HTML/raw email dispatches.
- **`kms/`**: Data keys generation and master key wrapping.
- **`secretsmanager/`**: Secret credential values and JSON loaders.
- **`ssm/`**: Parameter store string retrievals.
- **`eventbridge/`**: Event publication to custom buses.
- **`dynamodb/`**: CRUD operations and transactional queries.
- **`cloudwatch/`**: Custom telemetry metrics records.
- **`xray/`**: Subsegments lifecycle tracers.

---

## 2. Usage Blueprint

Each client wrapper is decoupled via interfaces and supports mock setups during testing. Exposes `Ping(ctx)` health checks to monitor connectivity states dynamically.
