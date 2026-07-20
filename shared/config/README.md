# Context Configurations Loader

This package parses configuration parameters across all PRAHARI microservices.

---

## 1. Features

- **Consolidated Config struct**: Aggregates properties for AWS, DB, Redis, Kafka, HTTP, gRPC, Telemetry, and Security settings.
- **Fail-Fast Structural Validation**: Checks all fields upon load.
- **Multiple Config Sources**: Supports Dotenv, standard environment parameters, AWS Systems Manager, and AWS Secrets Manager.
- **Thread-safe Watcher**: Enables runtime modifications tracking.

---

## 2. Config Loading Precedence

The loader resolves parameters sequentially:
1. **Defaults**: Standard bootable fallbacks.
2. **Dotenv (`.env`)**: Local developer overrides.
3. **Environment**: Container and Kubernetes parameters.
4. **AWS SSM**: Dynamic parameters.
5. **AWS Secrets Manager**: Passwords and client secrets.
