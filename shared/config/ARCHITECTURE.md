# Architecture: Context Configurations Loader

This document outlines the core architecture and pipeline behind the `config` shared library.

---

## 1. Design Overview

To support multi-tenant scaling and Kubernetes-native deployment:
- **No Global Mutables**: Configuration is loaded into immutable structs and passed down via Dependency Injection.
- **Fail-Fast Boundaries**: Syntactic validation errors (e.g. incorrect ports) trigger startup crashes immediately.
- **Mockability**: External configuration endpoints (SSM, Secrets Manager) are abstracted behind mockable interfaces in `sources/`.

---

## 2. Override Hierarchy Pipeline

The config loading priority rules:
1. `SetDefaults` sets core port defaults.
2. Local `.env` file values are injected.
3. System environment parameters (Kubernetes environment variables) map to struct properties.
4. AWS SSM Parameter Store queries replace system fields.
5. AWS Secrets Manager pulls values (DB Passwords) to overwrite fields.
