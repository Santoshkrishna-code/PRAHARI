# Go Microservice Template

This template establishes the architectural design pattern and folder layout for all Go-based microservices in the **PRAHARI** platform.

## Folder Directory Mapping

* **`cmd/`**: Entrypoint initializing resources and starting the HTTP listener.
* **`internal/domain/`**: Enterprise Domain Objects and ports interfaces (DDD).
* **`internal/service/`**: Domain core business rule orchestrator.
* **`internal/api/`**: Ingress API controller adapter.
* **`internal/repository/`**: Egress database adapter.
* **`internal/middleware/`**: JWT validation and request tracking wrappers.
* **`internal/config/`**: Local environment configs loader.

## How to Initialize a New Service

1. Copy the contents of this folder into `services/<new-service-name>`.
2. Rename the module in `go.mod` to `prahari/services/<new-service-name>`.
3. Add the service entry in the root-level `go.work` file.
4. Replace references to `Permit` with your service's core domain entity (e.g. `Incident`, `Worker`, `Asset`).
5. Run tests:
   ```bash
   make test
   ```
