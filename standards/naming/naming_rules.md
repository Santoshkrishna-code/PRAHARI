# PRAHARI Naming Conventions
## Enterprise Naming Rules

This document establishes the naming formats for Kafka topics, AWS resources, databases, Docker image tags, and environmental variables across the PRAHARI ecosystem.

---

## 1. Kafka Event Topics

Kafka topic names must follow a standard namespace hierarchy to prevent collisions and simplify ACL routing policies.

**Pattern**: `prahari.<environment>.<domain>.<event-type>.<version>`

* **`environment`**: `dev`, `stage`, `prod`
* **`domain`**: Area originating the event (e.g., `telemetry`, `cv`, `permits`, `incidents`, `notifications`).
* **`event-type`**: Noun describing the outcome (e.g., `heartrate-stream`, `ppe-violation`, `permit-approved`, `incident-created`).
* **`version`**: Major version number of the schema (e.g., `v1`, `v2`).

### Examples
- `prahari.dev.telemetry.vitals-stream.v1`
- `prahari.prod.cv.ppe-violation.v1`
- `prahari.stage.permits.status-changed.v1`

---

## 2. AWS Resource Naming

AWS resources must follow a structured naming pattern to enable automated resource tagging and billing audits.

**Pattern**: `prahari-<environment>-<region-short>-<resource-type>-<name>`

* **`environment`**: `dev`, `stage`, `prod`
* **`region-short`**: `ue1` (us-east-1), `uw2` (us-west-2), etc.
* **`resource-type`**: AWS resource abbreviation (e.g., `vpc`, `eks`, `rds`, `s3`, `iam`, `msk`).
* **`name`**: Descriptive suffix.

### Examples
- `prahari-dev-ue1-vpc-main` (VPC)
- `prahari-prod-ue1-eks-core` (EKS Cluster)
- `prahari-stage-ue1-rds-permitdb` (Postgres RDS Instance)
- `prahari-prod-ue1-s3-incidents-bucket` (S3 Bucket)

---

## 3. Database Schemas

- **Table Names**: lowercase, `snake_case`, plural (e.g., `permits`, `work_zones`, `incident_reports`).
- **Column Names**: lowercase, `snake_case` (e.g., `created_at`, `approved_by`).
- **Primary Keys**: UUID columns named `id` (or `{table_singular}_id`).
- **Foreign Keys**: `{referenced_table_singular}_id` (e.g., `worker_id` referencing `workers.id`).

---

## 4. Docker Image Tags

**Pattern**: `<registry>/prahari-<service-name>:<environment>-<commit-sha>`

### Examples
- `123456789012.dkr.ecr.us-east-1.amazonaws.com/prahari-auth-service:dev-d92ba7f`
- `123456789012.dkr.ecr.us-east-1.amazonaws.com/prahari-permit-service:staging-f2910ba`
- `123456789012.dkr.ecr.us-east-1.amazonaws.com/prahari-incident-service:v1.0.0` (Production release tag)
