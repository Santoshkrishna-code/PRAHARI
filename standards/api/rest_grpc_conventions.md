# PRAHARI Platform API Conventions
## REST & gRPC Specifications

This document defines the strict API design standards to ensure consistency, discoverability, and compatibility across all PRAHARI services.

---

## 1. REST API Design Guidelines

### 1.1 URI Path Schema
1. URIs MUST be lowercase.
2. Resource names MUST be plural (e.g., `/api/v1/permits`, NOT `/api/v1/permit`).
3. Paths MUST indicate hierarchy: `/api/v1/plants/{plant_id}/assets/{asset_id}`.
4. URIs MUST be versioned using a major version prefix (e.g., `/api/v1/`).

### 1.2 HTTP Method Mapping
- **`GET`**: Retrieve a resource or list. Safe and idempotent.
- **`POST`**: Create a new resource or execute a non-CRUD action. Non-idempotent.
- **`PUT`**: Replace an existing resource completely. Idempotent.
- **`PATCH`**: Partially update an existing resource. Non-idempotent.
- **`DELETE`**: Remove a resource. Idempotent.

### 1.3 Request & Response Body Casing
- All JSON fields in request and response payloads MUST use **`snake_case`**.

```json
{
  "permit_id": "ptw-99201",
  "worker_id": "usr-8819",
  "risk_assessment_required": true
}
```

### 1.4 Standard Error Response Envelope
Every error response MUST return the standard error schema with a relevant HTTP status code:

```json
{
  "error": {
    "code": "INVALID_ARGUMENT",
    "message": "The permit request parameters are invalid.",
    "details": [
      {
        "field": "expiration_time",
        "issue": "Permit expiration cannot be in the past."
      }
    ]
  }
}
```

### 1.5 Pagination
For list queries, pagination parameters MUST use offset/limit or cursor tokens:
- Query parameters: `page_size` (default: 20, max: 100) and `page_token` (base64 cursor).
- Response structure:
  ```json
  {
    "items": [],
    "next_page_token": "eyJPZmZzZXQiOjIwfQ=="
  }
  ```

---

## 2. gRPC API Design Guidelines

gRPC is used for high-performance, internal microservice-to-microservice calls.

### 2.1 Proto3 Style Guidelines
1. Files MUST be located under the service's `proto/` directory or shared under `shared/protobuf/`.
2. Packages MUST match the domain hierarchy: `package prahari.permit.v1;`.
3. Message names MUST use `PascalCase`. Field names MUST use `snake_case`.

```protobuf
syntax = "proto3";

package prahari.permit.v1;

option go_package = "prahari/shared/protobuf/permitv1";

message CreatePermitRequest {
  string worker_id = 1;
  string zone_id = 2;
  int64 duration_seconds = 3;
}
```

### 2.2 Standard RPC Method Names
- `Get[Resource]`: Retrieve a single item.
- `List[Resource]`: Query list of items.
- `Create[Resource]`: Add new item.
- `Update[Resource]`: Modify existing item.
- `Delete[Resource]`: Remove item.

### 2.3 gRPC Status Codes Mapping
Inter-service gRPC errors MUST map to standard codes:
- **`OK` (0)**: Success.
- **`INVALID_ARGUMENT` (3)**: Client specified an invalid argument.
- **`UNAUTHENTICATED` (16)**: Request lacks valid credentials.
- **`PERMISSION_DENIED` (7)**: Client lacks permissions (RBAC/ABAC).
- **`NOT_FOUND` (5)**: Resource not found.
- **`ALREADY_EXISTS` (6)**: Resource already exists.
- **`INTERNAL` (13)**: Internal system error.
