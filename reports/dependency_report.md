# Dependency Verification & System Environment Audit Report

**Audit Date**: 2026-07-22  
**Auditor**: Independent Enterprise Software Auditor  

---

## 1. System & Language Runtimes

| Component | Target Version | Observed Version | Verification Command | Status |
|---|---|---|---|:---:|
| **Python** | 3.11+ | Python 3.13.9 | `python3 --version` | **VERIFIED** |
| **Go** | 1.21+ | go1.26.5 darwin/arm64 | `go version` | **VERIFIED** |
| **Node.js** | 20+ | v26.0.0 | `node -v` | **VERIFIED** |
| **NPM** | 9+ | 11.12.1 | `npm -v` | **VERIFIED** |
| **Git LFS** | Installed | git-lfs/3.6.0 | `git lfs version` | **VERIFIED** |

---

## 2. Python Virtual Environment Dependencies ([ai/platform/.venv](file:///Users/Santosh/PRAHARI/ai/platform/.venv))

| Package | Version Installed | Purpose | Status |
|---|---|---|:---:|
| **FastAPI** | `0.115.0+` | Web framework & OpenAPI documentation | **VERIFIED** |
| **Pydantic** | `2.9.0+` | Data schema validation & serialization | **VERIFIED** |
| **SQLAlchemy** | `2.0.35+` | Async ORM & database session management | **VERIFIED** |
| **Pytest** | `9.1.1` | Automated integration and unit testing framework | **VERIFIED** |
| **aiosqlite** | `0.20.0` | Async SQLite driver for isolated testing | **VERIFIED** |
| **asyncpg** | `0.29.0` | High-performance PostgreSQL async database driver | **VERIFIED** |

---

## 3. Infrastructure Services (`docker-compose.yml`)

- **PostgreSQL 15-alpine**: Defined on port `5432` with database `prahari_db`.
- **TimescaleDB pg15**: Defined on port `5433` with database `prahari_telemetry`.
- **Redis 7-alpine**: Defined on port `6379`.
- **Confluent Kafka 7.3.0 (KRaft mode)**: Defined on port `9092`.
- **OpenSearch 2.8.0**: Defined on port `9200`.
- **Neo4j 5-community**: Defined on port `7474` / `7687`.
