# PRAHARI: Auth Service Microservice

The **Auth Service** is the platform's central gatekeeper, managing user authentication, registration, session refreshes, and access token validation. It uses **Amazon Cognito User Pools** as the backing identity directory.

---

## 1. Domain Architecture

Following clean hexagonal architecture principles:
- **`cmd/server/main.go`**: Entrypoint initiating the application container.
- **`bootstrap/`**: DI container builder loading configurations, structured logging, and AWS sessions.
- **`internal/domain/`**: Bounded context entities, JWT claims, and service interfaces.
- **`internal/dto/`**: JSON request/response binders.
- **`internal/repository/`**: Cognito SDK caller and InMemory read-through caching.
- **`internal/service/`**: Core registration, login, and token refresh business logic.
- **`internal/api/v1/`**: Chi Router endpoints and handlers (Login, Register, Refresh, Profile).
- **`internal/middleware/`**: JWT validation, unstructured logs, panic recovery, CORS, rate limiting, and security headers.
- **`internal/observability/`**: Metrics collection and OpenTelemetry/X-Ray tracing.

---

## 2. API Contract Endpoints

### REST HTTP API (Port 8001)

| Endpoint | Method | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `/api/v1/auth/register` | `POST` | Registers a new worker/supervisor. | No |
| `/api/v1/auth/login` | `POST` | Validates credentials and returns JWTs. | No |
| `/api/v1/auth/refresh` | `POST` | Renews expired access tokens. | No |
| `/api/v1/auth/profile` | `GET` | Fetches parsed user profile claims. | Yes (Bearer) |
| `/health/live` | `GET` | Process liveness check. | No |
| `/health/ready` | `GET` | Downstream AWS Cognito readiness check. | No |

---

## 3. Emitted Event Schemas

The service publishes audit streams to MSK (Kafka) on user status changes:
- `prahari.<env>.auth.user-registered.v1`
- `prahari.<env>.auth.user-logged-in.v1`
- `prahari.<env>.auth.token-refreshed.v1`

---

## 4. Local Development

### 1. Build and Run Local Binary
```bash
make build
./bin/auth-service
```

### 2. Execute Tests
```bash
make test
```

### 3. Build Production Container
```bash
make docker-build
```
