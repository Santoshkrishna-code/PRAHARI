# Architecture: AWS SDK Wrapper Subsystem

This document outlines the architectural patterns and design specifications governing the `shared/aws` package.

---

## 1. Key Design Tenets

- **Strict Context Propagation**: Every SDK call accepts `context.Context` parameters to allow parent cancellation signals and deadline timeouts to bubble down.
- **Mocking Strategy**: Clients wrap the SDK structs and expose narrow interfaces (like `S3API` and `SSMAPI`) to allow mocking in unit tests.
- **Unified Health checks**: All adapters implement a standardized `aws.HealthChecker` interface exposing a `Ping(ctx)` method, supporting runtime diagnostic assertions.

---

## 2. Integration Mappings

```
        +----------------------------------------+
        |   Microservice Dependency Container    |
        +-------------------+--------------------+
                            |
                            v
        +-------------------+--------------------+
        |      Master AWS SDK Manager Facade     |
        +-------------------+--------------------+
                            |
     +-----------+----------+----------+-----------+
     |           |                     |           |
     v           v                     v           v
  [ S3 ]      [ SQS ]             [ DynamoDB ]   [ KMS ]
 (s3/client) (sqs/producer)       (ddb/client)  (kms/client)
     |           |                     |           |
     +-----------+----------+----------+-----------+
                            |
                            v
                     AWS SDK v2 Client
```
