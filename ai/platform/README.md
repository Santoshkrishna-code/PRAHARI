# PRAHARI AI Platform Foundation (Volume 0)

This is the foundation service layer (Volume 0) of the PRAHARI AI Platform. It provides a reusable, highly performant python framework incorporating clean architecture, DDD patterns, asynchronous database/cache providers, structural logging, request tracking, and full OpenTelemetry/Prometheus observability.

---

## 1. Technical Stack
- **Base Environment**: Python 3.13
- **Web Layer**: FastAPI & Uvicorn
- **Dependency Management**: Poetry
- **Database Adapters**: SQLAlchemy 2.0 (asyncpg)
- **Cache Registry**: Redis
- **Observability**: OpenTelemetry & Prometheus metrics
- **Code Standards**: Ruff & MyPy

---

## 2. Directory Layout
```
ai/platform/
├── app/
│   ├── main.py                 # FastAPI lifecycle & application setup
│   ├── core/                   # Configuration, common DTOs, retry loops, exceptions
│   ├── infrastructure/         # DB, Redis, Logging, Telemetry, Registry providers
│   └── interfaces/             # HTTP router controllers & Middlewares
```

---

## 3. How to Run Locally

### 3.1 Prerequisites
Ensure you have Docker and Docker Compose installed.

### 3.2 Running the Application
Spin up the service and its dependencies:
```bash
docker compose up --build
```

The service will start and expose:
- **API Documentation**: [http://localhost:8000/docs](http://localhost:8000/docs)
- **JSON Schema**: [http://localhost:8000/openapi.json](http://localhost:8000/openapi.json)
- **System Health**: [http://localhost:8000/health](http://localhost:8000/health)
- **Prometheus Telemetry**: [http://localhost:8000/metrics](http://localhost:8000/metrics)
