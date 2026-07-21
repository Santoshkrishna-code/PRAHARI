# PRAHARI Platform: High Level Design (HLD)

## 1. Executive Summary
This document defines the high-level system architecture of the PRAHARI Platform. PRAHARI uses a containerized, event-driven, and microservices-based design to support enterprise occupational health and safety requirements. The platform integrates traditional relational databases with distributed message brokers, vector search engines, computer vision inferencing runtimes, and multi-agent AI networks.

---

## 2. Architecture Principles
- **Separation of Concerns**: Enforce strict boundary layers through clean, hexagonal code structures.
- **Event-Driven Asynchrony**: Offload latency-tolerant processes to Kafka message brokers.
- **Stateless Services**: Ensure scalability by keeping application containers free of local session state.
- **Fail-Secure Defaults**: Deny access by default (Zero-Trust) and isolate network spaces.
- **Continuous Observability**: Track all actions through tracing, structured logging, and real-time metrics.

---

## 3. System Context
The following diagram illustrates how external systems and users interact with the PRAHARI Platform boundary.

```mermaid
graph TD
    User([Safety Officer / Operator]) -->|HTTPS / WSS| GW[API Gateway]
    IPCam([RTSP IP Cameras]) -->|RTSP Stream| CVEdge[CV Edge Processing Unit]
    CVEdge -->|Inferences / HTTPS| GW
    Sensors([IoT Sensors / PLC]) -->|MQTT / LoRaWAN| IoTGW[IoT Gateway / AWS IoT Core]
    IoTGW -->|Kafka Events| MSK[Kafka Message Broker]
    
    subgraph Platform Boundary
        GW -->|gRPC / HTTP| Services[Microservices Fleet]
        MSK <--> Services
        Services --> DB[(PostgreSQL Cluster)]
        Services --> Cache[(Redis Cache)]
        Services --> VectorDB[(Vector DB / pgvector)]
    end
    
    Services -->|Logs / Traces| OTel[OpenTelemetry Collector]
    OTel -->|Telemetry Data| Obs[CloudWatch / Prometheus / Jaeger]
```

---

## 4. Enterprise Architecture & Service Landscape
The platform core consists of specialized microservices communicating via REST (external), gRPC (internal synchronous), and Kafka (asynchronous).

- **Gateway Service**: Envoy-based API Gateway handling SSL termination, CORS, routing, and rate limiting.
- **Auth & User Services**: Coordinates authentication with AWS Cognito and manages tenancy, RBAC, and ABAC profiles.
- **Chemical Service**: Manages chemical registration, SDS lifecycles, and storage compatibility audits.
- **PHA Service**: Manages Process Hazard Analysis studies, HAZOP matrices, LOPA evaluations, and checklist executions.
- **Incident Service**: Manages the incident intake pipeline, investigations, root cause analyses, and action assignments.
- **AI Agent Service**: Orchestrates LangGraph-based AI agents, vector indexing, RAG, and safety summaries.
- **CV Inference Service**: Tracks cameras, ingests edge inferences, manages YOLO model configs, and triggers alarms.
- **Digital Twin Service**: Correlates chemical spatial data, IoT events, and visual feeds to present plant safety twins.

---

## 5. Deployment Topology (AWS Architecture)
The platform is deployed inside a multi-Availability Zone (AZ) Amazon Virtual Private Cloud (VPC) segregated into public, private, and isolated subnets.

```mermaid
graph TB
    subgraph VPC [VPC: 10.0.0.0/16]
        subgraph Public Subnets [Public Subnets - Multi-AZ]
            ALB[Application Load Balancer]
            NAT[NAT Gateways]
        end
        
        subgraph Private Subnets [Private Subnets - Multi-AZ]
            EKS[Amazon EKS Cluster]
            MSK[Amazon MSK Kafka Cluster]
        end
        
        subgraph Isolated Subnets [Isolated Subnets - Multi-AZ]
            RDS[(Amazon Aurora PostgreSQL)]
            Cache[(Amazon ElastiCache Redis)]
        end
    end
    
    Route53[Route 53] --> ALB
    ALB --> EKS
    EKS --> MSK
    EKS --> RDS
    EKS --> Cache
```

---

## 6. Container Architecture (C4 Level 2)
The container diagram details the runtime components and communications within the EKS namespace.

