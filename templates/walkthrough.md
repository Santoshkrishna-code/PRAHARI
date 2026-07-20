# Master Walkthrough: Generator & Templates Framework

This document summarizes the complete implementation of the **PRAHARI Platform Generator & Service Templates** under `templates/`.

---

## 1. Directory Structure

- **`generator/`**: CLI tool parsing prompt inputs and validating clean architecture imports.
- **`registry/`**: Catalog of all generated microservices.
- **`microservice/`**: Production-ready DDD service template containing OTel metrics, log wrappers, database migration hooks, and Docker configurations.
- **`worker/`**: Background partition consumer loops.
- **`lambda/`**: AWS serverless invoke event handlers.
- **`cronjob/`**: Scheduled tasks executors.

---

## 2. CLI Generator Execution

The `prahari` CLI constructs code structures and registers generated systems automatically:

```bash
# Compile and build binary
go build -o bin/prahari ./templates/generator/cmd/prahari/...

# Scaffold a new service named 'incident'
./bin/prahari new service incident
```
