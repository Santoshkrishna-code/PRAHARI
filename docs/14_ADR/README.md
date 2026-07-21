# PRAHARI Platform: Architecture Decision Records (ADRs)

This document contains the Architecture Decision Records (ADRs) mapping out the major technology choices and design paradigms implemented across the PRAHARI Platform.

---

## Index of ADRs
- [ADR-001: Authentication Strategy (AWS Cognito)](#adr-001-authentication-strategy-aws-cognito)
- [ADR-002: Authorization Policy Engine (Casbin ABAC/RBAC)](#adr-002-authorization-policy-engine-casbin-abacrbac)
- [ADR-003: CQRS Pattern Separation](#adr-003-cqrs-pattern-separation)
- [ADR-004: Messaging Architecture (Kafka Broker)](#adr-004-messaging-architecture-kafka-broker)
- [ADR-005: Relational Transactional Database (PostgreSQL)](#adr-005-relational-transactional-database-postgresql)
- [ADR-006: In-Memory Caching & Rate Limiting (Redis)](#adr-006-in-memory-caching--rate-limiting-redis)
- [ADR-007: Container Orchestration (Kubernetes / EKS)](#adr-007-container-orchestration-kubernetes--eks)
- [ADR-008: AI Agent Architecture (LangGraph Python Agents)](#adr-008-ai-agent-architecture-langgraph-python-agents)
- [ADR-009: Computer Vision Model Inferences (YOLOv8 & TensorRT)](#adr-009-computer-vision-model-inferences-yolov8--tensorrt)
- [ADR-010: Digital Twin Synchronizations (Three.js & WebSockets)](#adr-010-digital-twin-synchronizations-threejs--websockets)
- [ADR-011: Event Sourcing & Eventual Consistency (Transactional Outbox)](#adr-011-event-sourcing--eventual-consistency-transactional-outbox)
- [ADR-012: Observability Architecture (OpenTelemetry, Prometheus & Jaeger)](#adr-012-observability-architecture-opentelemetry-prometheus--jaeger)
- [ADR-013: Multi-Tenancy Architecture (Database-per-Tenant Separation)](#adr-013-multi-tenancy-architecture-database-per-tenant-separation)
- [ADR-014: Storage Encryption Strategy (AWS KMS Envelope Encryption)](#adr-014-storage-encryption-strategy-aws-kms-envelope-encryption)
- [ADR-015: Deployment Methodology (ArgoCD GitOps Pipelines)](#adr-015-deployment-methodology-argocd-gitops-pipelines)

---

## ADR-001: Authentication Strategy (AWS Cognito)
- **Status**: Approved
- **Context**: The platform requires secure, scalable user credential storage, MFA enforcement, and identity federation (SAML/OIDC) for enterprise tenants.
- **Decision**: Use **AWS Cognito User Pools** as the primary Identity Provider (IdP), exposed via a facade Auth Service.
- **Alternatives**: Self-hosted Keycloak (rejected due to database backup and clustering operational overhead).
- **Consequences & Trade-offs**: Reduces operational cost and handles security audit burdens, but ties the primary login flow to AWS APIs.

---

## ADR-002: Authorization Policy Engine (Casbin ABAC/RBAC)
- **Status**: Approved
- **Context**: Enterprise tenants require customizable, site-specific permissions combined with attribute limits (e.g. "Only Lead Engineers can approve PHA studies under MOC-09").
- **Decision**: Integrate **Casbin** as an ABAC (Attribute-Based Access Control) policy engine inside our microservice middlewares.
- **Alternatives**: Custom SQL authorization queries (rejected due to code duplication and difficulty auditing access tables).
- **Consequences & Trade-offs**: Provides highly flexible permission management, but requires caching policies in Redis to avoid performance bottlenecks.

---

## ADR-003: CQRS Pattern Separation
- **Status**: Approved
- **Context**: Read requests (e.g. viewing active safety charts) outnumber writes (e.g. registering chemicals) by a 50:1 ratio. Real-time dashboards must load in < 150ms.
- **Decision**: Separate read and write paths (Command Query Responsibility Segregation). Writes use PostgreSQL, reads use dedicated materialized views or Redis cache pools.
- **Alternatives**: Unified CRUD layers (rejected due to DB locking during intensive analytics queries).
- **Consequences & Trade-offs**: Optimizes read performance, but introduces eventual consistency latency.

---

## ADR-004: Messaging Architecture (Kafka Broker)
- **Status**: Approved
- **Context**: Microservices need to broadcast plant state changes (e.g. chemical intake, CV alert triggers) to downstream reporting, metrics, and notification services asynchronously.
- **Decision**: Deploy **Apache Kafka** (via AWS MSK) as the primary enterprise event broker.
- **Alternatives**: RabbitMQ (rejected due to lacking event replay capabilities and lower partitioning throughput).
- **Consequences & Trade-offs**: Highly scalable event streaming with message replay capabilities, but requires managing partition strategies and client retry loops.

---

## ADR-005: Relational Transactional Database (PostgreSQL)
- **Status**: Approved
- **Context**: The core database must support complex queries, transactional guarantees (ACID), and handle spatial structures (pgvector for AI embeddings).
- **Decision**: Standardize on **PostgreSQL (AWS Aurora)**.
- **Alternatives**: MongoDB (rejected due to lacking robust multi-row foreign key constraints and transactional integrity).
- **Consequences & Trade-offs**: Excellent SQL indexing, pgvector support, and mature ecosystem, but scaling writes requires partitioning.

---

## ADR-006: In-Memory Caching & Rate Limiting (Redis)
- **Status**: Approved
- **Context**: Prevent database bloat on frequently accessed assets (e.g. compatibility lookup vectors) and enforce API rate limits at the gateway layer.
- **Decision**: Deploy **Amazon ElastiCache Redis** for distributed caching and sliding-window rate limit counters.
- **Alternatives**: In-memory Go maps (rejected due to lacking synchronization across multiple microservice pods).
- **Consequences & Trade-offs**: Sub-millisecond response times, but introduces cache invalidation complexity.

---

## ADR-007: Container Orchestration (Kubernetes / EKS)
- **Status**: Approved
- **Context**: The platform consists of dozens of stateless microservices and stateful collectors that must run reliably, autoscale, and deploy via GitOps.
- **Decision**: Use **Amazon EKS (Elastic Kubernetes Service)** as the container runtime.
- **Alternatives**: AWS ECS (rejected due to limited scaling options and lack of native support for GitOps tooling like ArgoCD).
- **Consequences & Trade-offs**: Industry-standard container scaling, but requires managing cluster network security groups and ingress configurations.

---

## ADR-008: AI Agent Architecture (LangGraph Python Agents)
- **Status**: Approved
- **Context**: Automating SDS audits and Process Safety recommendations requires multi-turn execution paths, custom memory, and RAG lookups.
- **Decision**: Implement AI agents using **LangGraph** within Python containers.
- **Alternatives**: Raw LLM API calls with basic template formatting (rejected due to high rate of hallucinations and lack of logical loops).
- **Consequences & Trade-offs**: Clean state-machine orchestration, but adds execution latency (1-3 seconds per run).

---

## ADR-009: Computer Vision Model Inferences (YOLOv8 & TensorRT)
- **Status**: Approved
- **Context**: Edge gateways must process multiple high-resolution RTSP camera feeds and detect PPE violations in under 100ms.
- **Decision**: Use **YOLOv8** compiled to **TensorRT engine format** executing on NVIDIA GPU hardware.
- **Alternatives**: Cloud-only Amazon Rekognition (rejected due to high latency, bandwidth consumption, and operating costs).
- **Consequences & Trade-offs**: Sub-10ms model execution times at the edge, but requires specialized local GPU-accelerated gateways.

---

## ADR-010: Digital Twin Synchronizations (Three.js & WebSockets)
- **Status**: Approved
- **Context**: Operators require real-time 3D visualization of chemical container positions, alarms, and sensor states without reloading.
- **Decision**: Build the frontend twin in **Three.js (WebGL)** and synchronize state via **WebSockets** connected to the gateway.
- **Alternatives**: Static 2D floor plans (rejected as premium enterprise visual twins are required).
- **Consequences & Trade-offs**: Interactive 3D graphics, but requires clients to possess capable browser hardware.

---

## ADR-011: Event Sourcing & Eventual Consistency (Transactional Outbox)
- **Status**: Approved
- **Context**: Prevent data inconsistencies where a database write succeeds but publishing the corresponding Kafka event fails.
- **Decision**: Implement the **Transactional Outbox Pattern** across all Go microservices.
- **Alternatives**: Direct inline publishing to Kafka inside the DB transaction block (rejected as it causes lockups on connection failures).
- **Consequences & Trade-offs**: Guarantees at-least-once event delivery to Kafka, but adds slight event delivery latency (50-200ms).

---

## ADR-012: Observability Architecture (OpenTelemetry, Prometheus & Jaeger)
- **Status**: Approved
- **Context**: Need a vendor-neutral observability architecture to trace requests across distributed network networks, database operations, and Kafka topics.
- **Decision**: Instrument code with **OpenTelemetry APIs**, routing metrics to Prometheus and traces to Jaeger.
- **Alternatives**: Vendor-specific SDKs like Datadog (rejected due to vendor lock-in).
- **Consequences & Trade-offs**: Full trace observability across runtime layers, but adds trace context header propagation overhead.

---

## ADR-013: Multi-Tenancy Architecture (Database-per-Tenant Separation)
- **Status**: Approved
- **Context**: Enterprise chemical customers require absolute data isolation to satisfy internal security regulations and compliance audits.
- **Decision**: Implement **Database-per-Tenant isolation** at the RDS layer.
- **Alternatives**: Shared database with tenant_id filtering columns (rejected due to data leak risks).
- **Consequences & Trade-offs**: High isolation and ease of compliance, but increases database management overhead.

---

## ADR-014: Storage Encryption Strategy (AWS KMS Envelope Encryption)
- **Status**: Approved
- **Context**: Sensitive Safety Data Sheets and audit trails must be encrypted at rest with customer-managed keys.
- **Decision**: Implement **Envelope Encryption** utilizing AWS KMS Customer Managed Keys.
- **Alternatives**: Default AWS-managed S3 bucket encryption (rejected due to lack of granular key rotation control).
- **Consequences & Trade-offs**: Compliance-aligned security footprint, but increases AWS KMS API call costs.

---

## ADR-015: Deployment Methodology (ArgoCD GitOps Pipelines)
- **Status**: Approved
- **Context**: Ensure that the deployed state in Kubernetes always matches the declarative configs in git.
- **Decision**: Standardize on **ArgoCD** as our continuous deployment tool.
- **Alternatives**: Push-based deployments via GitHub Actions runner commands (rejected due to lack of automatic drift detection).
- **Consequences & Trade-offs**: Automated drift detection and zero-touch cluster synchronization, but requires maintaining a separate GitOps config repository.