```mermaid
graph TD
    Client([React Frontend / App]) -->|HTTPS / Port 443| Ingress[Kubernetes NGINX Ingress]
    Ingress -->|HTTP / Port 80| Gateway[Gateway Service]
    
    subgraph Kubernetes Pods
        Gateway -->|gRPC / Port 8080| AuthService[Auth Service]
        Gateway -->|gRPC / Port 8080| ChemicalService[Chemical Service]
        Gateway -->|gRPC / Port 8080| PHAService[PHA Service]
        Gateway -->|gRPC / Port 8080| AIService[AI Agent Service]
        
        ChemicalService -->|Publish Events| Kafka[Kafka Cluster]
        PHAService -->|Publish Events| Kafka
        
        KafkaStream[Kafka Consumers] -->|Subscribe| Kafka
        KafkaStream -->|Write Audits| Postgres[(PostgreSQL DB)]
    end
    
    AuthService -->|User Auth| Cognito[AWS Cognito User Pool]
    ChemicalService -->|SDS Storage| S3[AWS S3 Bucket]
    AIService -->|Vector Storage| PGVector[(Aurora pgvector)]
```

---

## 7. Core Flows & Integration

### 7.1 Authentication & Authorization Flow
```mermaid
sequenceDiagram
    autonumber
    actor Client as Client App
    participant GW as API Gateway
    participant Auth as Auth Service
    participant Cognito as AWS Cognito
    participant Service as Business Service (e.g. Chemical)

    Client->>GW: POST /api/v1/auth/login
    GW->>Auth: Forward to Auth Service
    Auth->>Cognito: InitiateAuth (User, Password)
    Cognito-->>Auth: Authentication Result (ID Token, Access Token, Refresh Token)
    Auth-->>Client: Return JWT tokens
    
    Client->>GW: GET /api/v1/chemicals (Authorization: Bearer <token>)
    GW->>GW: Verify JWT signature & Check Rate Limits
    GW->>Service: Forward request with user headers
    Service->>Service: Authorize user via Casbin RBAC/ABAC policy
    Service-->>Client: 200 OK (Chemical Records)
```

### 7.2 Async Transactional Outbox Pattern
To prevent distributed transaction failures, all microservices write state changes and events to a local PostgreSQL database inside a single database transaction, after which a CDC (Change Data Capture) system or dedicated Outbox Poller publishes the events to Kafka.

```mermaid
graph LR
    Service[Service Logic] -->|1. Begin Tx| DB[(PostgreSQL)]
    DB -->|2. Write Entity State| EntityTable[Chemicals Table]
    DB -->|3. Write Event Payload| OutboxTable[Outbox Table]
    Service -->|4. Commit Tx| DB
    Debezium[Debezium Connector / Poller] -->|5. Scan Outbox| OutboxTable
    Debezium -->|6. Publish Event| Kafka[Kafka Broker]
    Kafka -->|7. Consume Event| Analytics[Analytics Service]
```

---

## 8. Layer Integration Specifications

### 8.1 Caching Layer (Redis)
- **Cache-Aside Pattern**: Read requests query Redis first. On cache miss, fetch from PostgreSQL, write back to Redis (TTL 3600s).
- **Distributed Locks**: Use Redlock algorithm on Redis to serialize operations on high-concurrency entities (e.g., container transfer actions).
- **Rate Limiting**: Sliding window rate-limiting metrics stored inside Redis sorted sets.

### 8.2 Streaming Layer (Kafka)
- **Topic Naming Convention**: `<tenant_id>.<domain_area>.<entity_type>.<event_action>` (e.g. `1001.ehs.chemical.created`).
- **Partitioning Strategy**: Partition keys mapped to `plant_id` to ensure strictly ordered event processing per physical plant site.
- **Dead Letter Queue (DLQ)**: Failed events write to dedicated DLQ topics (e.g. `prahari.ehs.chemical.created.dlq`) with exponential backoff retry consumers.

### 8.3 Observability Layer (OpenTelemetry)
- **Context Propagation**: Trace context propagated using W3C Trace Context standard (`traceparent` header) across HTTP, gRPC, and Kafka metadata.
- **Log Aggregation**: Structured JSON logs scraped via fluent-bit and shipped to AWS CloudWatch.
- **Metrics Scraping**: Prometheus pulls JVM, Go, EKS, and business metrics, sending alerts via Alertmanager.

---

## 9. Disaster Recovery & High Availability
- **Multi-AZ Availability**: All microservice workloads replicate across three Availability Zones. Primary databases use multi-AZ synchronous replication with automatic failover.
- **Cross-Region Replication**: Database transactions replicate asynchronously to a secondary AWS DR region. S3 assets replicate via cross-region replication.
- **Failover Plan**: Route 53 DNS routing policies switch customer traffic to the secondary region if the primary region experiences service degradation.
