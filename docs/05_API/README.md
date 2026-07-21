# PRAHARI Platform: API Documentation Specification

This document defines the interface standards, security schemas, request/response models, pagination protocols, error formats, and rate limits for the PRAHARI Platform.

---

## 1. Authentication & Security Scheme
All API endpoints (except public health checks `/healthz`) require JWT authentication. 

### Header Format
```http
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1...
X-Tenant-ID: tenant-12003
```

---

## 2. API conventions (Filtering, Pagination & Sorting)
To maintain performance on large datasets, all collections (`GET` requests) must implement pagination, filtering, and sorting conventions.

### 2.1 Pagination Parameters
- `page`: Page index (default: `1`, 1-indexed).
- `limit`: Number of items per page (default: `20`, maximum: `100`).

### 2.2 Filtering Query Operators
Filters are specified via query parameters using standard field names and value matches:
- `?status=ACTIVE` (Equality)
- `?molecular_weight=gt:100.5` (Greater than)
- `?created_at=between:2026-01-01T00:00:00Z,2026-01-31T23:59:59Z` (Range filter)

### 2.3 Sorting Query Parameters
- `sort`: Field name prefixed with `+` for ascending or `-` for descending. E.g., `sort=-created_at,name` sorting first by creation date descending, then by name ascending.

---

## 3. RFC 7807 Error Details
All API errors are returned in standard **RFC 7807: Problem Details for HTTP APIs** format.

### Example Error Response
```json
{
  "type": "https://prahari.io/errors/compatibility-conflict",
  "title": "Chemical Compatibility Conflict",
  "status": 409,
  "detail": "Chemical 'Hydrofluoric Acid' cannot be placed in Storage Area 'Zone C-3' due to proximity with incompatible basic solutions.",
  "instance": "/api/v1/inventory/containers",
  "invalid_params": [
    {
      "name": "storage_area_id",
      "reason": "Zone C-3 contains 3 containers of Sodium Hydroxide (NaOH) which violates rule CHEM-COMP-09."
    }
  ]
}
```

---

## 4. API Endpoints Reference

### 4.1 Chemical Registration (`POST /api/v1/chemicals`)

**Request Payload**
```json
{
  "plant_id": "pl-00912",
  "name": "Sodium Cyanide",
  "cas_number": "143-33-9",
  "formula": "NaCN",
  "molecular_weight": 49.007,
  "physical_state": "SOLID",
  "max_allowable_qty": 200.00
}
```

**Response Payload (201 Created)**
```json
{
  "id": "chem-5501-ab8",
  "plant_id": "pl-00912",
  "name": "Sodium Cyanide",
  "cas_number": "143-33-9",
  "formula": "NaCN",
  "molecular_weight": 49.007,
  "physical_state": "SOLID",
  "is_restricted": true,
  "max_allowable_qty": 200.00,
  "status": "APPROVED_INTAKE",
  "created_at": "2026-07-21T13:20:00Z",
  "updated_at": "2026-07-21T13:20:00Z"
}
```

### 4.2 Start PHA Study (`POST /api/v1/studies`)

**Request Payload**
```json
{
  "plant_id": "pl-00912",
  "unit_id": "ut-distillation",
  "title": "Tower-101 HAZOP and LOPA Re-evaluation",
  "method": "HAZOP",
  "leader_id": "usr-lead-991",
  "scribe_id": "usr-scrb-220",
  "target_date": "2026-12-31T00:00:00Z"
}
```

**Response Payload (201 Created)**
```json
{
  "id": "study-8891-cca",
  "study_number": "PHA-2026-0081",
  "plant_id": "pl-00912",
  "unit_id": "ut-distillation",
  "title": "Tower-101 HAZOP and LOPA Re-evaluation",
  "method": "HAZOP",
  "status": "DRAFT",
  "leader_id": "usr-lead-991",
  "scribe_id": "usr-scrb-220",
  "target_date": "2026-12-31T00:00:00Z",
  "revalidation_due_at": "2031-12-31T00:00:00Z",
  "created_at": "2026-07-21T13:20:10Z"
}
```

---

## 5. Rate Limits & SLA
Rate limits are applied at the API Gateway using sliding window counters stored in Redis:
- **Standard User Accounts**: 1,200 requests / minute per Tenant.
- **Edge Devices (CV stream alerts)**: 10,000 requests / minute per IP.
- **System-to-System Integrations**: 3,000 requests / minute per Service Key.

Rate limit status is communicated in HTTP response headers:
```http
X-RateLimit-Limit: 1200
X-RateLimit-Remaining: 1198
X-RateLimit-Reset: 58
```
When rate limits are exceeded, the gateway returns a `429 Too Many Requests` status code with an RFC 7807 error detail.
